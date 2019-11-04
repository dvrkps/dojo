package randomservice

import (
	"testing"

	"github.com/dvrkps/dojo/consul/guard"
	"github.com/dvrkps/dojo/consul/logger"
)

func TestConfigurationErr(t *testing.T) {
	tests := []struct {
		ok     bool
		label  string
		config Configuration
	}{
		{
			ok:    true,
			label: "all good",
			config: Configuration{
				Guard:      &guard.Guard{},
				Log:        &logger.Logger{},
				NoWorkers:  1,
				MaxNumber:  1,
				OutputSize: 1,
			},
		},
		{
			label:  "nil guard",
			config: Configuration{},
		},
		{
			label: "nil logger",
			config: Configuration{
				Guard: &guard.Guard{},
				Log:   nil,
			},
		},
		{
			label: "workers < 1",
			config: Configuration{
				Guard:     &guard.Guard{},
				Log:       &logger.Logger{},
				NoWorkers: 0,
			},
		},
		{
			label: "max number < 1",
			config: Configuration{
				Guard:     &guard.Guard{},
				Log:       &logger.Logger{},
				NoWorkers: 1,
				MaxNumber: 0,
			},
		},
		{
			label: "output size < 1",
			config: Configuration{
				Guard:      &guard.Guard{},
				Log:        &logger.Logger{},
				NoWorkers:  1,
				MaxNumber:  1,
				OutputSize: 0,
			},
		},
	}
	for _, tt := range tests {
		err := tt.config.Err()
		if !tt.ok {
			if err == nil {
				t.Errorf("%s: Err() = <nil>; want <error>", tt.label)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: Err() = %v; want <nil>", tt.label, err)
		}
	}
}
