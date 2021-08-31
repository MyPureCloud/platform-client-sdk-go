package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationsresponse
type Notificationsresponse struct { 
	// Entities
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

func (o *Notificationsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Notificationsresponse
	
	return json.Marshal(&struct { 
		Entities *[]Wfmusernotification `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Notificationsresponse) UnmarshalJSON(b []byte) error {
	var NotificationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &NotificationsresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := NotificationsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Notificationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
