package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger

	showVersion   bool
	Version       string
	apiAddress    = flag.String("api-address", "", "Celesit Node Gateway API Address")
	listenAddress = flag.String("listen-address", "8000", "Binary Listen address")
	Timeout       = flag.Int(("timeout"), 10, "Exporter Timeout Second When Calling Your Node")
)

func main() {
	// init
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")

	logger = logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	flag.Parse()

	if showVersion {
		fmt.Println("version:", Version)
		os.Exit(0)
	}
	if *apiAddress == "" {
		logger.Fatalln("Please specify --api-address")
	}

	// Printing flag informations
	logger.Infoln("Start Celestia Node Exporter")
	logger.Infoln("Version: ", *&Version)
	logger.Infoln("Listening on: ", *listenAddress+"/metrics")
	logger.Infoln("Using Celestia Node Gateway API Addresss: ", *apiAddress)
	logger.Infoln("API Call Timeout: ", *Timeout, "seconds")

	// create prometheus new registry and registring custom metrics
	registry := prometheus.NewRegistry()
	registry.MustRegister(failCount, chainIDMetric, heightMetric, syncTimeMetric)

	// create new prom handler
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog:          logger,
		EnableOpenMetrics: false,
		ErrorHandling:     promhttp.ContinueOnError,
		Timeout:           time.Duration(time.Duration(*Timeout).Seconds()),
	})

	// set created handlers
	http.Handle("/metrics", handler)
	http.HandleFunc("/", DefaultDescriptionHandler)
	logger.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", *listenAddress), nil))
}

func DefaultDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	chainID, height, blockTimestamp, err := getCelestiaNodeInfo()
	if err != nil {
		// if collecting block data, increase fail count
		failCount.Inc()
	}

	// update metrics value
	chainIDMetric.WithLabelValues(chainID).Set(1)
	heightMetric.Set(float64(height))
	syncTimeMetric.Set(float64(time.Now().Unix() - blockTimestamp.Unix()))

	w.Write([]byte(`
	<html>
		<head>
			<title>
				Celestia Node Exporter
			</title>
		</head>
		<body>
			<h1>
				Celestia Node Exporter
			</h1>
			<p>
				<a href="/metrics">Metrics</a>
			</p>
		</body>
	</html>
	`))

}

func getCelestiaNodeInfo() (chainID string, height uint64, blockTimestamp time.Time, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequest("GET", *apiAddress+"/head", nil)
	if err != nil {
		logger.Errorln("failed creating http requester:", err)
		return "", 0, time.Time{}, fmt.Errorf(err.Error())
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorln("failed doing http request:", err)
		return "", 0, time.Time{}, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorln("failed reading response body data:", err)
		return "", 0, time.Time{}, fmt.Errorf(err.Error())

	}

	var result CelestiaHeadResponse
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Errorln("failed unmarshaling json data, check your request URL", resp.Request.URL)
		return "", 0, time.Time{}, fmt.Errorf(err.Error())
	}

	height, err = strconv.ParseUint(result.Header.Height, 10, 64)
	if err != nil {
		logger.Errorln("failed parsing uint from string height")
		return "", 0, time.Time{}, fmt.Errorf(err.Error())
	}

	// success requesting and parsing data
	logger.Infoln("got celestia node inforamtions, Chain ID: ", result.Header.ChainID)
	logger.Infoln("got celestia node inforamtions, Height: ", result.Header.Height)
	logger.Infoln("got celestia node inforamtions, Block Timestamp: ", result.Header.Time)
	return result.Header.ChainID, height, result.Header.Time, nil
}
