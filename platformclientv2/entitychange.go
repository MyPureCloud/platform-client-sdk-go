package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entitychange
type Entitychange struct { 
	// EntityId - Id of the entity that was changed
	EntityId *string `json:"entityId,omitempty"`


	// EntityName - Name of the entity that was changed
	EntityName *string `json:"entityName,omitempty"`


	// EntityType - Type of the entity that was changed
	EntityType *string `json:"entityType,omitempty"`


	// OldValues - Previous values for the entity.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - New values for the entity.
	NewValues *[]string `json:"newValues,omitempty"`

}

func (o *Entitychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entitychange
	
	return json.Marshal(&struct { 
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		OldValues *[]string `json:"oldValues,omitempty"`
		
		NewValues *[]string `json:"newValues,omitempty"`
		*Alias
	}{ 
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		
		EntityType: o.EntityType,
		
		OldValues: o.OldValues,
		
		NewValues: o.NewValues,
		Alias:    (*Alias)(o),
	})
}

func (o *Entitychange) UnmarshalJSON(b []byte) error {
	var EntitychangeMap map[string]interface{}
	err := json.Unmarshal(b, &EntitychangeMap)
	if err != nil {
		return err
	}
	
	if EntityId, ok := EntitychangeMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := EntitychangeMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if EntityType, ok := EntitychangeMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if OldValues, ok := EntitychangeMap["oldValues"].([]interface{}); ok {
		OldValuesString, _ := json.Marshal(OldValues)
		json.Unmarshal(OldValuesString, &o.OldValues)
	}
	
	if NewValues, ok := EntitychangeMap["newValues"].([]interface{}); ok {
		NewValuesString, _ := json.Marshal(NewValues)
		json.Unmarshal(NewValuesString, &o.NewValues)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Entitychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
