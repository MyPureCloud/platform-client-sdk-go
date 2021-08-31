package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationsresponse
type Updatenotificationsresponse struct { 
	// Entities
	Entities *[]Updatenotificationresponse `json:"entities,omitempty"`

}

func (o *Updatenotificationsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatenotificationsresponse
	
	return json.Marshal(&struct { 
		Entities *[]Updatenotificationresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatenotificationsresponse) UnmarshalJSON(b []byte) error {
	var UpdatenotificationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatenotificationsresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := UpdatenotificationsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatenotificationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
