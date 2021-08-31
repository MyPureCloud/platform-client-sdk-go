package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmuserentitylisting
type Wfmuserentitylisting struct { 
	// Entities
	Entities *[]User `json:"entities,omitempty"`

}

func (o *Wfmuserentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmuserentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]User `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmuserentitylisting) UnmarshalJSON(b []byte) error {
	var WfmuserentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &WfmuserentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WfmuserentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmuserentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
