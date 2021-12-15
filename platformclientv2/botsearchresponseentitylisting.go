package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botsearchresponseentitylisting
type Botsearchresponseentitylisting struct { 
	// Entities
	Entities *[]Botsearchresponse `json:"entities,omitempty"`

}

func (o *Botsearchresponseentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botsearchresponseentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Botsearchresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Botsearchresponseentitylisting) UnmarshalJSON(b []byte) error {
	var BotsearchresponseentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &BotsearchresponseentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BotsearchresponseentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botsearchresponseentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
