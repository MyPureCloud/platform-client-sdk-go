package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitlisting
type Businessunitlisting struct { 
	// Entities
	Entities *[]Businessunitlistitem `json:"entities,omitempty"`

}

func (o *Businessunitlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Businessunitlisting
	
	return json.Marshal(&struct { 
		Entities *[]Businessunitlistitem `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Businessunitlisting) UnmarshalJSON(b []byte) error {
	var BusinessunitlistingMap map[string]interface{}
	err := json.Unmarshal(b, &BusinessunitlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BusinessunitlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Businessunitlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
