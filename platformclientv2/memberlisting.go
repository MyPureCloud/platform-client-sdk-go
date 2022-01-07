package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Memberlisting
type Memberlisting struct { 
	// Entities
	Entities *[]Member `json:"entities,omitempty"`

}

func (o *Memberlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Memberlisting
	
	return json.Marshal(&struct { 
		Entities *[]Member `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Memberlisting) UnmarshalJSON(b []byte) error {
	var MemberlistingMap map[string]interface{}
	err := json.Unmarshal(b, &MemberlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := MemberlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Memberlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
