package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicalertnotification
type V2mobiusrulestopicalertnotification struct { 
	// Recipient
	Recipient *string `json:"recipient,omitempty"`


	// NotificationTypes
	NotificationTypes *[]string `json:"notificationTypes,omitempty"`

}

func (o *V2mobiusrulestopicalertnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusrulestopicalertnotification
	
	return json.Marshal(&struct { 
		Recipient *string `json:"recipient,omitempty"`
		
		NotificationTypes *[]string `json:"notificationTypes,omitempty"`
		*Alias
	}{ 
		Recipient: o.Recipient,
		
		NotificationTypes: o.NotificationTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusrulestopicalertnotification) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicalertnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicalertnotificationMap)
	if err != nil {
		return err
	}
	
	if Recipient, ok := V2mobiusrulestopicalertnotificationMap["recipient"].(string); ok {
		o.Recipient = &Recipient
	}
    
	if NotificationTypes, ok := V2mobiusrulestopicalertnotificationMap["notificationTypes"].([]interface{}); ok {
		NotificationTypesString, _ := json.Marshal(NotificationTypes)
		json.Unmarshal(NotificationTypesString, &o.NotificationTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicalertnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
