package main

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

type GenConfig password.GeneratorInput

func GenPassword(c *GenConfig) (string, error) {
	gen, err := password.NewGenerator((*password.GeneratorInput)(c))
	if err != nil {
		return "", err
	}
	return gen.Generate(10, 0, 0, false, true)
}

func GenPasswordDefault() string {
	p, _ := GenPassword(&GenConfig{})
	return p
}

func GetPassword(key string) string {
	return fmt.Sprintf("%s password", key)
}

func AddPassword(name, url, password string) {

	passwords[name] = Password{
		Name:     name,
		URL:      url,
		Password: password,
	}
}
