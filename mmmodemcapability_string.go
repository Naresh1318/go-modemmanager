// Code generated by "stringer -type=MMModemCapability"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModemCapabilityNone-0]
	_ = x[MmModemCapabilityPots-1]
	_ = x[MmModemCapabilityCdmaEvdo-2]
	_ = x[MmModemCapabilityGsmUmts-4]
	_ = x[MmModemCapabilityLte-8]
	_ = x[MmModemCapabilityLteAdvanced-16]
	_ = x[MmModemCapabilityIridium-32]
	_ = x[MmModemCapabilityAny-4294967295]
}

const (
	_MMModemCapability_name_0 = "MmModemCapabilityNoneMmModemCapabilityPotsMmModemCapabilityCdmaEvdo"
	_MMModemCapability_name_1 = "MmModemCapabilityGsmUmts"
	_MMModemCapability_name_2 = "MmModemCapabilityLte"
	_MMModemCapability_name_3 = "MmModemCapabilityLteAdvanced"
	_MMModemCapability_name_4 = "MmModemCapabilityIridium"
	_MMModemCapability_name_5 = "MmModemCapabilityAny"
)

var (
	_MMModemCapability_index_0 = [...]uint8{0, 21, 42, 67}
)

func (i MMModemCapability) String() string {
	switch {
	case i <= 2:
		return _MMModemCapability_name_0[_MMModemCapability_index_0[i]:_MMModemCapability_index_0[i+1]]
	case i == 4:
		return _MMModemCapability_name_1
	case i == 8:
		return _MMModemCapability_name_2
	case i == 16:
		return _MMModemCapability_name_3
	case i == 32:
		return _MMModemCapability_name_4
	case i == 4294967295:
		return _MMModemCapability_name_5
	default:
		return "MMModemCapability(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
