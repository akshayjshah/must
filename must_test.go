package must

import (
	"fmt"
	"net/url"
)

var google = Must(url.Parse("https://www.google.com"))

func Example() {
	fmt.Println(google.Scheme)
	// output:
	// https
}
