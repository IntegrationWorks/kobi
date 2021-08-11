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
