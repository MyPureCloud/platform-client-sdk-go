package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwardingeventcallforwarding
type Callforwardingeventcallforwarding struct { 
	// User
	User *Callforwardingeventuser `json:"user,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Calls
	Calls *[]Callforwardingeventcall `json:"calls,omitempty"`


	// Voicemail
	Voicemail *string `json:"voicemail,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (u *Callforwardingeventcallforwarding) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwardingeventcallforwarding

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		User *Callforwardingeventuser `json:"user,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Calls *[]Callforwardingeventcall `json:"calls,omitempty"`
		
		Voicemail *string `json:"voicemail,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Enabled: u.Enabled,
		
		Calls: u.Calls,
		
		Voicemail: u.Voicemail,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callforwardingeventcallforwarding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
