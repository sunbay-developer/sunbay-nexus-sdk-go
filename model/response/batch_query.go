package response

import (
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"
)

// BatchQueryResponse represents a batch query transaction response.
// The fields from the 'data' object are automatically flattened into this struct by the HTTP client.
type BatchQueryResponse struct {
	common.BaseResponse

	// BatchList is the list of batch statistics grouped by channel code and price currency
	BatchList []common.BatchQueryItem `json:"batchList"`
}
