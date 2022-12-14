package main

type MetersReadings struct {
	SynchrometerX struct {
		Data struct {
			Cts []struct {
				Ct     float64 `json:"ct"`
				EExpWs float64 `json:"eExp_Ws"`
				EImpWs float64 `json:"eImp_Ws"`
				IA     float64 `json:"i_A"`
				PW     float64 `json:"p_W"`
				QVar   float64 `json:"q_VAR"`
				VV     float64 `json:"v_V"`
			} `json:"cts"`
			FirmwareVersion string `json:"firmwareVersion"`
		} `json:"data"`
		Error string `json:"error"`
	} `json:"synchrometerX"`
	SynchrometerY struct {
		Data struct {
			Cts []struct {
				Ct     float64 `json:"ct"`
				EExpWs float64 `json:"eExp_Ws"`
				EImpWs float64 `json:"eImp_Ws"`
				IA     float64 `json:"i_A"`
				PW     float64 `json:"p_W"`
				QVar   float64 `json:"q_VAR"`
				VV     float64 `json:"v_V"`
			} `json:"cts"`
			FirmwareVersion string `json:"firmwareVersion"`
		} `json:"data"`
		Error string `json:"error"`
	} `json:"synchrometerY"`
}
