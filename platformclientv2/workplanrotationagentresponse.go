package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanrotationagentresponse
type Workplanrotationagentresponse struct { 
	// User - The user associated with this work plan rotation
	User *Userreference `json:"user,omitempty"`


	// DateRange - The date range to which this agent is effective in the work plan rotation
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Position - Start position of the work plan in the pattern for this agent in the work plan rotation. Position value starts from 0
	Position *int `json:"position,omitempty"`

}

func (o *Workplanrotationagentresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanrotationagentresponse
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Position *int `json:"position,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		DateRange: o.DateRange,
		
		Position: o.Position,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanrotationagentresponse) UnmarshalJSON(b []byte) error {
	var WorkplanrotationagentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanrotationagentresponseMap)
	if err != nil {
		return err
	}
	
	if User, ok := WorkplanrotationagentresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if DateRange, ok := WorkplanrotationagentresponseMap["dateRange"].(map[string]interface{}); ok {
		DateRangeString, _ := json.Marshal(DateRange)
		json.Unmarshal(DateRangeString, &o.DateRange)
	}
	
	if Position, ok := WorkplanrotationagentresponseMap["position"].(float64); ok {
		PositionInt := int(Position)
		o.Position = &PositionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanrotationagentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
