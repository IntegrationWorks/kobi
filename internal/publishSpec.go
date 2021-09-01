package internal

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	kongAddress    string
	kongAuthHeader string
	workspace      string
)

func PublishSpecToPortal(workspaceFlag string, filepath string) error {

	filepath = filepath + FILE_EXTENSION

	if workspaceFlag == "" {
		workspace = "default"
	} else {
		workspace = workspaceFlag
	}

	kongAddress = os.Getenv("KOBI_KONG_ADDR")
	kongAuthHeader = os.Getenv("KOBI_KONG_TOKEN")

	if kongAddress == "" {
		kongAddress = "http://127.0.0.1:8001"
	}

	kongAddress += "/" + workspace + "/files"

	// kongAddress = "https://enu0vp4i5npd.x.pipedream.net"
	// filepath = "temp.json"

	_, err := url.ParseRequestURI(kongAddress)

	if err != nil {
		fmt.Printf("Invalid URL %s configured for Kong. Please set the environment vairable \"kobi_KONG_ADDR\" correctly.", kongAddress)
		return err
	}

	fmt.Printf("Kong address is %s\n", kongAddress)

	fmt.Printf("Building Form")
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("path")
	_, err = io.Copy(fw, strings.NewReader("specs/"+filepath))
	if err != nil {
		return err
	}
	fw, err = writer.CreateFormFile("contents", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	// Close multipart writer.
	writer.Close()

	fmt.Println("Creating HTTP Request...")
	httpReq, err := http.NewRequest(http.MethodPost, kongAddress, bytes.NewReader(body.Bytes()))

	if err != nil {
		fmt.Print("Failed to read create HTTP request")
		return err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	if kongAuthHeader != "" {
		fmt.Println("Adding authentication token")
		httpReq.Header.Set("kong-admin-token", kongAuthHeader)
	}

	fmt.Println("Sending HTTP Request...")

	res, err := http.DefaultClient.Do(httpReq)

	if err != nil {
		fmt.Print("Failed to send HTTP request to Kong")
		return err
	}

	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print("Failed to read response from Kong")
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		fmt.Print("Kong returned error")
		return fmt.Errorf("kong returned http response code %d", res.StatusCode)
	}

	fmt.Println("Success...")

	return nil

}
