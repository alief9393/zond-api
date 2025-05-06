package dto

type ValidatorResponse struct {
	PublicKey        string `json:"public_key" example:"0xabc123..."`
	Index            int64  `json:"index" example:"1234"`
	Balance          string `json:"balance" example:"32000000000"`
	Status           string `json:"status" example:"active"`
	EffectiveBalance string `json:"effective_balance" example:"32000000000"`
	ActivationEpoch  int64  `json:"activation_epoch" example:"194"`
	ExitEpoch        int64  `json:"exit_epoch" example:"0"`
	ChainID          int64  `json:"chain_id" example:"1"`
	RetrievedFrom    string `json:"retrieved_from" example:"lighthouse"`
}

type ValidatorsResponse struct {
	Validators []ValidatorResponse `json:"validators"`
}

type ValidatorDetailResponse struct {
	PublicKey        string `json:"public_key" example:"0xabc123..."`
	Index            int    `json:"index" example:"1234"`
	Balance          int64  `json:"balance" example:"32000000000"`
	Status           string `json:"status" example:"active"`
	EffectiveBalance int64  `json:"effective_balance" example:"32000000000"`
	ActivationEpoch  int64  `json:"activation_epoch" example:"194"`
	ExitEpoch        int64  `json:"exit_epoch" example:"0"`
	ChainID          int    `json:"chain_id" example:"1"`
	RetrievedFrom    string `json:"retrieved_from" example:"prysm"`
}
