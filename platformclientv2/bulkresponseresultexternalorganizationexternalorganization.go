package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultexternalorganizationexternalorganization
type Bulkresponseresultexternalorganizationexternalorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalorganization `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorexternalorganization `json:"error,omitempty"`

}

func (o *Bulkresponseresultexternalorganizationexternalorganization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultexternalorganizationexternalorganization
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Externalorganization `json:"entity,omitempty"`
		
		VarError *Bulkerrorexternalorganization `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Success: o.Success,
		
		Entity: o.Entity,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresponseresultexternalorganizationexternalorganization) UnmarshalJSON(b []byte) error {
	var BulkresponseresultexternalorganizationexternalorganizationMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresponseresultexternalorganizationexternalorganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkresponseresultexternalorganizationexternalorganizationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Success, ok := BulkresponseresultexternalorganizationexternalorganizationMap["success"].(bool); ok {
		o.Success = &Success
	}
	
	if Entity, ok := BulkresponseresultexternalorganizationexternalorganizationMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if VarError, ok := BulkresponseresultexternalorganizationexternalorganizationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationexternalorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
