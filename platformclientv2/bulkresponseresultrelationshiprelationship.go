package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultrelationshiprelationship
type Bulkresponseresultrelationshiprelationship struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Relationship `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorrelationship `json:"error,omitempty"`

}

func (o *Bulkresponseresultrelationshiprelationship) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultrelationshiprelationship
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Relationship `json:"entity,omitempty"`
		
		VarError *Bulkerrorrelationship `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Success: o.Success,
		
		Entity: o.Entity,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresponseresultrelationshiprelationship) UnmarshalJSON(b []byte) error {
	var BulkresponseresultrelationshiprelationshipMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresponseresultrelationshiprelationshipMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkresponseresultrelationshiprelationshipMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Success, ok := BulkresponseresultrelationshiprelationshipMap["success"].(bool); ok {
		o.Success = &Success
	}
	
	if Entity, ok := BulkresponseresultrelationshiprelationshipMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if VarError, ok := BulkresponseresultrelationshiprelationshipMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultrelationshiprelationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
