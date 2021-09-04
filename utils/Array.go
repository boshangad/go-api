package utils

func InArrayWithString(raw string, raws []string) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if raw == checkRaw {
			return true
		}
	}
	return false
}

func InArrayWithInt(raw int, raws []int) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if raw == checkRaw {
			return true
		}
	}
	return false
}

func InArrayWithInt8(raw int8, raws []int8) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if raw == checkRaw {
			return true
		}
	}
	return false
}

func InArrayWithInt16(raw int16, raws []int16) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if raw == checkRaw {
			return true
		}
	}
	return false
}

func InArrayWithInt32(raw int32, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(int32); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithInt64(raw int64, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(int64); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithUint(raw uint, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(uint); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithUint8(raw uint8, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(uint8); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithUint16(raw uint16, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(uint16); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithUint32(raw uint32, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(uint32); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}

func InArrayWithUint64(raw uint64, raws []interface{}) bool {
	if raws == nil || len(raws) < 1 {
		return false
	}
	for _, checkRaw := range raws {
		if data, ok := checkRaw.(uint64); ok {
			if raw == data {
				return true
			}
		}
	}
	return false
}
