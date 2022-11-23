package main

import (
	"time"
)

type MetersAggregates struct {
	Battery struct {
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
		NumMetersAggregated               float64   `json:"num_meters_aggregated"`
		Timeout                           float64   `json:"timeout"`
	} `json:"battery"`
	Load struct {
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
		Timeout                           float64   `json:"timeout"`
	} `json:"load"`
	Site struct {
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
		NumMetersAggregated               float64   `json:"num_meters_aggregated"`
		Timeout                           float64   `json:"timeout"`
	} `json:"site"`
	Solar struct {
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
		NumMetersAggregated               float64   `json:"num_meters_aggregated"`
		Timeout                           float64   `json:"timeout"`
	} `json:"solar"`
}
