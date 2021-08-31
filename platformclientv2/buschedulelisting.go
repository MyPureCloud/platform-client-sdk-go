package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulelisting
type Buschedulelisting struct { 
	// Entities
	Entities *[]Buschedulelistitem `json:"entities,omitempty"`

}

func (o *Buschedulelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulelisting
	
	return json.Marshal(&struct { 
		Entities *[]Buschedulelistitem `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulelisting) UnmarshalJSON(b []byte) error {
	var BuschedulelistingMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BuschedulelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
