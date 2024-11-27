package utils

import (
	"github.com/spf13/cast"
	"reflect"
)

func DistinctU64(elements []uint64) []uint64 {
	var result []uint64
	switch reflect.TypeOf(elements).Kind() {
	case reflect.Slice:
		var encountered = make(map[interface{}]bool)
		s := reflect.ValueOf(elements)
		for i := 0; i < s.Len(); i++ {
			if !encountered[s.Index(i).Interface()] == true {
				encountered[s.Index(i).Interface()] = true
				result = append(result, cast.ToUint64(s.Index(i).Interface()))
			}
		}
		break
	default:
		return result
	}

	return result
}

func Distinct64(elements []int64) []int64 {
	var result []int64
	switch reflect.TypeOf(elements).Kind() {
	case reflect.Slice:
		var encountered = make(map[interface{}]bool)
		s := reflect.ValueOf(elements)
		for i := 0; i < s.Len(); i++ {
			if !encountered[s.Index(i).Interface()] == true {
				encountered[s.Index(i).Interface()] = true
				result = append(result, cast.ToInt64(s.Index(i).Interface()))
			}
		}
		break
	default:
		return result
	}

	return result
}

func DistinctStr(elements []string) []string {
	var result []string
	switch reflect.TypeOf(elements).Kind() {
	case reflect.Slice:
		var encountered = make(map[interface{}]bool)
		s := reflect.ValueOf(elements)
		for i := 0; i < s.Len(); i++ {
			if !encountered[s.Index(i).Interface()] == true {
				encountered[s.Index(i).Interface()] = true
				result = append(result, cast.ToString(s.Index(i).Interface()))
			}
		}
		break
	default:
		return result
	}

	return result
}

func RemoveSliceZero(s []string) []string {
	var ret []string
	for _, u := range s {
		if u != "" {
			ret = append(ret, u)
		}
	}
	return ret
}
