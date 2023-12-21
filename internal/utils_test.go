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

func TestGetRepositoryParams12Semantic(t *testing.T) {
	bianVersion := "12"
	apiType := "semantic"
	repoPath, fileExtention := GetRepositoryParams(bianVersion, apiType)
	assert.Equal(t, FILE_EXTENSION_YAML, fileExtention)
	assert.Equal(t, REPO_PATH_12_SEMANTIC, repoPath)
}

func TestGetRepositoryParams12ISO(t *testing.T) {
	bianVersion := "12"
	apiType := "iso"
	repoPath, fileExtention := GetRepositoryParams(bianVersion, apiType)
	assert.Equal(t, FILE_EXTENSION_YAML, fileExtention)
	assert.Equal(t, REPO_PATH_12_ISO, repoPath)
}

func TestGetRepositoryParams11(t *testing.T) {
	bianVersion := "11"
	apiType := ""
	repoPath, fileExtention := GetRepositoryParams(bianVersion, apiType)
	assert.Equal(t, FILE_EXTENSION_YAML, fileExtention)
	assert.Equal(t, REPO_PATH_11, repoPath)
}

func TestGetRepositoryParams10(t *testing.T) {
	bianVersion := "10"
	apiType := ""
	repoPath, fileExtention := GetRepositoryParams(bianVersion, apiType)
	assert.Equal(t, FILE_EXTENSION_YAML, fileExtention)
	assert.Equal(t, REPO_PATH_10, repoPath)
}

func TestGetRepositoryParams9(t *testing.T) {
	bianVersion := "9.1"
	apiType := ""
	repoPath, fileExtention := GetRepositoryParams(bianVersion, apiType)
	assert.Equal(t, FILE_EXTENSION_JSON, fileExtention)
	assert.Equal(t, REPO_PATH_9_1, repoPath)
}
