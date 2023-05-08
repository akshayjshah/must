// Package must provides a generic equivalent of the standard library's
// [template.Must] or [regexp.MustCompile].
package must

// Must adapts a typical Go constructor to either panic or succeed. This is
// usually undesirable, but it's occasionally useful when creating global
// variables using a pure function and static inputs. Think of it as a generic
// version of the standard library's [template.Must] or [regexp.MustCompile].
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
