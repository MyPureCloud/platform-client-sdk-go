package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Change
type Change struct { 
	// Entity
	Entity *Auditentity `json:"entity,omitempty"`


	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - The old values which were modified and/or removed by this action.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - The new values which were modified and/or added by this action.
	NewValues *[]string `json:"newValues,omitempty"`

}

func (o *Change) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Change
	
	return json.Marshal(&struct { 
		Entity *Auditentity `json:"entity,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		Entity: o.Entity,
		
		Property: o.Property,
		
		OldValues: o.OldValues,
		
		NewValues: o.NewValues,
		Alias:    (*Alias)(o),
	})
}

func (o *Change) UnmarshalJSON(b []byte) error {
	var ChangeMap map[string]interface{}
	err := json.Unmarshal(b, &ChangeMap)
	if err != nil {
		return err
	}
	
	if Entity, ok := ChangeMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Property, ok := ChangeMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if OldValues, ok := ChangeMap["oldValues"].([]interface{}); ok {
		OldValuesString, _ := json.Marshal(OldValues)
		json.Unmarshal(OldValuesString, &o.OldValues)
	}
	
	if NewValues, ok := ChangeMap["newValues"].([]interface{}); ok {
		NewValuesString, _ := json.Marshal(NewValues)
		json.Unmarshal(NewValuesString, &o.NewValues)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Change) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
