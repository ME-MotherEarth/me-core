package api

// MECTSupply represents the structure for mect supply that is returned by api routes
type MECTSupply struct {
	InitialMinted string `json:"initialMinted"`
	Supply        string `json:"supply"`
	Burned        string `json:"burned"`
	Minted        string `json:"minted"`
}
