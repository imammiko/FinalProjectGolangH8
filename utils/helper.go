package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func FormatVlidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func SendImageToCloud(image multipart.File, imageName string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var param = url.Values{}
	param.Set("name", imageName)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("fileName")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, strings.NewReader(imageName))
	if err != nil {
		return nil, err
	}
	fw, err = writer.CreateFormFile("file", imageName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, image)
	if err != nil {
		return nil, err
	}
	writer.Close()
	url := Getenv("URLIMAGEKITUPLOAD", "")
	fmt.Println(url, "sini")
	PirvateKey := Getenv("PRIVATEKEYIMAGEKIT", "")
	privatekeyEncode := base64.StdEncoding.EncodeToString([]byte(PirvateKey))
	fmt.Println(privatekeyEncode, "sini")
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Basic "+privatekeyEncode+"Og==")
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		err := errors.New("messsage : upload failed")
		return nil, err
	}
	defer rsp.Body.Close()
	var responseData = make(map[string]interface{})
	err = json.NewDecoder(rsp.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
