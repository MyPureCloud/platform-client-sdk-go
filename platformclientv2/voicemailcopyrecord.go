package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Voicemailcopyrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailcopyrecord

	
	Date := new(string)
	if u.Date != nil {
		
		*Date = timeutil.Strftime(u.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	

	return json.Marshal(&struct { 
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Date *string `json:"date,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Group: u.Group,
		
		Date: Date,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
