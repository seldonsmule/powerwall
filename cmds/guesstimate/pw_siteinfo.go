package main

type SiteInfo struct {
	GridCode struct {
		Country           string `json:"country"`
		Distributor       string `json:"distributor"`
		GridCode          string `json:"grid_code"`
		GridCodeOverrides []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"grid_code_overrides"`
		GridFreqSetting    float64 `json:"grid_freq_setting"`
		GridPhaseSetting   string  `json:"grid_phase_setting"`
		GridVoltageSetting float64 `json:"grid_voltage_setting"`
		Region             string  `json:"region"`
		Retailer           string  `json:"retailer"`
		State              string  `json:"state"`
		Utility            string  `json:"utility"`
	} `json:"grid_code"`
	MaxSiteMeterPowerKW    float64 `json:"max_site_meter_power_kW"`
	MaxSystemEnergyKWh     float64 `json:"max_system_energy_kWh"`
	MaxSystemPowerKW       float64 `json:"max_system_power_kW"`
	MinSiteMeterPowerKW    float64 `json:"min_site_meter_power_kW"`
	NominalSystemEnergyKWh float64 `json:"nominal_system_energy_kWh"`
	NominalSystemPowerKW   float64 `json:"nominal_system_power_kW"`
	SiteName               string  `json:"site_name"`
	Timezone               string  `json:"timezone"`
}
