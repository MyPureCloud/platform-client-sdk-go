package platformclientv2
import (
	"encoding/json"
)

// Updateworkplanrotationagentrequest
type Updateworkplanrotationagentrequest struct { 
	// UserId - The ID of an agent in this work plan rotation
	UserId *string `json:"userId,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int32 `json:"position,omitempty"`


	// Delete - If marked true for this agent when updating, then this agent will be removed from this work plan rotation
	Delete *bool `json:"delete,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationagentrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
