package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Expirededgelisting
type Expirededgelisting struct { 
	// Entities
	Entities *[]Edgeidnamepair `json:"entities,omitempty"`

}

func (o *Expirededgelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Expirededgelisting
	
	return json.Marshal(&struct { 
		Entities *[]Edgeidnamepair `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Expirededgelisting) UnmarshalJSON(b []byte) error {
	var ExpirededgelistingMap map[string]interface{}
	err := json.Unmarshal(b, &ExpirededgelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ExpirededgelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Expirededgelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
