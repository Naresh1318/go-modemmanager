// Code generated by "stringer -type=MMSmsPduType"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsPduTypeUnknown-0]
	_ = x[MmSmsPduTypeDeliver-1]
	_ = x[MmSmsPduTypeSubmit-2]
	_ = x[MmSmsPduTypeStatusReport-3]
	_ = x[MmSmsPduTypeCdmaDeliver-32]
	_ = x[MmSmsPduTypeCdmaSubmit-33]
	_ = x[MmSmsPduTypeCdmaCancellation-34]
	_ = x[MmSmsPduTypeCdmaDeliveryAcknowledgement-35]
	_ = x[MmSmsPduTypeCdmaUserAcknowledgement-36]
	_ = x[MmSmsPduTypeCdmaReadAcknowledgement-37]
}

const (
	_MMSmsPduType_name_0 = "MmSmsPduTypeUnknownMmSmsPduTypeDeliverMmSmsPduTypeSubmitMmSmsPduTypeStatusReport"
	_MMSmsPduType_name_1 = "MmSmsPduTypeCdmaDeliverMmSmsPduTypeCdmaSubmitMmSmsPduTypeCdmaCancellationMmSmsPduTypeCdmaDeliveryAcknowledgementMmSmsPduTypeCdmaUserAcknowledgementMmSmsPduTypeCdmaReadAcknowledgement"
)

var (
	_MMSmsPduType_index_0 = [...]uint8{0, 19, 38, 56, 80}
	_MMSmsPduType_index_1 = [...]uint8{0, 23, 45, 73, 112, 147, 182}
)

func (i MMSmsPduType) String() string {
	switch {
	case i <= 3:
		return _MMSmsPduType_name_0[_MMSmsPduType_index_0[i]:_MMSmsPduType_index_0[i+1]]
	case 32 <= i && i <= 37:
		i -= 32
		return _MMSmsPduType_name_1[_MMSmsPduType_index_1[i]:_MMSmsPduType_index_1[i+1]]
	default:
		return "MMSmsPduType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
