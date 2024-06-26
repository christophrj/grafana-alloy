package triton

import (
	"testing"

	"github.com/grafana/alloy/syntax"
	"github.com/stretchr/testify/require"
)

func TestAlloyConfig(t *testing.T) {
	var exampleAlloyConfig = `
	account = "TRITON-987654"
	dns_suffix = "triton.example.com"
	endpoint = "triton.example.com"
	groups = ["group1", "group2"]
	tls_config {
		ca_file = "/path/to/ca_file"
	}
`

	var args Arguments
	err := syntax.Unmarshal([]byte(exampleAlloyConfig), &args)
	require.NoError(t, err)
}

func TestBadAlloyConfig(t *testing.T) {
	var exampleAlloyConfig = `
	account = "TRITON-987654"
	dns_suffix = "triton.example.com"
	endpoint = "triton.example.com"
	groups = ["group1", "group2"]
	tls_config {
		ca_file = "/path/to/ca_file"
		ca_pem = "not a real pem"
	}
`

	// Make sure the TLSConfig Validate function is being utilized correctly
	var args Arguments
	err := syntax.Unmarshal([]byte(exampleAlloyConfig), &args)
	require.ErrorContains(t, err, "at most one of ca_pem and ca_file must be configured")
}
