package gofuncmap

import (
	"html/template"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_stringFuncs(t *testing.T) {
	txt := `My name is {{ .Name | toUpper }}`
	tmpl, err := template.New("test_funcs").Funcs(Funcs).Parse(txt)
	require.Nil(t, err)
	err = tmpl.Execute(os.Stdout, map[string]string{"Name": "www"})
	assert.Nil(t, err)
}
