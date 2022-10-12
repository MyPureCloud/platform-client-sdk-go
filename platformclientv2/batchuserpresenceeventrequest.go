package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchuserpresenceeventrequest - A maximum of 100 events are allowed per request
type Batchuserpresenceeventrequest struct { 
	// UserPresenceEvents - UserPresence events for this batch
	UserPresenceEvents *[]Userpresenceevent `json:"userPresenceEvents,omitempty"`

}

func (o *Batchuserpresenceeventrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchuserpresenceeventrequest
	
	return json.Marshal(&struct { 
		UserPresenceEvents *[]Userpresenceevent `json:"userPresenceEvents,omitempty"`
		*Alias
	}{ 
		UserPresenceEvents: o.UserPresenceEvents,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchuserpresenceeventrequest) UnmarshalJSON(b []byte) error {
	var BatchuserpresenceeventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BatchuserpresenceeventrequestMap)
	if err != nil {
		return err
	}
	
	if UserPresenceEvents, ok := BatchuserpresenceeventrequestMap["userPresenceEvents"].([]interface{}); ok {
		UserPresenceEventsString, _ := json.Marshal(UserPresenceEvents)
		json.Unmarshal(UserPresenceEventsString, &o.UserPresenceEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchuserpresenceeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
