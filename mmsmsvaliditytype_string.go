// Code generated by "stringer -type=MMSmsValidityType"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsValidityTypeUnknown-0]
	_ = x[MmSmsValidityTypeRelative-1]
	_ = x[MmSmsValidityTypeAbsolute-2]
	_ = x[MmSmsValidityTypeEnhanced-3]
}

const _MMSmsValidityType_name = "MmSmsValidityTypeUnknownMmSmsValidityTypeRelativeMmSmsValidityTypeAbsoluteMmSmsValidityTypeEnhanced"

var _MMSmsValidityType_index = [...]uint8{0, 24, 49, 74, 99}

func (i MMSmsValidityType) String() string {
	if i >= MMSmsValidityType(len(_MMSmsValidityType_index)-1) {
		return "MMSmsValidityType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMSmsValidityType_name[_MMSmsValidityType_index[i]:_MMSmsValidityType_index[i+1]]
}
