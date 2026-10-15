package greeting

import "fmt"

// Message returns a greeting for the supplied name.
func Message(name string) string {
	if name == "" {
		name = "DevOps Engineer"
	}

	return fmt.Sprintf("Hello, %s! Your Go CI pipeline is working.", name)
}
