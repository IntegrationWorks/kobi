package internal

import (
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
)

func ValidateNoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("positional arguments are not valid for this command, " +
			"please use flags instead")
	}
	return nil
}

func ValidateBianVersionAndAPIType(bianVersion string, apiType string) error {

	if !slices.Contains(SUPPORTED_VERSIONS(), bianVersion) {
		return fmt.Errorf("api version %s not supported. supported versions are %v", bianVersion, SUPPORTED_VERSIONS())
	}

	if bianVersion != BIAN_VERSION_12 {
		if apiType != "" && apiType != SEMANTIC_API {
			return fmt.Errorf("api type 'iso' only compatible with bian version 12. set the bian version with the --bian-version (-b) flag.")
		}
	}

	if apiType != "" && apiType != SEMANTIC_API && apiType != ISO_API {
		return fmt.Errorf("api type %s is not supported. supported values are %s or %s and are only valid for bian version %s", apiType, SEMANTIC_API, ISO_API, BIAN_VERSION_12)
	}

	return nil
}

func GetRepositoryParams(bianVersion string, apiType string) (repoPath string, fileExtension string) {

	repoPath = ""
	fileExtension = ""

	if bianVersion == BIAN_VERSION_12 {
		fileExtension = FILE_EXTENSION_YAML
		if apiType == SEMANTIC_API {
			repoPath = REPO_PATH_12_SEMANTIC
		} else {
			repoPath = REPO_PATH_12_ISO
		}
	} else if bianVersion == BIAN_VERSION_11 {
		fileExtension = FILE_EXTENSION_YAML
		repoPath = REPO_PATH_11
	} else if bianVersion == BIAN_VERSION_10 {
		fileExtension = FILE_EXTENSION_YAML
		repoPath = REPO_PATH_10
	} else {
		fileExtension = FILE_EXTENSION_JSON
		repoPath = REPO_PATH_9_1
	}

	return repoPath, fileExtension

}
