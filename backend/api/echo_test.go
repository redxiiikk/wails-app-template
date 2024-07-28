package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoApi(t *testing.T) {
	type test struct {
		name     string
		arg     string
		expected string
	}

	tests := []test{
		{
			name: "should get same message",
			arg: "Echo Message",
			expected: "Echo Message",
		},
	}

	echoApi := createEchoApiInstance()

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			response, err := echoApi.echo(EchoRequest{Message: tc.arg})
			assert.Nil(tt, err)
			assert.Equal(tt, tc.expected, response.Message)
		})
	}
}

func createEchoApiInstance() *EchoApi {
	return &EchoApi{}
}