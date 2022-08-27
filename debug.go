package collection

import "fmt"

// Debug returns debug format function that can be used in Debug method
func Debug[T any](name string, ff func(T) string) func(T) string {
	return func(item T) string {
		return fmt.Sprintf("%v:%v", name, ff(item))
	}
}
