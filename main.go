package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/idtoken"
)

func main() {
	idToken := os.Getenv("ID_TOKEN")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientId)
	if err != nil {
		fmt.Println("validate err:", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
