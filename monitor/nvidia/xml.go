package nvidia

import "encoding/xml"

type NvidiaSmi struct {
	XMLName       xml.Name `xml:"nvidia_smi_log"`
	Timestamp     string   `xml:"timestamp"`
	DriverVersion string   `xml:"driver_version"`
	CudaVersion   string   `xml:"cuda_version"`
	AttachedGPUs  int      `xml:"attached_gpus"`
	GPU           []GPU    `xml:"gpu"`
}

type GPU struct {
	ID                         string                `xml:"id,attr"`
	ProductName                string                `xml:"product_name"`
	ProductBrand               string                `xml:"product_brand"`
	ProductArchitecture        string                `xml:"product_architecture"`
	DisplayMode                string                `xml:"display_mode"`
	DisplayAttached            string                `xml:"display_attached"`
	DisplayActive              string                `xml:"display_active"`
	PersistenceMode            string                `xml:"persistence_mode"`
	AddressingMode             string                `xml:"addressing_mode"`
	MigMode                    MigMode               `xml:"mig_mode"`
	MigDevices                 string                `xml:"mig_devices"`
	AccountingMode             string                `xml:"accounting_mode"`
	AccountingModeBufferSize   int                   `xml:"accounting_mode_buffer_size"`
	DriverModel                DriverModel           `xml:"driver_model"`
	Serial                     string                `xml:"serial"`
	UUID                       string                `xml:"uuid"`
	MinorNumber                int                   `xml:"minor_number"`
	VbiosVersion               string                `xml:"vbios_version"`
	MultigpuBoard              string                `xml:"multigpu_board"`
	BoardID                    string                `xml:"board_id"`
	BoardPartNumber            string                `xml:"board_part_number"`
	GPUPartNumber              string                `xml:"gpu_part_number"`
	GPUFruPartNumber           string                `xml:"gpu_fru_part_number"`
	PlatformInfo               PlatformInfo          `xml:"platformInfo"`
	InforomVersion             InforomVersion        `xml:"inforom_version"`
	InforomBbxFlush            InforomBbxFlush       `xml:"inforom_bbx_flush"`
	GpuOperationMode           GpuOperationMode      `xml:"gpu_operation_mode"`
	C2cMode                    string                `xml:"c2c_mode"`
	GpuVirtualizationMode      GpuVirtualizationMode `xml:"gpu_virtualization_mode"`
	GpuResetStatus             GpuResetStatus        `xml:"gpu_reset_status"`
	GpuRecoveryAction          string                `xml:"gpu_recovery_action"`
	GspFirmwareVersion         string                `xml:"gsp_firmware_version"`
	IBMNPU                     IBMNPU                `xml:"ibmnpu"`
	PCI                        PCI                   `xml:"pci"`
	FanSpeed                   string                `xml:"fan_speed"`
	PerformanceState           string                `xml:"performance_state"`
	ClocksEventReasons         ClocksEventReasons    `xml:"clocks_event_reasons"`
	ClocksEventReasonsCounters ClocksEventCounters   `xml:"clocks_event_reasons_counters"`
	SparseOperationMode        string                `xml:"sparse_operation_mode"`
	FbMemoryUsage              MemoryUsage           `xml:"fb_memory_usage"`
	Bar1MemoryUsage            MemoryUsage           `xml:"bar1_memory_usage"`
	CcProtectedMemoryUsage     MemoryUsage           `xml:"cc_protected_memory_usage"`
	ComputeMode                string                `xml:"compute_mode"`
	Utilization                Utilization           `xml:"utilization"`
	EncoderStats               EncoderStats          `xml:"encoder_stats"`
	FbcStats                   EncoderStats          `xml:"fbc_stats"`
	DramEncryptionMode         DramEncryptionMode    `xml:"dram_encryption_mode"`
	EccMode                    EccMode               `xml:"ecc_mode"`
	EccErrors                  EccErrors             `xml:"ecc_errors"`
	RetiredPages               RetiredPages          `xml:"retired_pages"`
	RemappedRows               string                `xml:"remapped_rows"`
	Temperature                Temperature           `xml:"temperature"`
	SupportedGpuTargetTemp     SupportedTempRange    `xml:"supported_gpu_target_temp"`
	GPUPowerReadings           PowerReadings         `xml:"gpu_power_readings"`
	GPUMemoryPowerReadings     SimplePowerReadings   `xml:"gpu_memory_power_readings"`
	ModulePowerReadings        ModulePowerReadings   `xml:"module_power_readings"`
	PowerSmoothing             string                `xml:"power_smoothing"`
	PowerProfiles              PowerProfiles         `xml:"power_profiles"`
	Clocks                     Clocks                `xml:"clocks"`
	ApplicationsClocks         ClockPair             `xml:"applications_clocks"`
	DefaultApplicationsClocks  ClockPair             `xml:"default_applications_clocks"`
	DeferredClocks             DeferredClocks        `xml:"deferred_clocks"`
	MaxClocks                  MaxClocks             `xml:"max_clocks"`
	MaxCustomerBoostClocks     MaxCustomerBoostClock `xml:"max_customer_boost_clocks"`
	ClockPolicy                ClockPolicy           `xml:"clock_policy"`
	Voltage                    Voltage               `xml:"voltage"`
	Fabric                     Fabric                `xml:"fabric"`
	SupportedClocks            SupportedClocks       `xml:"supported_clocks"`
	Processes                  string                `xml:"processes"`
	AccountedProcesses         string                `xml:"accounted_processes"`
	Capabilities               Capabilities          `xml:"capabilities"`
}

