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

func (o *Voicemailcopyrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailcopyrecord
	
	Date := new(string)
	if o.Date != nil {
		
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Date *string `json:"date,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Group: o.Group,
		
		Date: Date,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailcopyrecord) UnmarshalJSON(b []byte) error {
	var VoicemailcopyrecordMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailcopyrecordMap)
	if err != nil {
		return err
	}
	
	if User, ok := VoicemailcopyrecordMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := VoicemailcopyrecordMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if dateString, ok := VoicemailcopyrecordMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateString)
		o.Date = &Date
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailcopyrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
