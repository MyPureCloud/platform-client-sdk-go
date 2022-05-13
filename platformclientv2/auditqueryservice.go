package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryservice
type Auditqueryservice struct { 
	// Name - Name of the Service
	Name *string `json:"name,omitempty"`


	// Entities - List of Entities
	Entities *[]Auditqueryentity `json:"entities,omitempty"`

}

func (o *Auditqueryservice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryservice
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Entities *[]Auditqueryentity `json:"entities,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryservice) UnmarshalJSON(b []byte) error {
	var AuditqueryserviceMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryserviceMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AuditqueryserviceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Entities, ok := AuditqueryserviceMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryservice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
