package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
)

func main() {
	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(context.Background(), "forstmeier", nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	for _, org := range orgs {
		fmt.Printf("org: %+v\n", org)
	}
}
