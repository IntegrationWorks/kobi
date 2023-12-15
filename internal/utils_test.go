package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBianVersionAndAPITypeValidParams(t *testing.T) {
	bianVersion := "10"
	apiType := ""
	error := ValidateBianVersionAndAPIType(bianVersion, apiType)
	assert.Nil(t, error)
}

func TestValidateIncorrectVersion(t *testing.T) {
	bianVersion := "1"
	apiType := ""
	error := ValidateBianVersionAndAPIType(bianVersion, apiType)
	assert.NotNil(t, error)
	assert.Equal(t, "api version 1 not supported. supported versions are [9.1 10 11 12]", error.Error())
}

func TestValidateApiTypeWithNotV12(t *testing.T) {
	bianVersion := "10"
	apiType := "iso"
	error := ValidateBianVersionAndAPIType(bianVersion, apiType)
	assert.NotNil(t, error)
	assert.Equal(t, "api type 'iso' only compatible with bian version 12. set the bian version with the --bian-version (-b) flag.", error.Error())
}
func TestValidateApiTypeWithV12(t *testing.T) {
	bianVersion := "12"
	apiType := "iso"
	error := ValidateBianVersionAndAPIType(bianVersion, apiType)
	assert.Nil(t, error)
}

func TestValidateAPITypeInvalid(t *testing.T) {
	bianVersion := "12"
	apiType := "bork"
	error := ValidateBianVersionAndAPIType(bianVersion, apiType)
	assert.NotNil(t, error)
	assert.Equal(t, "api type bork is not supported. supported values are semantic or iso and are only valid for bian version 12", error.Error())
}
