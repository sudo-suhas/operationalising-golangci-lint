package snippets

import (
	"errors"
	"fmt"
)

// START BROKEN // OMIT

type MyError struct { // HL
	Code string
	// ...
}

func (e *MyError) Error() string { // HL
	// ...
	return "my-error" // OMIT
}

func demo() error {
	if err := doSomething(); err != nil {
		var e MyError
		if errors.As(err, &e) { // special case // HL
			// ...
			return fmt.Errorf("demo: %s", e.Code) // OMIT
		}
		return fmt.Errorf("demo: %w", err)
	}

	// ... continue happy path
	// END BROKEN // OMIT
	return nil
}

func demoFixed() error {
	if err := doSomething(); err != nil {
		// START FIXED // OMIT
		var e *MyError // HL
		if errors.As(err, &e) {
			// ...
			// END FIXED // OMIT
			return fmt.Errorf("demo: %s", e.Code) // OMIT
		}
		return fmt.Errorf("demo: %w", err)
	}

	// ... continue happy path
	// END BROKEN // OMIT
	return nil
}

func doSomething() error { return nil }
