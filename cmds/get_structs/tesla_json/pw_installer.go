package main

type Installer struct {
	BackupConfiguration      string   `json:"backup_configuration"`
	Company                  string   `json:"company"`
	CustomerID               string   `json:"customer_id"`
	Email                    string   `json:"email"`
	HasResidualCurrentDevice bool     `json:"has_residual_current_device"`
	InstallationTypes        []string `json:"installation_types"`
	Location                 string   `json:"location"`
	Mounting                 string   `json:"mounting"`
	Phone                    string   `json:"phone"`
	RunSitemaster            bool     `json:"run_sitemaster"`
	SolarInstallation        string   `json:"solar_installation"`
	SolarInstallationType    string   `json:"solar_installation_type"`
	VerifiedConfig           bool     `json:"verified_config"`
	Wiring                   string   `json:"wiring"`
}
