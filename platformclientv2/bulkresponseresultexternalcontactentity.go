package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultexternalcontactentity
type Bulkresponseresultexternalcontactentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

func (o *Bulkresponseresultexternalcontactentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultexternalcontactentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Externalcontact `json:"entity,omitempty"`
		
		VarError *Bulkerrorentity `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Success: o.Success,
		
		Entity: o.Entity,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresponseresultexternalcontactentity) UnmarshalJSON(b []byte) error {
	var BulkresponseresultexternalcontactentityMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresponseresultexternalcontactentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkresponseresultexternalcontactentityMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Success, ok := BulkresponseresultexternalcontactentityMap["success"].(bool); ok {
		o.Success = &Success
	}
	
	if Entity, ok := BulkresponseresultexternalcontactentityMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if VarError, ok := BulkresponseresultexternalcontactentityMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
