package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduleruntopicbuschedulingrunprogressnotification
type Wfmbuscheduleruntopicbuschedulingrunprogressnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Result
	Result *Wfmbuscheduleruntopicbuschedulerun `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulingrunprogressnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
