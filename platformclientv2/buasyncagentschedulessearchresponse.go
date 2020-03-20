package platformclientv2
import (
	"encoding/json"
)

// Buasyncagentschedulessearchresponse
type Buasyncagentschedulessearchresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Buagentschedulessearchresponse `json:"result,omitempty"`


	// Progress - Percent progress for the operation
	Progress *int32 `json:"progress,omitempty"`


	// DownloadUrl - The URL from which to download the result if it is too large to pass directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buasyncagentschedulessearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
