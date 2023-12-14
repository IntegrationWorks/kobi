package internal

import (
	"fmt"

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

	if bianVersion != BIAN_VERSION_12 && bianVersion != BIAN_VERSION_9_1 {
		return fmt.Errorf("api version %s not supported. supported versions are %s and %s", bianVersion, BIAN_VERSION_9_1, BIAN_VERSION_12)
	}

	if bianVersion == BIAN_VERSION_9_1 {
		if apiType != "" {
			return fmt.Errorf("api type only compatible with bian version 12. set the bian version with the --bian-version (-b) flag.")
		}
	}

	if apiType != "" && apiType != SEMANTIC_API && apiType != ISO_API {
		return fmt.Errorf("api type %s is not supported. supported values are %s or %s and are only valid for bian version %s", apiType, SEMANTIC_API, ISO_API, BIAN_VERSION_12)
	}

	return nil
}
