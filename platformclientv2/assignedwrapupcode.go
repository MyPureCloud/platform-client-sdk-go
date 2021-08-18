package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignedwrapupcode
type Assignedwrapupcode struct { 
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The duration in seconds of the wrap-up segment.
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrap-up segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`

}

func (u *Assignedwrapupcode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignedwrapupcode

	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	

	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		*Alias
	}{ 
		Code: u.Code,
		
		Notes: u.Notes,
		
		Tags: u.Tags,
		
		DurationSeconds: u.DurationSeconds,
		
		EndTime: EndTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assignedwrapupcode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
