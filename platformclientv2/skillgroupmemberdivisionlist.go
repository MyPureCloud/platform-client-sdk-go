package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgroupmemberdivisionlist
type Skillgroupmemberdivisionlist struct { 
	// Entities
	Entities *[]Division `json:"entities,omitempty"`

}

func (o *Skillgroupmemberdivisionlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgroupmemberdivisionlist
	
	return json.Marshal(&struct { 
		Entities *[]Division `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgroupmemberdivisionlist) UnmarshalJSON(b []byte) error {
	var SkillgroupmemberdivisionlistMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgroupmemberdivisionlistMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := SkillgroupmemberdivisionlistMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgroupmemberdivisionlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
