package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailcopyrecord
type Voicemailcopyrecord struct { 
	// User - The user that the voicemail message was copied to/from
	User *User `json:"user,omitempty"`


	// Group - The group that the voicemail message was copied to/from
	Group *Group `json:"group,omitempty"`


	// Date - The date when the voicemail was copied. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
