package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicalertnotification
type V2mobiusalertstopicalertnotification struct { 
	// Recipient
	Recipient *string `json:"recipient,omitempty"`


	// NotificationTypes
	NotificationTypes *[]string `json:"notificationTypes,omitempty"`

}

func (o *V2mobiusalertstopicalertnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopicalertnotification
	
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

func (o *V2mobiusalertstopicalertnotification) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicalertnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicalertnotificationMap)
	if err != nil {
		return err
	}
	
	if Recipient, ok := V2mobiusalertstopicalertnotificationMap["recipient"].(string); ok {
		o.Recipient = &Recipient
	}
    
	if NotificationTypes, ok := V2mobiusalertstopicalertnotificationMap["notificationTypes"].([]interface{}); ok {
		NotificationTypesString, _ := json.Marshal(NotificationTypes)
		json.Unmarshal(NotificationTypesString, &o.NotificationTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicalertnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
