package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)



func RequestToAssembly(url string, body []byte, isJSON bool, requestType string) (*http.Client, *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest(requestType, url, bytes.NewBuffer(body))
	if isJSON {
		req.Header.Set("content-type", "application/json")
	}
	req.Header.Set("authorization", ApiKeyToAssembly)
	return client, req
}

func decodeJSON(res *http.Response) map[string]interface{} {
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	return result
}

func TranscriptMessage(data []byte) (string, error) {
	// Setup HTTP client and set header
	client, req := RequestToAssembly(UploadURLToAssembly, data, false, "POST")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// Decode json and store it in a map
	result := decodeJSON(res)
	// Print the upload_url
	AUDIO_URL := result["upload_url"].(string)
	// Prepare json data
	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, err := json.Marshal(values)

	if err != nil {
		return "", err
	}

	// Setup HTTP client and set header
	client, req = RequestToAssembly(TRANSCRIPT_URL, jsonData, true, "POST")
	res, err = client.Do(req)

	if err != nil {
		return "", err
	}

	// Decode json and store it in a map
	result = decodeJSON(res)

	POLLING_URL := TRANSCRIPT_URL + "/" + result["id"].(string)

	// Send GET request
	client, req = RequestToAssembly(POLLING_URL, nil, true, "GET")
	var text string
	for {
		res, err = client.Do(req)

		if err != nil {
			return "", err
		}

		result = decodeJSON(res)

		// Check status and print the transcribed text
		if result["status"] == "completed" {
			text = result["text"].(string)
			break
		}

		time.Sleep(2 * time.Second)
	}
	return text, nil
}

//85 строк
