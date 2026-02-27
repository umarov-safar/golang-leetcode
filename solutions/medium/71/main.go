package main

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	res := make([]string, 0)
	for _, p := range paths {
		if p == "" || p == "." {
			continue
		}

		if p == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			continue
		}

		res = append(res, p)
	}

	return "/" + strings.Join(res, "/")
}
