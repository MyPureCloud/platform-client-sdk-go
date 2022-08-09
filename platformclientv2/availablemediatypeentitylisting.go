package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availablemediatypeentitylisting
type Availablemediatypeentitylisting struct { 
	// Entities
	Entities *[]Availablemediatype `json:"entities,omitempty"`

}

func (o *Availablemediatypeentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availablemediatypeentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Availablemediatype `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Availablemediatypeentitylisting) UnmarshalJSON(b []byte) error {
	var AvailablemediatypeentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &AvailablemediatypeentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AvailablemediatypeentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availablemediatypeentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
