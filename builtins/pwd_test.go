//Leif Messinger
//Description: Tests the pwd command
//4-4-2023 CSCE4600
package builtins

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvironmentVariables(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		setEnv  map[string]string
		wantOut string
		wantErr error
	}{
		{
			name: "success no args",
			wantOut: fmt.Sprintln(os.Getwd(), "\n"),
		},
		{
			name: "success no args, root directory",
			setEnv: map[string]string{
				"PWD": "/"
			},
			wantOut: fmt.Sprintln("/", "\n")
		}
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// setup
			for k, v := range tt.setEnv {
				t.Setenv(k, v)
			}

			// test
			var out bytes.Buffer
			if err := PrintWorkingDirectory(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("EnvironmentVariables() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("EnvironmentVariables() unexpected error: %v", err)
			}
			if got := out.String(); got != tt.wantOut {
				t.Errorf("EnvironmentVariables() got = %v, want %v", got, tt.wantOut)
			}
		})
	}
}
