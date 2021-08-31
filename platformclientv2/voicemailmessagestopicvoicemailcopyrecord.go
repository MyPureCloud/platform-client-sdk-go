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

func (o *Voicemailmessagestopicvoicemailcopyrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmessagestopicvoicemailcopyrecord
	
	return json.Marshal(&struct { 
		User *Voicemailmessagestopicowner `json:"user,omitempty"`
		
		Group *Voicemailmessagestopicowner `json:"group,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Group: o.Group,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailmessagestopicvoicemailcopyrecord) UnmarshalJSON(b []byte) error {
	var VoicemailmessagestopicvoicemailcopyrecordMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmessagestopicvoicemailcopyrecordMap)
	if err != nil {
		return err
	}
	
	if User, ok := VoicemailmessagestopicvoicemailcopyrecordMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := VoicemailmessagestopicvoicemailcopyrecordMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
