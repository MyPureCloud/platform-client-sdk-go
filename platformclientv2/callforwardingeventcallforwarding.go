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

func (o *Callforwardingeventcallforwarding) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwardingeventcallforwarding
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		User: o.User,
		
		Enabled: o.Enabled,
		
		Calls: o.Calls,
		
		Voicemail: o.Voicemail,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Callforwardingeventcallforwarding) UnmarshalJSON(b []byte) error {
	var CallforwardingeventcallforwardingMap map[string]interface{}
	err := json.Unmarshal(b, &CallforwardingeventcallforwardingMap)
	if err != nil {
		return err
	}
	
	if User, ok := CallforwardingeventcallforwardingMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Enabled, ok := CallforwardingeventcallforwardingMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Calls, ok := CallforwardingeventcallforwardingMap["calls"].([]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Voicemail, ok := CallforwardingeventcallforwardingMap["voicemail"].(string); ok {
		o.Voicemail = &Voicemail
	}
    
	if modifiedDateString, ok := CallforwardingeventcallforwardingMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callforwardingeventcallforwarding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
