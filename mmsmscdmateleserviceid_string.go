// Code generated by "stringer -type=MMSmsCdmaTeleserviceId"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsCdmaTeleserviceIdUnknown-0]
	_ = x[MmSmsCdmaTeleserviceIdCmt91-4096]
	_ = x[MmSmsCdmaTeleserviceIdWpt-4097]
	_ = x[MmSmsCdmaTeleserviceIdWmt-4098]
	_ = x[MmSmsCdmaTeleserviceIdVmn-4099]
	_ = x[MmSmsCdmaTeleserviceIdWap-4100]
	_ = x[MmSmsCdmaTeleserviceIdWemt-4101]
	_ = x[MmSmsCdmaTeleserviceIdScpt-4102]
	_ = x[MmSmsCdmaTeleserviceIdCatpt-4103]
}

const (
	_MMSmsCdmaTeleserviceId_name_0 = "MmSmsCdmaTeleserviceIdUnknown"
	_MMSmsCdmaTeleserviceId_name_1 = "MmSmsCdmaTeleserviceIdCmt91MmSmsCdmaTeleserviceIdWptMmSmsCdmaTeleserviceIdWmtMmSmsCdmaTeleserviceIdVmnMmSmsCdmaTeleserviceIdWapMmSmsCdmaTeleserviceIdWemtMmSmsCdmaTeleserviceIdScptMmSmsCdmaTeleserviceIdCatpt"
)

var (
	_MMSmsCdmaTeleserviceId_index_1 = [...]uint8{0, 27, 52, 77, 102, 127, 153, 179, 206}
)

func (i MMSmsCdmaTeleserviceId) String() string {
	switch {
	case i == 0:
		return _MMSmsCdmaTeleserviceId_name_0
	case 4096 <= i && i <= 4103:
		i -= 4096
		return _MMSmsCdmaTeleserviceId_name_1[_MMSmsCdmaTeleserviceId_index_1[i]:_MMSmsCdmaTeleserviceId_index_1[i+1]]
	default:
		return "MMSmsCdmaTeleserviceId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
