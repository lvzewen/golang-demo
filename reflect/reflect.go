package main

import (
	"reflect"
	"fmt"
)

func main() {
	fmt.Printf("contains %d", Contains([]string{"1", "2", "3"}, "2"))
	fmt.Printf("string contains %d", StringContains([]string{"1", "2", "3"}, "2"))
}

// Contains ...
func Contains(arr interface{}, key interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(arr).Kind() {
		case reflect.Slice: {
			s := reflect.ValueOf(arr)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(key, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// StringContains ...
func StringContains(arr []string, val string) (index int) {
	index = -1
	
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			index = i
			return
		}
	}
	return index
}
