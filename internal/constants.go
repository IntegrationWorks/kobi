package internal

const BIAN_ORG string = "bian-official"
const BIAN_REPO string = "public"
const BIAN_BRANCH string = "main"

const REPO_PATH_9_1 string = "release9.1/semantic-apis/swaggers/"
const REPO_PATH_10 string = "release10.0.0/semantic-apis/oas3/yamls/"
const REPO_PATH_11 string = "release11.0.0/semantic-apis/oas3/yamls/"

const REPO_PATH_12_SEMANTIC string = "release12.0.0/semantic-apis/oas3/yamls/"
const REPO_PATH_12_ISO string = "release12.0.0/apis-iso20022_ext-ddd/oas3/yamls/"

const FILE_EXTENSION_JSON string = ".json"
const FILE_EXTENSION_YAML string = ".yaml"

const KONG_ORG string = "Kong"
const KONG_PORTAL_REPO string = "kong-portal-templates"

const BIAN_VERSION_9_1 = "9.1"
const BIAN_VERSION_10 = "10"
const BIAN_VERSION_11 = "11"
const BIAN_VERSION_12 = "12"

func SUPPORTED_VERSIONS() []string {
	return []string{BIAN_VERSION_9_1, BIAN_VERSION_10, BIAN_VERSION_11, BIAN_VERSION_12}
}

const SEMANTIC_API = "semantic"
const ISO_API = "iso"
