package tezosconnect

import (
	"path/filepath"
	"testing"
)

func TestWriteConfig(t *testing.T) {
	dir := t.TempDir()
	configFilename := filepath.Join(dir + "/config.yaml")
	extraTezosConfigPath := filepath.Join(dir + "/conflate/extra.yaml")
	p := Config{}
	t.Run("TestWriteConfig", func(t *testing.T) {
		err := p.WriteConfig(configFilename, extraTezosConfigPath)
		if err != nil {
			t.Logf("unable to write to config files: %v", err)
		}
	})
}
