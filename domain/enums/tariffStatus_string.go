// Code generated by "stringer -type=TariffStatus -trimprefix=TariffStatus -output=tariffStatus_string.go"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TariffStatusEnabled-1]
	_ = x[TariffStatusDisabled-2]
}

const _TariffStatus_name = "EnabledDisabled"

var _TariffStatus_index = [...]uint8{0, 7, 15}

func (i TariffStatus) String() string {
	i -= 1
	if i >= TariffStatus(len(_TariffStatus_index)-1) {
		return "TariffStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TariffStatus_name[_TariffStatus_index[i]:_TariffStatus_index[i+1]]
}
