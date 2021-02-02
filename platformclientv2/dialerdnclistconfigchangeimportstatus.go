package platformclientv2
import (
	"encoding/json"
)

// Dialerdnclistconfigchangeimportstatus
type Dialerdnclistconfigchangeimportstatus struct { 
	// ImportState
	ImportState *string `json:"importState,omitempty"`


	// TotalRecords
	TotalRecords *int `json:"totalRecords,omitempty"`


	// CompletedRecords
	CompletedRecords *int `json:"completedRecords,omitempty"`


	// PercentageComplete
	PercentageComplete *int `json:"percentageComplete,omitempty"`


	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangeimportstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
