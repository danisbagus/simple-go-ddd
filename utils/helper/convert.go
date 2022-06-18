package helper

import "strconv"

func StringToInt64(sValue string, defValue int64) int64 {
	value := defValue
	iValue, err := strconv.ParseInt(sValue, 0, 32)
	if err == nil {
		value = iValue
	}
	return value
}

func StringToUint64(sValue string, defValue uint64) uint64 {
	value := StringToInt64(sValue, int64(defValue))
	return uint64(value)
}
