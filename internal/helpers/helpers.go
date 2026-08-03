package helpers

import "os"

func GetEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
