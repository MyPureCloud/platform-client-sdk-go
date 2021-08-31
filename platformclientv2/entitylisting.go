package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entitylisting
type Entitylisting struct { 
	// Entities
	Entities *[]interface{} `json:"entities,omitempty"`

}

func (o *Entitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entitylisting
	
	return json.Marshal(&struct { 
		Entities *[]interface{} `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Entitylisting) UnmarshalJSON(b []byte) error {
	var EntitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &EntitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := EntitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Entitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
