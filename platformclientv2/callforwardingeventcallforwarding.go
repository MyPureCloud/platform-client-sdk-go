package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Callforwardingeventcallforwarding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
