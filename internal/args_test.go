package internal

import (
	"testing"
)

func TestVerboseIsBoolOption(t *testing.T) {

	args := []string{"-verbose"}
	cmd, err := Read(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !cmd.verbose {
		t.Error("verbose should be true when -verbose is passed")
	}
}

func TestVerboseIsOptional(t *testing.T) {
	args := []string{}
	cmd, err := Read(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if cmd.verbose {
		t.Error("verbose should be false when -verbose is not passed")
	}
}
