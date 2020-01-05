# gofuncmap

Use golang functions in template

## Usage

```go
package main

import (
	"html/template"
	"os"

	"github.com/gggwvg/gofuncmap"
)

func main() {
	var (
		user = map[string]string{"Name": "www"}
		txt  = `My name is {{ .Name | toUpper }}`
	)
	tmpl, err := template.New("example").Funcs(gofuncmap.Funcs).Parse(txt)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	// Output:
	// My name is WWW
}
```
