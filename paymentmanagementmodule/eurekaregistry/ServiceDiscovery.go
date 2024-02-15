package eureka

import (
	"io/ioutil"
	"net/http"
	"strings"

	app "com.example.paymentmanagement/config"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/micro/micro/v3/service/logger"
)

func Client(w http.ResponseWriter, req *http.Request, restServer string) {
	uri := app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	cleanURL := strings.TrimSuffix(uri, "/apps/")
	client := eureka.NewClient([]string{cleanURL})
	res, _ := client.GetApplication(restServer)
	homePageURL := res.Instances[0].HomePageUrl
	url := homePageURL + "rest/services/" + restServer
	clientWithAuth := &http.Client{
		Transport: &headerTransport{
			Transport: http.DefaultTransport,
			Header:    req.Header,
		},
	}
	response, err := clientWithAuth.Get(url)
	if err != nil {
		logger.Errorf("Error sending GET request: %s", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("Error reading response body: %s", err)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		logger.Errorf("Error writing response: %s", err)
		return
	}
}

type headerTransport struct {
	Transport http.RoundTripper
	Header    http.Header
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token := t.Header.Get("Authorization")
	req.Header.Set("Authorization", token)
	return t.Transport.RoundTrip(req)
}
