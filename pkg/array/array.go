package array

import "reflect"

func In(val interface{}, arr interface{}) int {
	values := reflect.ValueOf(arr)
	if reflect.TypeOf(arr).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}
