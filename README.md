# kerrors

A simple Go package for standardized error handling using three focused error types: 

- **ValueError**: For invalid input, type errors, and validation issues
- **SystemError**: For system-level problems like I/O failures, crashes, or internal bugs
- **NetworkError**: For failures related to communication, timeouts, or connectivity

This library helps you simplify error propagation and categorization by wrapping underlying errors into one of three high-level categories

## Setup

There are 2 options for setup:

<details><summary>1. Using `go mod tidy`</summary>

Initialize your project with `go mod init <name>`. Then import it in your code and use it:

```go
import (
	"os"

    "github.com/descent098/kerrors"
)

func readFile() error {
	_, err := os.ReadFile("/nonexistent.txt")
	if err != nil {
		return kerrors.NewSystemError("Failed to read file", err)
	}
	return nil
}
```

Then run `go mod tidy` to download.

</details>

<details><summary>2. Using `go install`</summary>

Install the package using `go install`:

```bash
go install github.com/descent098/kerrors
```

Import it in your code:

```go
import "github.com/descent098/kerrors"
```

</details>

## Usage

### ValueError

```go
package main

import (
	"errors"
	"fmt"

	"github.com/descent098/kerrors"
)

func processInput(input string) error {
	if input == "" {
		baseErr := errors.New("input is blank")
		return kerrors.NewValueError("Input cannot be empty", baseErr)
	}
	return nil
}

func main() {
	err := processInput("")
	if err != nil {
		fmt.Println("Error:", err)

		// Check if it's a ValueError
		var ve *kerrors.ValueError
		if errors.As(err, &ve) {
			fmt.Println("Caught a ValueError")
		}

		// Check if the base error is wrapped
		if errors.Is(err, errors.New("input is blank")) {
			fmt.Println("Underlying error matched")
		}

		// Unwrap and inspect
		fmt.Println("Wrapped error:", errors.Unwrap(ve))
	}
}
```

### SystemError

```go
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/descent098/kerrors"
)

func readFile() error {
	_, err := os.ReadFile("/nonexistent.txt")
	if err != nil {
		return kerrors.NewSystemError("Failed to read file", err)
	}
	return nil
}

func main() {
	err := readFile()
	if err != nil {
		fmt.Println("Error:", err)

		// Check type
		var se *kerrors.SystemError
		if errors.As(err, &se) {
			fmt.Println("Caught a SystemError")
		}

		// Check wrapped os error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist (wrapped)")
		}

		// Unwrap
		fmt.Println("Wrapped error:", errors.Unwrap(se))
	}
}
```

### NetworkError

```go
package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/descent098/kerrors"
)

func connect() error {
	underlying := &net.DNSError{Err: "connection timed out", Name: "example.com"}
	return kerrors.NewNetworkError("Unable to connect to server", underlying)
}

func main() {
	err := connect()
	if err != nil {
		fmt.Println("Error:", err)

		// Check if it's a NetworkError
		var ne *kerrors.NetworkError
		if errors.As(err, &ne) {
			fmt.Println("Caught a NetworkError")
		}

		// Check for wrapped DNS error
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			fmt.Println("Underlying DNS error:", dnsErr)
		}

		// Unwrap
		fmt.Println("Wrapped error:", errors.Unwrap(ne))
	}
}
```

