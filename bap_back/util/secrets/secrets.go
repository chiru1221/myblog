package secrets

import (
	"io/ioutil"
	"log"
)

func SecretsFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
