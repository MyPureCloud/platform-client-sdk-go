package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedmappingpreviewlisting
type Agentownedmappingpreviewlisting struct { 
	// Entities
	Entities *[]Agentownedmappingpreview `json:"entities,omitempty"`

}

func (o *Agentownedmappingpreviewlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentownedmappingpreviewlisting
	
	return json.Marshal(&struct { 
		Entities *[]Agentownedmappingpreview `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentownedmappingpreviewlisting) UnmarshalJSON(b []byte) error {
	var AgentownedmappingpreviewlistingMap map[string]interface{}
	err := json.Unmarshal(b, &AgentownedmappingpreviewlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AgentownedmappingpreviewlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreviewlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
