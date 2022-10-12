package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchuserroutingstatuseventrequest - A maximum of 100 events are allowed per request
type Batchuserroutingstatuseventrequest struct { 
	// UserRoutingStatusEvents - UserRoutingStatus events for this batch
	UserRoutingStatusEvents *[]Userroutingstatusevent `json:"userRoutingStatusEvents,omitempty"`

}

func (o *Batchuserroutingstatuseventrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchuserroutingstatuseventrequest
	
	return json.Marshal(&struct { 
		UserRoutingStatusEvents *[]Userroutingstatusevent `json:"userRoutingStatusEvents,omitempty"`
		*Alias
	}{ 
		UserRoutingStatusEvents: o.UserRoutingStatusEvents,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchuserroutingstatuseventrequest) UnmarshalJSON(b []byte) error {
	var BatchuserroutingstatuseventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BatchuserroutingstatuseventrequestMap)
	if err != nil {
		return err
	}
	
	if UserRoutingStatusEvents, ok := BatchuserroutingstatuseventrequestMap["userRoutingStatusEvents"].([]interface{}); ok {
		UserRoutingStatusEventsString, _ := json.Marshal(UserRoutingStatusEvents)
		json.Unmarshal(UserRoutingStatusEventsString, &o.UserRoutingStatusEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchuserroutingstatuseventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
