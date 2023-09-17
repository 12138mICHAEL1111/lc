package main

import (
	"strings"
)

func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	stack := []string{}
	for _, v := range list {
		switch v{
		case "":
			continue
		case ".":
			continue
		case "..":
			if len(stack) == 0 {
				continue
			}else{
				stack = stack[:len(stack) - 1]
			}
		default:
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}