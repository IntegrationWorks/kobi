package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v37/github"
)

func ListServices(bianVersion string, apiType string) {

	fileExtension := ""
	repoPath := ""

	if bianVersion == BIAN_VERSION_12 {
		fileExtension = FILE_EXTENSION_YAML
		if apiType == SEMANTIC_API {
			repoPath = REPO_PATH_12_SEMANTIC
		} else {
			repoPath = REPO_PATH_12_ISO
		}
	} else {
		fileExtension = FILE_EXTENSION_JSON
		repoPath = REPO_PATH_9_1
	}
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
