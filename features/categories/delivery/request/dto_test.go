package request

import "testing"

func TestRequestCategory(t *testing.T) {
	dataReq := Category{
		Name: "success-story",
	}

	RequestToCore(dataReq)
}
