package main

import (
	"fmt"
	"net/url"
	"strings"
)

// ======================

func getPathParams(path string) []string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var params []string
	for _, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			params = append(params, part[1:len(part)-1])
		}
	}
	return params
}
// ================================

func main() {

	param := getPathParams("/api/{v1}/users/{1}")
	fmt.Println(param)

}
