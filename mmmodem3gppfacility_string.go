// Code generated by "stringer -type=MMModem3gppFacility"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModem3gppFacilityNone-0]
	_ = x[MmModem3gppFacilitySim-1]
	_ = x[MmModem3gppFacilityFixedDialing-2]
	_ = x[MmModem3gppFacilityPhSim-4]
	_ = x[MmModem3gppFacilityPhFsim-8]
	_ = x[MmModem3gppFacilityNetPers-16]
	_ = x[MmModem3gppFacilityNetSubPers-32]
	_ = x[MmModem3gppFacilityProviderPers-64]
	_ = x[MmModem3gppFacilityCorpPers-128]
}

const (
	_MMModem3gppFacility_name_0 = "MmModem3gppFacilityNoneMmModem3gppFacilitySimMmModem3gppFacilityFixedDialing"
	_MMModem3gppFacility_name_1 = "MmModem3gppFacilityPhSim"
	_MMModem3gppFacility_name_2 = "MmModem3gppFacilityPhFsim"
	_MMModem3gppFacility_name_3 = "MmModem3gppFacilityNetPers"
	_MMModem3gppFacility_name_4 = "MmModem3gppFacilityNetSubPers"
	_MMModem3gppFacility_name_5 = "MmModem3gppFacilityProviderPers"
	_MMModem3gppFacility_name_6 = "MmModem3gppFacilityCorpPers"
)

var (
	_MMModem3gppFacility_index_0 = [...]uint8{0, 23, 45, 76}
)

func (i MMModem3gppFacility) String() string {
	switch {
	case i <= 2:
		return _MMModem3gppFacility_name_0[_MMModem3gppFacility_index_0[i]:_MMModem3gppFacility_index_0[i+1]]
	case i == 4:
		return _MMModem3gppFacility_name_1
	case i == 8:
		return _MMModem3gppFacility_name_2
	case i == 16:
		return _MMModem3gppFacility_name_3
	case i == 32:
		return _MMModem3gppFacility_name_4
	case i == 64:
		return _MMModem3gppFacility_name_5
	case i == 128:
		return _MMModem3gppFacility_name_6
	default:
		return "MMModem3gppFacility(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
