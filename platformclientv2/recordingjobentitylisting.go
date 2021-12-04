package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobentitylisting
type Recordingjobentitylisting struct { 
	// Entities
	Entities *[]Recordingjob `json:"entities,omitempty"`

}

func (o *Recordingjobentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingjobentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Recordingjob `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingjobentitylisting) UnmarshalJSON(b []byte) error {
	var RecordingjobentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingjobentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := RecordingjobentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingjobentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
