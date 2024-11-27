package jsonx

import (
	json "github.com/json-iterator/go"
	"reflect"
	"unsafe"
)

func ToStr(obj interface{}) string {
	js, _ := json.Marshal(obj)
	return Bytes2str(js)
}
func ToBytes(obj interface{}) []byte {
	js, _ := json.Marshal(obj)
	return js
}

func ToMap(obj string) map[string]interface{} {
	ret := make(map[string]interface{})
	if err := json.Unmarshal(Str2byt(obj), &ret); err != nil {
		return nil
	}
	return ret
}

func ToSlice(obj string) []interface{} {
	var ret []interface{}
	if err := json.Unmarshal(Str2byt(obj), &ret); err != nil {
		return nil
	}
	return ret
}

func Bytes2str(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
}

func Str2byt(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func CopyStruct(src, dst interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonStr, dst); err != nil {
		return err
	}

	return nil
}
