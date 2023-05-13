package main

import "time"

type CelestiaHeadResponse struct {
	Header struct {
		Version            interface{} `json:"-"`
		ChainID            string      `json:"chain_id"`
		Height             string      `json:"height"`
		Time               time.Time   `json:"time"`
		LastBlockID        interface{} `json:"-"`
		LastCommitHash     string      `json:"last_commit_hash"`
		DataHash           string      `json:"data_hash"`
		ValidatorsHash     string      `json:"validators_hash"`
		NextValidatorsHash string      `json:"next_validators_hash"`
		ConsensusHash      string      `json:"consensus_hash"`
		AppHash            string      `json:"app_hash"`
		LastResultsHash    string      `json:"last_results_hash"`
		EvidenceHash       string      `json:"evidence_hash"`
		ProposerAddress    string      `json:"proposer_address"`
	} `json:"header"`
	ValidatorSet interface{} `json:"-"`
	Commit       interface{} `json:"-"`
	Dah          interface{} `json:"-"`
}
