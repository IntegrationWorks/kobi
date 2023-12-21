package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v37/github"
)

func ListServices(bianVersion string, apiType string) {

	repoPath, fileExtension := GetRepositoryParams(bianVersion, apiType)
	log.Printf("Reading BIAN specs from %s", repoPath)

	client := github.NewClient(nil)

	tree, _, err := client.Git.GetTree(
		context.Background(),
		BIAN_ORG,
		BIAN_REPO,
		BIAN_BRANCH,
		true,
	)

	if err != nil {
		log.Printf("Failed to connect to BIAN github repository: %s", err.Error())
		return
	}

	for _, t := range tree.Entries {
		if strings.HasPrefix(t.GetPath(), repoPath) {
			domainName := t.GetPath()
			domainName = strings.Replace(domainName, repoPath, "", 1)
			domainName = strings.Replace(domainName, fileExtension, "", 1)
			if domainName != "Readme.md" {
				fmt.Println(domainName)
			}

		}
	}
}
