package helper

import (
	"reflect"
	"runtime"
	"strings"
)

func getFuncName(function any) string {
	name := strings.Split(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), ".")
	straightName := strings.Split(name[len(name)-1], "-")[0]

	return straightName
}
