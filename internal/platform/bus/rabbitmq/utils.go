package rabbitmq

import (
	"regexp"
	"strings"

	"github.com/rfdez/voting-poll/kit/event"
)

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getServiceName(eventType event.Type) string {
	splited := strings.Split(string(eventType), ".")
	return splited[1]
}
