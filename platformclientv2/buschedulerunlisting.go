package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulerunlisting
type Buschedulerunlisting struct { 
	// Entities
	Entities *[]Buschedulerun `json:"entities,omitempty"`

}

func (o *Buschedulerunlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buschedulerunlisting
	
	return json.Marshal(&struct { 
		Entities *[]Buschedulerun `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Buschedulerunlisting) UnmarshalJSON(b []byte) error {
	var BuschedulerunlistingMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulerunlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BuschedulerunlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulerunlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
