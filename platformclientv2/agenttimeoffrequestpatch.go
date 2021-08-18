package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenttimeoffrequestpatch
type Agenttimeoffrequestpatch struct { 
	// MarkedAsRead - Whether this request has been read by the agent
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// Status - The status of this time off request. Can only be canceled if the requested date has not already passed
	Status *string `json:"status,omitempty"`


	// Notes - Notes about the time off request. Can only be edited while the request is still pending
	Notes *string `json:"notes,omitempty"`

}

func (u *Agenttimeoffrequestpatch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agenttimeoffrequestpatch

	

	return json.Marshal(&struct { 
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		*Alias
	}{ 
		MarkedAsRead: u.MarkedAsRead,
		
		Status: u.Status,
		
		Notes: u.Notes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agenttimeoffrequestpatch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
