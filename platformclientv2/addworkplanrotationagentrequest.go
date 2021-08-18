package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Addworkplanrotationagentrequest
type Addworkplanrotationagentrequest struct { 
	// UserId - The ID of an agent in this work plan rotation
	UserId *string `json:"userId,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int `json:"position,omitempty"`

}

func (u *Addworkplanrotationagentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addworkplanrotationagentrequest

	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Position *int `json:"position,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		
		DateRange: u.DateRange,
		
		Position: u.Position,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Addworkplanrotationagentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
