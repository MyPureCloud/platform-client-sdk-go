package platformclientv2
import (
	"encoding/json"
)

// Importstatus
type Importstatus struct { 
	// State - current status of the import
	State *string `json:"state,omitempty"`


	// TotalRecords - total number of records to be imported
	TotalRecords *int64 `json:"totalRecords,omitempty"`


	// CompletedRecords - number of records finished importing
	CompletedRecords *int64 `json:"completedRecords,omitempty"`


	// PercentComplete - percentage of records finished importing
	PercentComplete *int32 `json:"percentComplete,omitempty"`


	// FailureReason - if the import has failed, the reason for the failure
	FailureReason *string `json:"failureReason,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
