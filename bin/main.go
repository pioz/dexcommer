package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pioz/dexcommer"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	username := os.Getenv("DEXCOMMER_USERNAME")
	password := os.Getenv("DEXCOMMER_PASSWORD")
	applicationId := os.Getenv("DEXCOMMER_APPLICATION_ID")
	if applicationId == "" {
		applicationId = "d89443d2-327c-4a6f-89e5-496bbb0317db"
	}
	fmt.Println(dexcommer.ReadLastestGlucoseValues(username, password, applicationId))
}
