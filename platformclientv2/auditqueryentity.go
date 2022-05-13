package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryentity
type Auditqueryentity struct { 
	// Name - Name of the Entity
	Name *string `json:"name,omitempty"`


	// Actions - List of Actions
	Actions *[]string `json:"actions,omitempty"`

}

func (o *Auditqueryentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryentity
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Actions *[]string `json:"actions,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryentity) UnmarshalJSON(b []byte) error {
	var AuditqueryentityMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryentityMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AuditqueryentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Actions, ok := AuditqueryentityMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
