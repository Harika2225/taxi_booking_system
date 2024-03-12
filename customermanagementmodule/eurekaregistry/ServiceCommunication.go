package eureka

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	app "com.example.customermanagement/config"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/micro/micro/v3/service/logger"
)

func ClientCommunication(w http.ResponseWriter, restServer string, api string, requestData interface{}) {
	uri := app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	cleanURL := strings.TrimSuffix(uri, "/apps/")
	client := eureka.NewClient([]string{cleanURL})
	res, _ := client.GetApplication(restServer)
	homePageURL := res.Instances[0].HomePageUrl
	url := homePageURL + api
	logger.Infof(url)
	// Marshal the requestData structure to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		logger.Errorf("Error marshaling request data: %s", err)
		return
	}

	// Create a new request with the provided data
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Errorf("Error creating request: %s", err)
		return
	}

	// Set the appropriate headers
	request.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and send the request
	clientWithAuth := &http.Client{}
	response, err := clientWithAuth.Do(request)
	if err != nil {
		logger.Errorf("Error sending request: %s", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("Error reading response body: %s", err)
		return
	}

	// Write the response to the HTTP response writer
	_, err = w.Write(body)
	if err != nil {
		logger.Errorf("Error writing response: %s", err)
		return
	}
}
