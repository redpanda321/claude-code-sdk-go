package adapters

import (
	"errors"
	"testing"
)

type noopAdapter struct{}

func (noopAdapter) Framework() FrameworkInfo {
	return FrameworkInfo{Name: "noop", ImportPath: "test/noop", Version: "v0.0.0"}
}
func (noopAdapter) Wrap(src any) (any, error)   { return src, nil }
func (noopAdapter) Unwrap(src any) (any, error) { return src, nil }

var _ Adapter = (*noopAdapter)(nil)

func TestNoopAdapter_RoundTrip(t *testing.T) {
	a := noopAdapter{}
	in := "hello"
	wrapped, err := a.Wrap(in)
	if err != nil {
		t.Fatalf("Wrap: %v", err)
	}
	unwrapped, err := a.Unwrap(wrapped)
	if err != nil {
		t.Fatalf("Unwrap: %v", err)
	}
	if unwrapped != in {
		t.Fatalf("roundtrip = %v, want %v", unwrapped, in)
	}
}

func TestErrNoConcreteAdapter_Identity(t *testing.T) {
	if !errors.Is(ErrNoConcreteAdapter, ErrNoConcreteAdapter) {
		t.Fatal("errors.Is failed for sentinel")
	}
	if ErrNoConcreteAdapter.Error() == "" {
		t.Fatal("ErrNoConcreteAdapter.Error() returned empty string")
	}
}
