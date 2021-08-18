package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapup
type Wrapup struct { 
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`


	// Name - The user configured wrap up code name.
	Name *string `json:"name,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The length of time in seconds that the agent spent doing after call work.
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrapup was finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// Provisional - Indicates if this is a pending save and should not require a code to be specified.  This allows someone to save some temporary wrapup that will be used later.
	Provisional *bool `json:"provisional,omitempty"`

}

func (u *Wrapup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wrapup

	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	

	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Provisional *bool `json:"provisional,omitempty"`
		*Alias
	}{ 
		Code: u.Code,
		
		Name: u.Name,
		
		Notes: u.Notes,
		
		Tags: u.Tags,
		
		DurationSeconds: u.DurationSeconds,
		
		EndTime: EndTime,
		
		Provisional: u.Provisional,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
