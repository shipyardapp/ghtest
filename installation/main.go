package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// func main() {
// 	// Wrap the shared transport for use with the integration ID 1 authenticating with installation ID 99.
// 	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, int64(0), int64(0), "forstmeierian.private_key.pem")
// 	if err != nil {
// 		// Handle error.
// 	}

// 	// Use installation transport with client.
// 	client := github.NewClient(&http.Client{Transport: itr})

// 	// opts := &github.RepositoryListOptions{}
// 	opts := &github.ListOptions{}

// 	ctx := context.Background()
// 	repos, _, err := client.Apps.ListUserRepos(ctx, int64(o), opts)
// 	if err != nil {
// 		log.Println("error:", err.Error())
// 	}
// 	for _, repo := range repos {
// 		log.Println("repo:", *repo.Name)
// 	}
// }

// func main() {
// 	httpClient := &http.Client{}

// 	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", nil)

// 	q := url.Values{}
// 	q.Add("client_id", "ID")
// 	q.Add("client_secret", "SECRET")
// 	q.Add("code", "CODE")

// 	req.URL.RawQuery = q.Encode()

// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		log.Println("error:", err.Error())
// 	}
// 	log.Printf("resp: %+v\n", resp)
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	log.Println("body:", string(bodyBytes))
// }

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "......"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.ListOptions{}
	repos, _, err := client.Apps.ListUserRepos(ctx, int64(12044600), opts)
	if err != nil {
		log.Println("error:", err.Error())
	}
	for _, repo := range repos {
		log.Println("repo:", *repo.Name)
	}
}
