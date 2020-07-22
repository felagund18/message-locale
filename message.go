package message_locale

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var messages map[string]string

func LoadResource(path string, code string) {
	messages = map[string]string{}

	content, err := ioutil.ReadFile(path + "/" + code + ".lang")
	if err != nil {
		log.Println(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}

		t := strings.Split(line, "=")
		messages[strings.TrimSpace(t[0])] = strings.TrimSpace(t[1])
	}
}

func Format(code string, val ...interface{}) string {
	message, ok := messages[code]
	if !ok {
		return ""
	}
	if len(val) == 0 {
		return message
	}
	return fmt.Sprintf(message, val...)
}
