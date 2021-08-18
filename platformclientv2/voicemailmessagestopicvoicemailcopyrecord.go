package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmessagestopicvoicemailcopyrecord
type Voicemailmessagestopicvoicemailcopyrecord struct { 
	// User
	User *Voicemailmessagestopicowner `json:"user,omitempty"`


	// Group
	Group *Voicemailmessagestopicowner `json:"group,omitempty"`

}

func (u *Voicemailmessagestopicvoicemailcopyrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmessagestopicvoicemailcopyrecord

	

	return json.Marshal(&struct { 
		User *Voicemailmessagestopicowner `json:"user,omitempty"`
		
		Group *Voicemailmessagestopicowner `json:"group,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Group: u.Group,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
