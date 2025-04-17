package dto

type ValidatorResponse struct {
	PublicKey        string `json:"public_key"`
	Index            int64  `json:"index"`
	Balance          string `json:"balance"`
	Status           string `json:"status"`
	EffectiveBalance string `json:"effective_balance"`
	ActivationEpoch  int64  `json:"activation_epoch"`
	ExitEpoch        int64  `json:"exit_epoch"`
	ChainID          int64  `json:"chain_id"`
	RetrievedFrom    string `json:"retrieved_from"`
}

type ValidatorsResponse struct {
	Validators []ValidatorResponse `json:"validators"`
}
