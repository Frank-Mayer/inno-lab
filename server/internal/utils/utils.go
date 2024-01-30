package utils

import (
	"os"
	"strconv"
)

func EnvBool(name string) bool {
	val, set := os.LookupEnv(name)
	if set {
		return val == "YES" || val == "yes" || val == "true" || val == "TRUE" || val == "1" || val == "y" || val == "Y"
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

func EnvInt(name string) int {
	val := EnvString(name)
	if i, err := strconv.Atoi(val); err == nil {
		return i
	} else {
		panic("Environment variable " + name + " not set to an integer")
	}
}
