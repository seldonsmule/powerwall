package main

type Networks []struct {
	Active   bool  `json:"active"`
	Dhcp     *bool `json:"dhcp"`
	Enabled  bool  `json:"enabled"`
	ExtraIps []struct {
		Ip      string  `json:"ip"`
		Netmask float64 `json:"netmask"`
	} `json:"extra_ips,omitempty"`
	IfaceNetworkInfo *struct {
		Gateway    string `json:"gateway"`
		HwAddress  string `json:"hw_address"`
		Interface  string `json:"interface"`
		IpNetworks []struct {
			Ip   string `json:"ip"`
			Mask string `json:"mask"`
		} `json:"ip_networks"`
		NetworkName    string  `json:"network_name"`
		SignalStrength float64 `json:"signal_strength"`
		State          string  `json:"state"`
		StateReason    string  `json:"state_reason"`
	} `json:"iface_network_info,omitempty"`
	Interface             string `json:"interface"`
	LastInternetConnected bool   `json:"lastInternetConnected"`
	LastTeslaConnected    bool   `json:"lastTeslaConnected"`
	NetworkName           string `json:"network_name"`
	Primary               bool   `json:"primary"`
	SecurityType          string `json:"security_type,omitempty"`
	Username              string `json:"username"`
}
