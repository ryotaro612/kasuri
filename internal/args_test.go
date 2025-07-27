package internal

import (
	"testing"
)

func TestVerboseIsBoolOption(t *testing.T) {
	t.Parallel()
	args := []string{"-verbose", "-token", "a"}
	actual, err := Read(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := Args{
		verbose: true,
		token:   "a",
	}

	if expected != actual {
		t.Errorf("The expected value is %#v, but got %#v", expected, actual)
	}

}

func TestVerboseIsOptional(t *testing.T) {
	t.Parallel()
	actual, err := Read([]string{"-token", "a"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := Args{
		verbose: false,
		token:   "a",
	}
	if expected != actual {
		t.Errorf("The expected value is %#v, but got %#v", expected, actual)
	}

}

func TestSlackTokenCanBePassedAsEnvVar(t *testing.T) {
	t.Setenv("ASHITABA_SLACK_TOKEN", "test_token")
	cmd, err := Read([]string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if cmd.token != "test_token" {
		t.Error("cmd.token should be set to the value of ASHITABA_SLACK_TOKEN")
	}
}
