package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name           string
		header         http.Header
		expected       string
		expectedErrMsg string
		isErrExpected  bool
	}

	tests := []test{
		{
			name: "happy flow should pass",
			header: http.Header{
				"Authorization": {"ApiKey 123"},
			},
			expected:       "1234",
			expectedErrMsg: "",
			isErrExpected:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetAPIKey(test.header)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("expected error, didn't get one")
					return
				}
				if err.Error() != test.expectedErrMsg {
					t.Errorf("expected %s\ngot %s",
						test.expectedErrMsg, err.Error())
				}

			default:
				if err != nil {
					t.Errorf("didn't expect error, got err %v\n",
						err.Error())
				}
				if got != test.expected {
					t.Errorf("expected %s\ngot %s",
						test.expected, got)
				}

			}

		})
	}

}
