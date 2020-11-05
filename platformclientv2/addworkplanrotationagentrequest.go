package platformclientv2
import (
	"encoding/json"
)

// Addworkplanrotationagentrequest
type Addworkplanrotationagentrequest struct { 
	// UserId - The ID of an agent in this work plan rotation
	UserId *string `json:"userId,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int32 `json:"position,omitempty"`

}

// String returns a JSON representation of the model
func (o *Addworkplanrotationagentrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
