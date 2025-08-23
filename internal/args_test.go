package internal

import (
	"testing"
)

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

func TestCommandLineArgumentSyntax(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		args     []string
		expected Args
	}{
		{
			name: "verbose flag is acceptable",
			args: []string{"-verbose", "-token", "a"},
			expected: Args{
				verbose: true,
				token:   "a",
			},
		},
		{
			name: "verbose flag is optional",
			args: []string{"-token", "a"},
			expected: Args{
				verbose: false,
				token:   "a",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		})
	}

}
