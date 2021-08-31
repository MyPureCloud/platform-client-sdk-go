package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationsrequest
type Updatenotificationsrequest struct { 
	// Entities - The notifications to update
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

func (o *Updatenotificationsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatenotificationsrequest
	
	return json.Marshal(&struct { 
		Entities *[]Wfmusernotification `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatenotificationsrequest) UnmarshalJSON(b []byte) error {
	var UpdatenotificationsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatenotificationsrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := UpdatenotificationsrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatenotificationsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
