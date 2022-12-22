package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/ncaak/pifiabot/models"
)

func getBytes(data interface{}) ([]byte, error) {
	var ba, err = json.Marshal(data)
	if err != nil {
		log.Println("ERROR :: Encoding json object")
		return []byte{}, err
	}

	return ba, nil
}

func getReplyJson(data models.Output) []byte {
	var ba, err = getBytes(models.Reply{
		ChatId:  data.ChatId,
		ReplyId: data.MessageId,
		Text:    data.Text,
	})
	if err != nil {
		log.Println("ERROR :: JSON Marshaling Reply model : " + err.Error())
		return []byte{}
	}
	return ba
}

func send(url string, body []byte) {
	var req, errReq = http.NewRequest("POST", url, bytes.NewBuffer(body))
	if errReq != nil {
		log.Println("ERROR :: Creating Request Object : " + errReq.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{
		Timeout: 30 * time.Second,
	}

	var _, errDo = client.Do(req)
	if errDo != nil {
		log.Println("ERROR :: Sending HTTP request : " + errReq.Error())
	}
}

func getMultipartBody(data models.SetWebhook) (*bytes.Buffer, string) {
	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)
	defer writer.Close()

	writer.WriteField("url", data.Url)
	cert, _ := writer.CreateFormField("certificate")
	cert.Write(data.Certificate)

	// writer.CreateFormFile()
	return body, writer.FormDataContentType()
}

func getMultipartRequest(method string, data models.SetWebhook) (*http.Request, error) {
	var body, contentType = getMultipartBody(data)

	var req, err = http.NewRequest("POST", method, body)
	if err != nil {
		log.Println("ERROR :: Creating Multipart Request")
		return req, err
	}

	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func debugRequest(req *http.Request) {
	requestDump, _ := httputil.DumpRequestOut(req, true)

	log.Printf("DEBUG :: Request dump : \n%s\n", string(requestDump)) // TODO : Activate this on "debug" configuration
}

func debugResponse(resp *http.Response) {
	responseDump, _ := httputil.DumpResponse(resp, true)

	log.Printf("DEBUG :: Response dump : \n%s\n", string(responseDump)) // TODO : Activate this on "debug" configuration
}

func handleFailedResponse(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return fmt.Sprintf("HTTP Code %d :\n%s\n", resp.StatusCode, body)
}
