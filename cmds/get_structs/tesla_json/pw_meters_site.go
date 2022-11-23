package main

import (
	"time"
)

type MetersSite []struct {
	CachedReadings struct {
		EnergyExported                    float64   `json:"energy_exported"`
		EnergyImported                    float64   `json:"energy_imported"`
		Frequency                         float64   `json:"frequency"`
		IACurrent                         float64   `json:"i_a_current"`
		IBCurrent                         float64   `json:"i_b_current"`
		ICCurrent                         float64   `json:"i_c_current"`
		InstantApparentPower              float64   `json:"instant_apparent_power"`
		InstantAverageCurrent             float64   `json:"instant_average_current"`
		InstantAverageVoltage             float64   `json:"instant_average_voltage"`
		InstantPower                      float64   `json:"instant_power"`
		InstantReactivePower              float64   `json:"instant_reactive_power"`
		InstantTotalCurrent               float64   `json:"instant_total_current"`
		LastCommunicationTime             time.Time `json:"last_communication_time"`
		LastPhaseEnergyCommunicationTime  time.Time `json:"last_phase_energy_communication_time"`
		LastPhasePowerCommunicationTime   time.Time `json:"last_phase_power_communication_time"`
		LastPhaseVoltageCommunicationTime time.Time `json:"last_phase_voltage_communication_time"`
		ReactivePowerA                    float64   `json:"reactive_power_a"`
		ReactivePowerB                    float64   `json:"reactive_power_b"`
		RealPowerA                        float64   `json:"real_power_a"`
		RealPowerB                        float64   `json:"real_power_b"`
		SerialNumber                      string    `json:"serial_number"`
		Timeout                           float64   `json:"timeout"`
		VL1n                              float64   `json:"v_l1n"`
		VL2n                              float64   `json:"v_l2n"`
	} `json:"Cached_readings"`
	Connection struct {
		DeviceSerial string   `json:"device_serial"`
		HTTPSConf    struct{} `json:"https_conf"`
		ShortID      string   `json:"short_id"`
	} `json:"connection"`
	Cts      []bool  `json:"cts"`
	ID       float64 `json:"id"`
	Inverted []bool  `json:"inverted"`
	Location string  `json:"location"`
	Type     string  `json:"type"`
}
