package testutil

import (
	"gopkg.in/yaml.v3"
	"os"
)

type RunnScenario struct {
	Runners map[string]string `yaml:"runners"`
}

// GetRunnerKeys parses a runn scenario YAML file and returns the keys under 'runners'.
func GetRunnerKeys(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var scenario RunnScenario
	if err := yaml.Unmarshal(data, &scenario); err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(scenario.Runners))
	for k := range scenario.Runners {
		keys = append(keys, k)
	}
	return keys, nil
}
