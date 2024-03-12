package eureka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	app "com.example.bookingmanagement/config"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/micro/micro/v3/service/logger"
)

func ClientCommunication(r *http.Request, w http.ResponseWriter, restServer string, api string, requestData interface{}) {
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
	request.Header.Set("Authorization", r.Header.Get("Authorization"))
	logger.Info("REQUEST BOOK AT BOOKING SERVICE IN EURKA CLINET", request.Body)
	fmt.Println(1)
	// Create a new HTTP client and send the request
	clientWithAuth := &http.Client{}
	fmt.Println(2)
	response, err := clientWithAuth.Do(request)
	fmt.Println(3)
	if err != nil {
		logger.Errorf("Error sending request: %s", err)
		return
	}
	fmt.Print("RS STATUS", response.Status)
	// defer response.Body.Close()
	logger.Info("qqqqqqqqqqqqqqqqqqqqqqqqq", response.Body, "responseeeeeeeeeeeeeeeeeeeeeeeeeeee")

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(body, "ppppppppppppppppppppppppppppppppppppp")
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
