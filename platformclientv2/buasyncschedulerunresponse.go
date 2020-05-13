package platformclientv2
import (
	"encoding/json"
)

// Buasyncschedulerunresponse
type Buasyncschedulerunresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Buschedulerun `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buasyncschedulerunresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