type MigMode struct {
	CurrentMig string `xml:"current_mig"`
	PendingMig string `xml:"pending_mig"`
}

type DriverModel struct {
	CurrentDM string `xml:"current_dm"`
	PendingDM string `xml:"pending_dm"`
}

type PlatformInfo struct {
	ChassisSerialNumber string `xml:"chassis_serial_number"`
	SlotNumber          string `xml:"slot_number"`
	TrayIndex           string `xml:"tray_index"`
	HostID              string `xml:"host_id"`
	PeerType            string `xml:"peer_type"`
	ModuleID            int    `xml:"module_id"`
	GpuFabricGuid       string `xml:"gpu_fabric_guid"`
}

type InforomVersion struct {
	ImgVersion string `xml:"img_version"`
	OemObject  string `xml:"oem_object"`
	EccObject  string `xml:"ecc_object"`
	PwrObject  string `xml:"pwr_object"`
}

type InforomBbxFlush struct {
	LatestTimestamp string `xml:"latest_timestamp"`
	LatestDuration  string `xml:"latest_duration"`
}

type GpuOperationMode struct {
	CurrentGom string `xml:"current_gom"`
	PendingGom string `xml:"pending_gom"`
}

type GpuVirtualizationMode struct {
	VirtualizationMode    string `xml:"virtualization_mode"`
	HostVgpuMode          string `xml:"host_vgpu_mode"`
	VgpuHeterogeneousMode string `xml:"vgpu_heterogeneous_mode"`
}

type GpuResetStatus struct {
	ResetRequired            string `xml:"reset_required"`
	DrainAndResetRecommended string `xml:"drain_and_reset_recommended"`
}

type IBMNPU struct {
	RelaxedOrderingMode string `xml:"relaxed_ordering_mode"`
}

type PCI struct {
	PciBus                string         `xml:"pci_bus"`
	PciDevice             string         `xml:"pci_device"`
	PciDomain             string         `xml:"pci_domain"`
	PciBaseClass          string         `xml:"pci_base_class"`
	PciSubClass           string         `xml:"pci_sub_class"`
	PciDeviceID           string         `xml:"pci_device_id"`
	PciBusID              string         `xml:"pci_bus_id"`
	PciSubSystemID        string         `xml:"pci_sub_system_id"`
	PciGpuLinkInfo        PciGpuLinkInfo `xml:"pci_gpu_link_info"`
	PciBridgeChip         PciBridgeChip  `xml:"pci_bridge_chip"`
	ReplayCounter         int            `xml:"replay_counter"`
	ReplayRolloverCounter int            `xml:"replay_rollover_counter"`
	TxUtil                string         `xml:"tx_util"`
	RxUtil                string         `xml:"rx_util"`
	AtomicCapsOutbound    string         `xml:"atomic_caps_outbound"`
	AtomicCapsInbound     string         `xml:"atomic_caps_inbound"`
}

type PciGpuLinkInfo struct {
	PcieGen    PcieGen    `xml:"pcie_gen"`
	LinkWidths LinkWidths `xml:"link_widths"`
}

type PcieGen struct {
	MaxLinkGen           int `xml:"max_link_gen"`
	CurrentLinkGen       int `xml:"current_link_gen"`
	DeviceCurrentLinkGen int `xml:"device_current_link_gen"`
	MaxDeviceLinkGen     int `xml:"max_device_link_gen"`
	MaxHostLinkGen       int `xml:"max_host_link_gen"`
}

