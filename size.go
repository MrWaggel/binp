package binp

// Size calculates the total capacity needed to hold all the values.
func Size(values []interface{}) (int, error) {
	var total int
	for _, v := range values {
		switch v.(type) {
		case int:
			total += 9
		case []int:
			total += len(v.([]int))*8 + 3 //  b[0] slice identifier + b[1,2] slice len
		case int64:
			total += 9
		case []int64:
			total += len(v.([]int64))*8 + 3 //  b[0] slice identifier + b[1,2] slice len
		case int32:
			total += 5
		case []int32:
			total += len(v.([]int32))*4 + 3 //  b[0] slice identifier + b[1,2] slice len
		case int16:
			total += 3
		case []int16:
			total += len(v.([]int16))*2 + 3 //  b[0] slice identifier + b[1,2] slice len
		case int8:
			total += 2
		case []int8:
			total += len(v.([]int8))*1 + 3 //  b[0] slice identifier + b[1,2] slice len
		case uint:
			total += 9
		case []uint:
			total += len(v.([]uint))*8 + 3
		case uint64:
			total += 9
		case []uint64:
			total += len(v.([]uint64))*8 + 3 //  b[0] slice identifier + b[1,2] slice len
		case uint32:
			total += 5
		case []uint32:
			total += len(v.([]uint32))*4 + 3 //  b[0] slice identifier + b[1,2] slice len
		case uint16:
			total += 3
		case []uint16:
			total += len(v.([]uint16))*2 + 3 //  b[0] slice identifier + b[1,2] slice len
		case byte:
			total += 2
		case []byte:
			total += len(v.([]byte))*1 + 5 //  b[0] slice identifier + b[1,2] slice len
		case bool:
			total += 1
		case []bool:
			total += len(v.([]bool))*1 + 3 //  b[0] slice identifier + b[1,2] slice len
		case float64:
			total += 9
		case []float64:
			total += len(v.([]float64))*8 + 3 //  b[0] slice identifier + b[1,2] slice len
		case float32:
			total += 5
		case []float32:
			total += len(v.([]float32))*4 + 3 //  b[0] slice identifier + b[1,2] slice len
		case string:
			total += 5 + len(v.(string))
		case []string:
			total += 3 // String Slice identifier + 2 bytes for slice len
			for _, e := range v.([]string) {
				total += len(e) + 4 // 4 is the string length identifier
			}
		default:
			return 0, ErrUnsupportedType

		}
	}
	return total, nil
}
