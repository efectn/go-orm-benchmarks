package helper

import (
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

// from https://stackoverflow.com/questions/56616196/how-to-convert-camel-case-string-to-snake-case, needs cleaner solution
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getFuncName(function any) string {
	name := strings.Split(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), ".")
	straightName := strings.Split(name[len(name)-1], "-")[0]

	return ToSnakeCase(straightName)
}