type LinkWidths struct {
	MaxLinkWidth     string `xml:"max_link_width"`
	CurrentLinkWidth string `xml:"current_link_width"`
}

type PciBridgeChip struct {
	BridgeChipType string `xml:"bridge_chip_type"`
	BridgeChipFw   string `xml:"bridge_chip_fw"`
}

type ClocksEventReasons struct {
	GpuIdle                   string `xml:"clocks_event_reason_gpu_idle"`
	ApplicationsClocksSetting string `xml:"clocks_event_reason_applications_clocks_setting"`
	SwPowerCap                string `xml:"clocks_event_reason_sw_power_cap"`
	HwSlowdown                string `xml:"clocks_event_reason_hw_slowdown"`
	HwThermalSlowdown         string `xml:"clocks_event_reason_hw_thermal_slowdown"`
	HwPowerBrakeSlowdown      string `xml:"clocks_event_reason_hw_power_brake_slowdown"`
	SyncBoost                 string `xml:"clocks_event_reason_sync_boost"`
	SwThermalSlowdown         string `xml:"clocks_event_reason_sw_thermal_slowdown"`
	DisplayClocksSetting      string `xml:"clocks_event_reason_display_clocks_setting"`
}

type ClocksEventCounters struct {
	SwPowerCap      string `xml:"clocks_event_reasons_counters_sw_power_cap"`
	SyncBoost       string `xml:"clocks_event_reasons_counters_sync_boost"`
	SwThermSlowdown string `xml:"clocks_event_reasons_counters_sw_therm_slowdown"`
	HwThermSlowdown string `xml:"clocks_event_reasons_counters_hw_therm_slowdown"`
	HwPowerBrake    string `xml:"clocks_event_reasons_counters_hw_power_brake"`
}

type MemoryUsage struct {
	Total    string `xml:"total"`
	Reserved string `xml:"reserved,omitempty"`
	Used     string `xml:"used"`
	Free     string `xml:"free"`
}

type Utilization struct {
	GpuUtil     string `xml:"gpu_util"`
	MemoryUtil  string `xml:"memory_util"`
	EncoderUtil string `xml:"encoder_util"`
	DecoderUtil string `xml:"decoder_util"`
	JpegUtil    string `xml:"jpeg_util"`
	OfaUtil     string `xml:"ofa_util"`
}

type EncoderStats struct {
	SessionCount   int `xml:"session_count"`
	AverageFps     int `xml:"average_fps"`
	AverageLatency int `xml:"average_latency"`
}

type DramEncryptionMode struct {
	CurrentDramEncryption string `xml:"current_dram_encryption"`
	PendingDramEncryption string `xml:"pending_dram_encryption"`
}

type EccMode struct {
	CurrentEcc string `xml:"current_ecc"`
	PendingEcc string `xml:"pending_ecc"`
}

type EccErrors struct {
	Volatile                   VolatileEccErrors      `xml:"volatile"`
	Aggregate                  AggregateEccErrors     `xml:"aggregate"`
	AggregateUncorrectableSram AggregateUncorrectable `xml:"aggregate_uncorrectable_sram_sources"`
}

type VolatileEccErrors struct {
	SramCorrectable         string `xml:"sram_correctable"`
	SramUncorrectableParity string `xml:"sram_uncorrectable_parity"`
	SramUncorrectableSecded string `xml:"sram_uncorrectable_secded"`
	DramCorrectable         string `xml:"dram_correctable"`
	DramUncorrectable       string `xml:"dram_uncorrectable"`
}

type AggregateEccErrors struct {
	SramCorrectable         string `xml:"sram_correctable"`
	SramUncorrectableParity string `xml:"sram_uncorrectable_parity"`
	SramUncorrectableSecded string `xml:"sram_uncorrectable_secded"`
	DramCorrectable         string `xml:"dram_correctable"`
	DramUncorrectable       string `xml:"dram_uncorrectable"`
	SramThresholdExceeded   string `xml:"sram_threshold_exceeded"`
}

type AggregateUncorrectable struct {
	SramL2              string `xml:"sram_l2"`
	SramSm              string `xml:"sram_sm"`
	SramMicrocontroller string `xml:"sram_microcontroller"`
	SramPcie            string `xml:"sram_pcie"`
	SramOther           string `xml:"sram_other"`
}

type RetiredPages struct {
	MultipleSingleBit RetirementInfo `xml:"multiple_single_bit_retirement"`
	DoubleBit         RetirementInfo `xml:"double_bit_retirement"`
	PendingBlacklist  string         `xml:"pending_blacklist"`
	PendingRetirement string         `xml:"pending_retirement"`
}

