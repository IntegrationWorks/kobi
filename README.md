# KOBI

`Kobi` is a CLI to manage and automate the deployment of [BIAN](https://bian.org) service domain API specifications to a [Kong](https://konghq.com/)

![KOBI CLI](doc/main.png)

## Features

* List all avaialbe BIAN service domain API Specifications for download or publishing
* Download a BIAN API service domain API Specification to the local machine
* Publish an API Specification from the BIAN API Repository directly to a running Kong Developer Portal
* Initialise the supporting files for a Kong Developer Portal

## Compatibility

Kobi is compatible with Kong Enterprise, version 1.0 and greater, and BIAN 9.0. Future releases will be compatibile with later versions of BIAN and Kong Konnect Cloud.

## Installation

The `kobi` CLI is developed on `golang` using `cobra` CLI. The executable `./kobi` is included in this repository, or it can be built from source (`go build`). For universal access, place the `kobi`  binary in a folder inclded on the path (e.g. `~/go/bin/kobi`).


## Configuration

`Kobi` uses environment variables to connect to the Kong Developer Portal.  

* `KOBI_KONG_ADDR` (Defaults to http://localhost:8001) - URL of Kong Admin API
* `KOBI_KONG_TOKEN` - RBAC token for configured Kong user with write permissions to the Kong Developer Portal Files API


## Usage

Use the `--help` flag to view usage instructions in the terminal

### Version

Use to list the installed version of `Kobi`

```bash

```