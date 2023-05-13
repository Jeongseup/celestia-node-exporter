# Celestia Node Exporter

Exports some information(chain-id, height, sync time interval) your celestia da type's node(bridge, light, full) for celestia node operator.

## Build & Install

```bash
# build in ./build
make build

# install in $GOBIN
make install
```

## Available Flags
```bash 
Usage of ./build/celestia-node-exporter:
  -listen-address string
        Binary Listen address (default ":8000")
  -rpc-address string
        Celesit Node API Address (default "http://localhost:26659")
  -timeout int
        Exporter Timeout Second When Calling Your Node (default 10)
  -v    show version
  -version
        show version
```

## Start
```bash
# check version
celestia-node-exporter --version

# simple start
celestia-node-exporter --api-address http://example.celestia-gateway-api.com:26659

# start with custom flags
celestia-node-exporter \ 
    --timeout 5 \
    --listen-address "8888" \
    --api-address http://example.celestia-gateway-api.com:26659 
```

## Exported Metrics
```
# HELP celestia_node_exporter_current_height exposing currrent your node's height.
# TYPE celestia_node_exporter_current_height gauge
celestia_node_exporter_current_height 0
# HELP celestia_node_exporter_current_synctime_interval exposing currrent your node's time interval to check synced node with now and blocktimestamp.
# TYPE celestia_node_exporter_current_synctime_interval gauge
celestia_node_exporter_current_synctime_interval 0
# HELP celestia_node_exporter_failed_counter Example metric with a string value.
# TYPE celestia_node_exporter_failed_counter counter
celestia_node_exporter_failed_counter 0
```

## References

* [sui-exporter](https://github.com/rpcpool/sui-exporter)
* [celestia-node](https://github.com/celestiaorg/celestia-node)