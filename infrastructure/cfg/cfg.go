package cfg

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var instance = make(Stage)

type Stage map[string]any
type Config struct{}

func (c Config) Init() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	b, err := lookupConfigFile(cwd)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &instance)
	if err != nil {
		return err
	}
	return nil
}

func lookupConfigFile(cwd string) ([]byte, error) {
	configPath := path.Join(cwd, "config.yaml")
	b, _ := os.ReadFile(configPath)
	if len(b) > 0 {
		return b, nil
	}
	configPath = path.Join(cwd, "config", "config.yaml")
	b, _ = os.ReadFile(configPath)
	if len(b) > 0 {
		return b, nil
	}
	cwd = filepath.Dir(cwd)
	configPath = path.Join(cwd, "config", "config.yaml")
	b, _ = os.ReadFile(configPath)
	if len(b) > 0 {
		return b, nil
	}
	return nil, fmt.Errorf("config file not found")
}

func GetString(pattern string) string {
	v, ok := lookup(instance, pattern).(string)
	if ok {
		return v
	}
	return ""
}

func GetInt(pattern string) int64 {
	v, ok := lookup(instance, pattern).(int)
	if ok {
		return int64(v)
	}
	return 0
}

func GetBool(pattern string) bool {
	v, ok := lookup(instance, pattern).(bool)
	if ok {
		return v
	}
	return false
}

func GetFloat(pattern string) float64 {
	v, ok := lookup(instance, pattern).(float64)
	if ok {
		return v
	}
	return 0
}

func GetInts(pattern string) []int {
	v, ok := lookup(instance, pattern).([]interface{})
	if ok {
		var ints []int
		for _, i := range v {
			if d, ok := i.(int); ok {
				ints = append(ints, d)
			}
		}
		return ints
	}
	return []int{}
}

func GetStrings(pattern string) []string {
	v, ok := lookup(instance, pattern).([]interface{})
	if ok {
		var sts []string
		for _, i := range v {
			if s, ok := i.(string); ok {
				sts = append(sts, s)
			}
		}
		return sts
	}
	return []string{}
}

func lookup(stage Stage, pattern string) any {
	parts := strings.Split(pattern, ".")
	for idx, part := range parts {
		v, ok := stage[part]
		if !ok {
			return nil
		}
		switch v.(type) {
		case Stage:
			newPattern := strings.Join(parts[idx+1:], ".")
			return lookup(v.(Stage), newPattern)
		default:
			return v
		}
	}
	return nil
}
