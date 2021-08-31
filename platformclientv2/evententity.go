package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evententity
type Evententity struct { 
	// EntityType - Type of entity the event pertains to. e.g. integration
	EntityType *string `json:"entityType,omitempty"`


	// Id - ID of the entity the event pertains to.
	Id *string `json:"id,omitempty"`

}

func (o *Evententity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evententity
	
	return json.Marshal(&struct { 
		EntityType *string `json:"entityType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		EntityType: o.EntityType,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Evententity) UnmarshalJSON(b []byte) error {
	var EvententityMap map[string]interface{}
	err := json.Unmarshal(b, &EvententityMap)
	if err != nil {
		return err
	}
	
	if EntityType, ok := EvententityMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
	
	if Id, ok := EvententityMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
