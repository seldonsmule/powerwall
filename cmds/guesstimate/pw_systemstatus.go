package main

import (
	"time"
)

type SystemStatus struct {
	AllEnableLinesHigh     bool    `json:"all_enable_lines_high"`
	AuxiliaryLoad          float64 `json:"auxiliary_load"`
	AvailableBlocks        float64 `json:"available_blocks"`
	AvailableChargerBlocks float64 `json:"available_charger_blocks"`
	BatteryBlocks          []struct {
		OpSeqState             string        `json:"OpSeqState"`
		PackagePartNumber      string        `json:"PackagePartNumber"`
		PackageSerialNumber    string        `json:"PackageSerialNumber"`
		Type                   string        `json:"Type"`
		BackupReady            bool          `json:"backup_ready"`
		ChargePowerClamped     bool          `json:"charge_power_clamped"`
		DisabledReasons        []interface{} `json:"disabled_reasons"`
		EnergyCharged          float64       `json:"energy_charged"`
		EnergyDischarged       float64       `json:"energy_discharged"`
		FOut                   float64       `json:"f_out"`
		IOut                   float64       `json:"i_out"`
		NominalEnergyRemaining float64       `json:"nominal_energy_remaining"`
		NominalFullPackEnergy  float64       `json:"nominal_full_pack_energy"`
		OffGrid                bool          `json:"off_grid"`
		POut                   float64       `json:"p_out"`
		PinvGridState          string        `json:"pinv_grid_state"`
		PinvState              string        `json:"pinv_state"`
		QOut                   float64       `json:"q_out"`
		VOut                   float64       `json:"v_out"`
		Version                string        `json:"version"`
		VfMode                 bool          `json:"vf_mode"`
		WobbleDetected         bool          `json:"wobble_detected"`
	} `json:"battery_blocks"`
	BatteryTargetPower               float64       `json:"battery_target_power"`
	BatteryTargetReactivePower       float64       `json:"battery_target_reactive_power"`
	BlocksControlled                 float64       `json:"blocks_controlled"`
	CanReboot                        string        `json:"can_reboot"`
	CommandSource                    string        `json:"command_source"`
	ExpectedEnergyRemaining          float64       `json:"expected_energy_remaining"`
	FfrPowerAvailabilityHigh         float64       `json:"ffr_power_availability_high"`
	FfrPowerAvailabilityLow          float64       `json:"ffr_power_availability_low"`
	GridFaults                       []interface{} `json:"grid_faults"`
	GridServicesPower                float64       `json:"grid_services_power"`
	HardwareCapabilityChargePower    float64       `json:"hardware_capability_charge_power"`
	HardwareCapabilityDischargePower float64       `json:"hardware_capability_discharge_power"`
	InstantaneousMaxApparentPower    float64       `json:"instantaneous_max_apparent_power"`
	InstantaneousMaxChargePower      float64       `json:"instantaneous_max_charge_power"`
	InstantaneousMaxDischargePower   float64       `json:"instantaneous_max_discharge_power"`
	InverterNominalUsablePower       float64       `json:"inverter_nominal_usable_power"`
	LastToggleTimestamp              time.Time     `json:"last_toggle_timestamp"`
	LoadChargeConstraint             float64       `json:"load_charge_constraint"`
	MaxApparentPower                 float64       `json:"max_apparent_power"`
	MaxChargePower                   float64       `json:"max_charge_power"`
	MaxDischargePower                float64       `json:"max_discharge_power"`
	MaxPowerEnergyRemaining          float64       `json:"max_power_energy_remaining"`
	MaxPowerEnergyToBeCharged        float64       `json:"max_power_energy_to_be_charged"`
	MaxSustainedRampRate             float64       `json:"max_sustained_ramp_rate"`
	NominalEnergyRemaining           float64       `json:"nominal_energy_remaining"`
	NominalFullPackEnergy            float64       `json:"nominal_full_pack_energy"`
	Primary                          bool          `json:"primary"`
	Score                            float64       `json:"score"`
	SmartInvDeltaP                   float64       `json:"smart_inv_delta_p"`
	SmartInvDeltaQ                   float64       `json:"smart_inv_delta_q"`
	SolarRealPowerLimit              float64       `json:"solar_real_power_limit"`
	SystemIslandState                string        `json:"system_island_state"`
}
