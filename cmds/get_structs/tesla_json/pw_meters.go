package main

type Meters []struct {
	Cts []struct {
		Inverted             []bool  `json:"inverted"`
		RealPowerScaleFactor float64 `json:"real_power_scale_factor,omitempty"`
		Type                 string  `json:"type"`
		Valid                []bool  `json:"valid"`
	} `json:"cts"`
	Serial  string `json:"serial"`
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}
