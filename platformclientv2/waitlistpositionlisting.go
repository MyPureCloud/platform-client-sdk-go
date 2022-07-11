package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Waitlistpositionlisting
type Waitlistpositionlisting struct { 
	// Entities
	Entities *[]Waitlistposition `json:"entities,omitempty"`

}

func (o *Waitlistpositionlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Waitlistpositionlisting
	
	return json.Marshal(&struct { 
		Entities *[]Waitlistposition `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Waitlistpositionlisting) UnmarshalJSON(b []byte) error {
	var WaitlistpositionlistingMap map[string]interface{}
	err := json.Unmarshal(b, &WaitlistpositionlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WaitlistpositionlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Waitlistpositionlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
