package main

import (
	"time"
)

type Powerwalls struct {
	BubbleShedding             bool        `json:"bubble_shedding"`
	CheckingIfOffgrid          bool        `json:"checking_if_offgrid"`
	Enumerating                bool        `json:"enumerating"`
	GatewayDin                 string      `json:"gateway_din"`
	GridCodeValidating         bool        `json:"grid_code_validating"`
	GridQualifying             bool        `json:"grid_qualifying"`
	Msa                        interface{} `json:"msa"`
	OnGridCheckError           string      `json:"on_grid_check_error"`
	PhaseDetectionLastError    string      `json:"phase_detection_last_error"`
	PhaseDetectionNotAvailable bool        `json:"phase_detection_not_available"`
	Powerwalls                 []struct {
		PackagePartNumber       string      `json:"PackagePartNumber"`
		PackageSerialNumber     string      `json:"PackageSerialNumber"`
		Type                    string      `json:"Type"`
		BcType                  interface{} `json:"bc_type"`
		CommissioningDiagnostic struct {
			Alert    bool   `json:"alert"`
			Category string `json:"category"`
			Checks   []struct {
				Checks    interface{} `json:"checks"`
				Debug     struct{}    `json:"debug"`
				EndTime   time.Time   `json:"end_time"`
				Message   string      `json:"message"`
				Name      string      `json:"name"`
				Results   struct{}    `json:"results"`
				StartTime time.Time   `json:"start_time"`
				Status    string      `json:"status"`
			} `json:"checks"`
			Disruptive bool        `json:"disruptive"`
			Inputs     interface{} `json:"inputs"`
			Name       string      `json:"name"`
		} `json:"commissioning_diagnostic"`
		GridReconnectionTimeSeconds float64 `json:"grid_reconnection_time_seconds"`
		GridState                   string  `json:"grid_state"`
		InConfig                    bool    `json:"in_config"`
		Type                        string  `json:"type"`
		UnderPhaseDetection         bool    `json:"under_phase_detection"`
		UpdateDiagnostic            struct {
			Alert    bool   `json:"alert"`
			Category string `json:"category"`
			Checks   []struct {
				Checks    interface{} `json:"checks"`
				Debug     interface{} `json:"debug"`
				EndTime   interface{} `json:"end_time"`
				Name      string      `json:"name"`
				Progress  float64     `json:"progress"`
				Results   interface{} `json:"results"`
				StartTime interface{} `json:"start_time"`
				Status    string      `json:"status"`
			} `json:"checks"`
			Disruptive bool        `json:"disruptive"`
			Inputs     interface{} `json:"inputs"`
			Name       string      `json:"name"`
		} `json:"update_diagnostic"`
		Updating bool `json:"updating"`
	} `json:"powerwalls"`
	RunningPhaseDetection bool        `json:"running_phase_detection"`
	States                interface{} `json:"states"`
	Sync                  struct {
		CommissioningDiagnostic struct {
			Alert    bool   `json:"alert"`
			Category string `json:"category"`
			Checks   []struct {
				Checks    interface{} `json:"checks"`
				Debug     struct{}    `json:"debug"`
				EndTime   time.Time   `json:"end_time"`
				Message   string      `json:"message"`
				Name      string      `json:"name"`
				Results   struct{}    `json:"results"`
				StartTime time.Time   `json:"start_time"`
				Status    string      `json:"status"`
			} `json:"checks"`
			Disruptive bool        `json:"disruptive"`
			Inputs     interface{} `json:"inputs"`
			Name       string      `json:"name"`
		} `json:"commissioning_diagnostic"`
		UpdateDiagnostic struct {
			Alert    bool   `json:"alert"`
			Category string `json:"category"`
			Checks   []struct {
				Checks    interface{} `json:"checks"`
				Debug     interface{} `json:"debug"`
				EndTime   interface{} `json:"end_time"`
				Name      string      `json:"name"`
				Progress  float64     `json:"progress"`
				Results   interface{} `json:"results"`
				StartTime interface{} `json:"start_time"`
				Status    string      `json:"status"`
			} `json:"checks"`
			Disruptive bool        `json:"disruptive"`
			Inputs     interface{} `json:"inputs"`
			Name       string      `json:"name"`
		} `json:"update_diagnostic"`
		Updating bool `json:"updating"`
	} `json:"sync"`
	Updating bool `json:"updating"`
}
