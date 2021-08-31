package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletopicentitylisting
type Availabletopicentitylisting struct { 
	// Entities
	Entities *[]Availabletopic `json:"entities,omitempty"`

}

func (o *Availabletopicentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletopicentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Availabletopic `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletopicentitylisting) UnmarshalJSON(b []byte) error {
	var AvailabletopicentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletopicentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AvailabletopicentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletopicentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
