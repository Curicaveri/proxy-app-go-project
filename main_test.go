package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/Curicaveri/proxy-app/api/handlers"
	"github.com/Curicaveri/proxy-app/api/middleware"
	"github.com/Curicaveri/proxy-app/api/server"
	"github.com/stretchr/testify/assert"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		app := server.SetUp()
		handlers.SetupRouter(app)
		middleware.InitQueue()
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("Server running...")
}

type Response struct {
	Status   string `json:"status,omitempty"`
	Response string `json:"result,omitempty"`
}

func TestAlgorithm(t *testing.T) {
	cases := []struct {
		Domain string
		Output string
	}{
		{
			Domain: "alpha",
			Output: "[alpha]",
		},
		{
			Domain: "omega",
			Output: "[alfa, omega]",
		},
		{
			Domain: "alpha",
			Output: "[alpha, alpha, omega]",
		},
		{
			Domain: "",
			Output: "error",
		},
	}

	valuesToCompare := &Response{}
	client := &http.Client{}
	for _, singleCase := range cases {
		req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
		req.Header.Add("domain", singleCase.Domain)
		response, _ := client.Do(req)

		bytes, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bytes, valuesToCompare)

		fmt.Println("RESPONSE: ", valuesToCompare.Response)

		assert.Equal(t, singleCase.Output, valuesToCompare.Response)
	}
}
