package main

// Generate a list of shortcuts for HTTP status code errors
//
// Go through all possible HTTP status codes and write out the stub function when we get something non-zero

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var output string

func init() {
	flag.StringVar(&output, "output", "-", "Where to write output")
}

func main() {
	flag.Parse()

	var f *os.File
	if output == "-" {
		f = os.Stdout
	} else {
		var err error
		f, err = os.Create(output)
		if err != nil {
			fmt.Printf("Could not open %s: %s\n", output, err)
			os.Exit(1)
			return
		}
		defer f.Close()
	}

	fmt.Fprintln(f, "package httperror\n")
	fmt.Fprintln(f, "import \"net/http\"")

	for i := 0; i < 1000; i += 1 {
		name := http.StatusText(i)

		if name == "" {
			continue
		}
		// 418 is magic
		if i == 418 {
			name = "Teapot"
		} else if i == 203 {
			name = "NonAuthoritativeInfo"
		} else if i == 407 {
			name = "ProxyAuthRequired"
		}

		shortName := strings.Replace(name, " ", "", -1)
		shortName = strings.Replace(shortName, "-", "", -1)

		fmt.Fprintf(f, `
// HTTP %d %s
func New%s(message ...string) *HTTPError {
	return New(http.Status%s, message...)
}
`, i, name, shortName, shortName)

		fmt.Fprintf(f, `
// HTTP %d %s
func Wrap%s(err error) *HTTPError {
	return Wrap(http.Status%s, err)
}
`, i, name, shortName, shortName)

	}
}
