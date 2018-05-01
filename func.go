package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"
	"github.com/goboilerplates/core"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(HandleRequest))
}

var (
	api GetSamplesAllAPI
	// TestResult for checking tests.
	TestResult bool
)

// CreateAPI creates API for GetSamplesAll.
func CreateAPI() GetSamplesAllAPI {
	if api == nil {
		return GetSamplesAllAPIImpl{
			Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
		}
	}
	return api
}

// HandleRequest handles request.
func HandleRequest(ctx context.Context, in io.Reader, out io.Writer) {
	api = CreateAPI()

	result, err := api.All()
	if err != nil {
		TestResult = false
		return
	}
	json.NewEncoder(out).Encode(&result)
	TestResult = true
}
