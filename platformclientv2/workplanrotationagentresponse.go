package platformclientv2
import (
	"encoding/json"
)

// Workplanrotationagentresponse
type Workplanrotationagentresponse struct { 
	// User - The user associated with this work plan rotation
	User *Userreference `json:"user,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int32 `json:"position,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationagentresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
