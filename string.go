package gofuncmap

import "strings"

var stringFuncs = map[string]interface{}{
	"toUpper": strings.ToUpper,
	"toLower": strings.ToLower,
}
