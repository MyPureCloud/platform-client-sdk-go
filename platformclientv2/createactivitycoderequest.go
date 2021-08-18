package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createactivitycoderequest - Activity Code
type Createactivitycoderequest struct { 
	// Name - The name of the activity code
	Name *string `json:"name,omitempty"`


	// Category - The activity code's category
	Category *string `json:"category,omitempty"`


	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsWorkTime - Indicates whether or not the activity should be counted as work time
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`


	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`

}

func (u *Createactivitycoderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createactivitycoderequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`
		
		AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Category: u.Category,
		
		LengthInMinutes: u.LengthInMinutes,
		
		CountsAsPaidTime: u.CountsAsPaidTime,
		
		CountsAsWorkTime: u.CountsAsWorkTime,
		
		AgentTimeOffSelectable: u.AgentTimeOffSelectable,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createactivitycoderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
