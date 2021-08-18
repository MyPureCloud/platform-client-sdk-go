package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateworkplanrotationagentrequest
type Updateworkplanrotationagentrequest struct { 
	// UserId - The ID of an agent in this work plan rotation
	UserId *string `json:"userId,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int `json:"position,omitempty"`


	// Delete - If marked true for this agent when updating, then this agent will be removed from this work plan rotation
	Delete *bool `json:"delete,omitempty"`

}

func (u *Updateworkplanrotationagentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateworkplanrotationagentrequest

	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Position *int `json:"position,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		
		DateRange: u.DateRange,
		
		Position: u.Position,
		
		Delete: u.Delete,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updateworkplanrotationagentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
