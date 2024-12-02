package tiger

import (
	"liquid/pkg/conf"
	"testing"
)

func TestSetupInstance(t *testing.T) {
	// expect not to throw errors
	conf.LoadTestConfig()
	SetupInstance()
}
