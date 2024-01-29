package utils

import (
	"os"
)

func EnvBool(name string) bool {
	val, set := os.LookupEnv(name)
	if set {
		return val == "YES"
	} else {
		return false
	}
}

func EnvString(name string) string {
	val, set := os.LookupEnv(name)
	if set {
		return val
	} else {
		panic("Environment variable " + name + " not set")
	}
}
