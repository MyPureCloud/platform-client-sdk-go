package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Subscriberresponse
type Subscriberresponse struct { 
	// MessageReturned - Suggested valid addresses
	MessageReturned *[]string `json:"messageReturned,omitempty"`


	// Status - http status
	Status *string `json:"status,omitempty"`

}

func (o *Subscriberresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Subscriberresponse
	
	return json.Marshal(&struct { 
		MessageReturned *[]string `json:"messageReturned,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		MessageReturned: o.MessageReturned,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Subscriberresponse) UnmarshalJSON(b []byte) error {
	var SubscriberresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SubscriberresponseMap)
	if err != nil {
		return err
	}
	
	if MessageReturned, ok := SubscriberresponseMap["messageReturned"].([]interface{}); ok {
		MessageReturnedString, _ := json.Marshal(MessageReturned)
		json.Unmarshal(MessageReturnedString, &o.MessageReturned)
	}
	
	if Status, ok := SubscriberresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Subscriberresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
