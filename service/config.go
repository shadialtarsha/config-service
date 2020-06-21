package service

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Config is struct that defines the underlying data of the config.
type Config struct {
	config map[string]interface{}
}

// SetFromBytes read the raw config from an input.
func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}

	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("config is not a map")
	}

	config, err := convertKeysToStrings(untypedConfig)
	if err != nil {
		return err
	}

	c.config = config

	return nil
}

// Get returns the specific config for the given service and a given env.
func (c *Config) Get(env string, serviceName string) (map[string]interface{}, error) {
	envConfig, ok := c.config[env].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("env config is not a map")
	}

	a, ok := envConfig["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config is not a map")
	}

	// If no config is defined for the service, return the base config.
	if _, ok = envConfig[serviceName]; !ok {
		return a, nil
	}

	b, ok := envConfig[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q config is not a map", serviceName)
	}

	// Merge the service config with the base config.
	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}

	return config, nil
}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToStrings(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil
}
