package src

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"main/data"
)

func LoadConfig(path string) (data.Config, error) {

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return data.Config{}, fmt.Errorf("src/confg/LoadConfig()|Error reading file %q|%v", path, err)
	}

	var config data.Config
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return data.Config{}, fmt.Errorf("src/confg/LoadConfig()|Error unmarshalling config|%v", err)
	}

	return config, nil
}
