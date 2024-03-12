package eureka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	app "com.example.drivermanagement/config"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/micro/micro/v3/service/logger"
)

func ClientCommunication(r *http.Request, w http.ResponseWriter, restServer string, api string, requestData interface{}) {
	uri := app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	logger.Info(uri)
	cleanURL := strings.TrimSuffix(uri, "/apps/")
	client := eureka.NewClient([]string{cleanURL})
	res, _ := client.GetApplication(restServer)
	homePageURL := res.Instances[0].HomePageUrl
	url := homePageURL + api
	fmt.Print("CALLING BOOKING ACCEPTED")
	// url := "http://localhost:9022/api/bookingAccepted"

	logger.Infof(url)
	// Marshal the requestData structure to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		logger.Errorf("Error marshaling request data: %s", err)
		return
	}
	logger.Info("REQUEST BODY:-", requestBody)
	// logger.Info("TOKEN:-", r.Header.Get("Authorization"))
	logger.Info(1)

	requestBodyBuffer := bytes.NewBuffer(requestBody)
	logger.Info(2)
	// Create a new request with the provided data
	request, err := http.NewRequest("POST", url, requestBodyBuffer)
	if err != nil {
		logger.Errorf("Error creating request: %s", err)
		return
	}
	logger.Info(3)

	// Set the appropriate headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", r.Header.Get("Authorization"))

	logger.Info(4)
	clientWithAuth := &http.Client{}
	logger.Info(5)
	response, err := clientWithAuth.Do(request)
	logger.Info(6)
	fmt.Print("RESPONSE STATUS:-", response)
	if err != nil {
		logger.Errorf("Error sending request: %s", err)
		return
	}

	// defer response.Body.Close()

	// // Read the response body
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	logger.Errorf("Error reading response body: %s", err)
	// 	return
	// }

	// // Write the response to the HTTP response writer
	// _, err = w.Write(body)
	// if err != nil {
	// 	logger.Errorf("Error writing response: %s", err)
	// 	return
	// }
}
