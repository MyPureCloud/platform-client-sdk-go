package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditchange
type Auditchange struct { 
	// Property
	Property *string `json:"property,omitempty"`


	// Entity
	Entity *Auditentityreference `json:"entity,omitempty"`


	// OldValues
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues
	NewValues *[]string `json:"newValues,omitempty"`

}

func (o *Auditchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditchange
	
	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		Entity *Auditentityreference `json:"entity,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Property: o.Property,
		
		Entity: o.Entity,
		
		OldValues: o.OldValues,
		
		NewValues: o.NewValues,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditchange) UnmarshalJSON(b []byte) error {
	var AuditchangeMap map[string]interface{}
	err := json.Unmarshal(b, &AuditchangeMap)
	if err != nil {
		return err
	}
	
	if Property, ok := AuditchangeMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if Entity, ok := AuditchangeMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if OldValues, ok := AuditchangeMap["oldValues"].([]interface{}); ok {
		OldValuesString, _ := json.Marshal(OldValues)
		json.Unmarshal(OldValuesString, &o.OldValues)
	}
	
	if NewValues, ok := AuditchangeMap["newValues"].([]interface{}); ok {
		NewValuesString, _ := json.Marshal(NewValues)
		json.Unmarshal(NewValuesString, &o.NewValues)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
