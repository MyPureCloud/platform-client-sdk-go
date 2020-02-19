package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduletopicbuschedulenotification
type Wfmbuscheduletopicbuschedulenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`


	// Result
	Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
