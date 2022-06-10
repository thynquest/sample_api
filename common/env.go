package common

import "os"

func GetEnv(key, defvalue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defvalue
	}
	return value
}
