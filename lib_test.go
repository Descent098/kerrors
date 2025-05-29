package kerrors_test

import (
	"errors"
	"os"
	"testing"

	"github.com/descent098/kerrors"
)

func TestValueError(t *testing.T) {
	t.Run("wraps custom error", func(t *testing.T) {
		base := errors.New("invalid value")
		err := kerrors.NewValueError("Value processing failed", base)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "Value processing failed") || !containsSubstring(err.Error(), "invalid value") {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, base) {
			t.Errorf("expected errors.Is to match base error")
		}

		var ve *kerrors.ValueError
		if !errors.As(err, &ve) {
			t.Errorf("expected errors.As to match ValueError")
		}
	})

	t.Run("wraps nil error", func(t *testing.T) {
		err := kerrors.NewValueError("Missing input", nil)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "Missing input" {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if errors.Unwrap(err) != nil {
			t.Errorf("expected nil unwrap, got: %v", errors.Unwrap(err))
		}

		var ve *kerrors.ValueError
		if !errors.As(err, &ve) {
			t.Errorf("expected errors.As to match ValueError")
		}
	})

	t.Run("wraps standard library error", func(t *testing.T) {
		err := kerrors.NewValueError("File error", os.ErrNotExist)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "File error") || !containsSubstring(err.Error(), os.ErrNotExist.Error()) {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected errors.Is to match os.ErrNotExist")
		}

		var ve *kerrors.ValueError
		if !errors.As(err, &ve) {
			t.Errorf("expected errors.As to match ValueError")
		}
	})
}

func TestSystemError(t *testing.T) {
	t.Run("wraps custom error", func(t *testing.T) {
		base := errors.New("kernel panic")
		err := kerrors.NewSystemError("Critical failure", base)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "Critical failure") || !containsSubstring(err.Error(), "kernel panic") {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, base) {
			t.Errorf("expected errors.Is to match base error")
		}

		var se *kerrors.SystemError
		if !errors.As(err, &se) {
			t.Errorf("expected errors.As to match SystemError")
		}
	})

	t.Run("wraps nil error", func(t *testing.T) {
		err := kerrors.NewSystemError("Internal fault", nil)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "Internal fault" {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if errors.Unwrap(err) != nil {
			t.Errorf("expected nil unwrap, got: %v", errors.Unwrap(err))
		}

		var se *kerrors.SystemError
		if !errors.As(err, &se) {
			t.Errorf("expected errors.As to match SystemError")
		}
	})

	t.Run("wraps standard library error", func(t *testing.T) {
		err := kerrors.NewSystemError("OS error", os.ErrPermission)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "OS error") || !containsSubstring(err.Error(), os.ErrPermission.Error()) {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, os.ErrPermission) {
			t.Errorf("expected errors.Is to match os.ErrPermission")
		}

		var se *kerrors.SystemError
		if !errors.As(err, &se) {
			t.Errorf("expected errors.As to match SystemError")
		}
	})
}

func TestNetworkError(t *testing.T) {
	t.Run("wraps custom error", func(t *testing.T) {
		base := errors.New("timeout")
		err := kerrors.NewNetworkError("Connection issue", base)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "Connection issue") || !containsSubstring(err.Error(), "timeout") {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, base) {
			t.Errorf("expected errors.Is to match base error")
		}

		var ne *kerrors.NetworkError
		if !errors.As(err, &ne) {
			t.Errorf("expected errors.As to match NetworkError")
		}
	})

	t.Run("wraps nil error", func(t *testing.T) {
		err := kerrors.NewNetworkError("Socket closed", nil)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "Socket closed" {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if errors.Unwrap(err) != nil {
			t.Errorf("expected nil unwrap, got: %v", errors.Unwrap(err))
		}

		var ne *kerrors.NetworkError
		if !errors.As(err, &ne) {
			t.Errorf("expected errors.As to match NetworkError")
		}
	})

	t.Run("wraps standard library error", func(t *testing.T) {
		err := kerrors.NewNetworkError("Dial failed", os.ErrDeadlineExceeded)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !containsSubstring(err.Error(), "Dial failed") || !containsSubstring(err.Error(), os.ErrDeadlineExceeded.Error()) {
			t.Errorf("unexpected error message: %s", err.Error())
		}
		if !errors.Is(err, os.ErrDeadlineExceeded) {
			t.Errorf("expected errors.Is to match os.ErrDeadlineExceeded")
		}

		var ne *kerrors.NetworkError
		if !errors.As(err, &ne) {
			t.Errorf("expected errors.As to match NetworkError")
		}
	})
}

// Small helper for checking if substring is in string
func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[0:len(substr)] == substr || containsSubstring(s[1:], substr)))
}