type RetirementInfo struct {
	RetiredCount    string `xml:"retired_count"`
	RetiredPagelist string `xml:"retired_pagelist"`
}

type Temperature struct {
	GpuTemp                string `xml:"gpu_temp"`
	GpuTempTlimit          string `xml:"gpu_temp_tlimit"`
	GpuTempMaxThreshold    string `xml:"gpu_temp_max_threshold"`
	GpuTempSlowThreshold   string `xml:"gpu_temp_slow_threshold"`
	GpuTempMaxGpuThreshold string `xml:"gpu_temp_max_gpu_threshold"`
	GpuTargetTemperature   string `xml:"gpu_target_temperature"`
	MemoryTemp             string `xml:"memory_temp"`
	GpuTempMaxMemThreshold string `xml:"gpu_temp_max_mem_threshold"`
}

type SupportedTempRange struct {
	GpuTargetTempMin string `xml:"gpu_target_temp_min"`
	GpuTargetTempMax string `xml:"gpu_target_temp_max"`
}

type PowerReadings struct {
	PowerState          string `xml:"power_state"`
	AveragePowerDraw    string `xml:"average_power_draw"`
	InstantPowerDraw    string `xml:"instant_power_draw"`
	CurrentPowerLimit   string `xml:"current_power_limit"`
	RequestedPowerLimit string `xml:"requested_power_limit"`
	DefaultPowerLimit   string `xml:"default_power_limit"`
	MinPowerLimit       string `xml:"min_power_limit"`
	MaxPowerLimit       string `xml:"max_power_limit"`
}

type SimplePowerReadings struct {
	AveragePowerDraw string `xml:"average_power_draw"`
	InstantPowerDraw string `xml:"instant_power_draw"`
}

type ModulePowerReadings struct {
	PowerState          string `xml:"power_state"`
	AveragePowerDraw    string `xml:"average_power_draw"`
	InstantPowerDraw    string `xml:"instant_power_draw"`
	CurrentPowerLimit   string `xml:"current_power_limit"`
	RequestedPowerLimit string `xml:"requested_power_limit"`
	DefaultPowerLimit   string `xml:"default_power_limit"`
	MinPowerLimit       string `xml:"min_power_limit"`
	MaxPowerLimit       string `xml:"max_power_limit"`
}

type PowerProfiles struct {
	RequestedProfiles string `xml:"power_profile_requested_profiles"`
	EnforcedProfiles  string `xml:"power_profile_enforced_profiles"`
}

type Clocks struct {
	GraphicsClock string `xml:"graphics_clock"`
	SmClock       string `xml:"sm_clock"`
	MemClock      string `xml:"mem_clock"`
	VideoClock    string `xml:"video_clock"`
}

type ClockPair struct {
	GraphicsClock string `xml:"graphics_clock"`
	MemClock      string `xml:"mem_clock"`
}

type DeferredClocks struct {
	MemClock string `xml:"mem_clock"`
}

type MaxClocks struct {
	GraphicsClock string `xml:"graphics_clock"`
	SmClock       string `xml:"sm_clock"`
	MemClock      string `xml:"mem_clock"`
	VideoClock    string `xml:"video_clock"`
}

type MaxCustomerBoostClock struct {
	GraphicsClock string `xml:"graphics_clock"`
}

type ClockPolicy struct {
	AutoBoost        string `xml:"auto_boost"`
	AutoBoostDefault string `xml:"auto_boost_default"`
}

type Voltage struct {
	GraphicsVolt string `xml:"graphics_volt"`
}

type Fabric struct {
	State       string       `xml:"state"`
	Status      string       `xml:"status"`
	CliqueId    string       `xml:"cliqueId"`
	ClusterUuid string       `xml:"clusterUuid"`
	Health      FabricHealth `xml:"health"`
}

type FabricHealth struct {
	Bandwidth               string `xml:"bandwidth"`
	RouteRecoveryInProgress string `xml:"route_recovery_in_progress"`
	RouteUnhealthy          string `xml:"route_unhealthy"`
	AccessTimeoutRecovery   string `xml:"access_timeout_recovery"`
}

type SupportedClocks struct {
	SupportedMemClocks []SupportedMemClock `xml:"supported_mem_clock"`
}

type SupportedMemClock struct {
	Value                   string   `xml:"value"`
	SupportedGraphicsClocks []string `xml:"supported_graphics_clock"`
}

type Capabilities struct {
	Egm string `xml:"egm"`
}
