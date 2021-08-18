package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationthreadingwindow
type Conversationthreadingwindow struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Settings - The conversation threading window timeout (Minutes) for each messaging type
	Settings *[]Conversationthreadingwindowsetting `json:"settings,omitempty"`


	// DefaultTimeoutMinutes - The default conversation threading window timeout (Minutes)
	DefaultTimeoutMinutes *int `json:"defaultTimeoutMinutes,omitempty"`

}

func (u *Conversationthreadingwindow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationthreadingwindow

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Settings *[]Conversationthreadingwindowsetting `json:"settings,omitempty"`
		
		DefaultTimeoutMinutes *int `json:"defaultTimeoutMinutes,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Settings: u.Settings,
		
		DefaultTimeoutMinutes: u.DefaultTimeoutMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
