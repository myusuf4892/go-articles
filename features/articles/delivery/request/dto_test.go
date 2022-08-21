package request

import (
	"testing"
)

func TestRequestArticle(t *testing.T) {
	dataReq := Article{
		Title:      "Sample article 1",
		CategoryID: 1,
	}

	RequestToCore(dataReq)
}
