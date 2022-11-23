package main

type Customer struct {
	GridServices    bool `json:"grid_services"`
	LimitedWarranty bool `json:"limited_warranty"`
	Marketing       bool `json:"marketing"`
	PrivacyNotice   bool `json:"privacy_notice"`
	Registered      bool `json:"registered"`
}
