package platformclientv2
import (
	"encoding/json"
)

// Dialercontactlistconfigchangeimportstatus
type Dialercontactlistconfigchangeimportstatus struct { 
	// ImportState
	ImportState *string `json:"importState,omitempty"`


	// TotalRecords
	TotalRecords *int32 `json:"totalRecords,omitempty"`


	// CompletedRecords
	CompletedRecords *int32 `json:"completedRecords,omitempty"`


	// PercentageComplete
	PercentageComplete *int32 `json:"percentageComplete,omitempty"`


	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangeimportstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
