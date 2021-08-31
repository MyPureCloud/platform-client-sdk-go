package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultexternalcontactexternalcontact
type Bulkresponseresultexternalcontactexternalcontact struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorexternalcontact `json:"error,omitempty"`

}

func (o *Bulkresponseresultexternalcontactexternalcontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultexternalcontactexternalcontact
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Externalcontact `json:"entity,omitempty"`
		
		VarError *Bulkerrorexternalcontact `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Success: o.Success,
		
		Entity: o.Entity,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresponseresultexternalcontactexternalcontact) UnmarshalJSON(b []byte) error {
	var BulkresponseresultexternalcontactexternalcontactMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresponseresultexternalcontactexternalcontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkresponseresultexternalcontactexternalcontactMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Success, ok := BulkresponseresultexternalcontactexternalcontactMap["success"].(bool); ok {
		o.Success = &Success
	}
	
	if Entity, ok := BulkresponseresultexternalcontactexternalcontactMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if VarError, ok := BulkresponseresultexternalcontactexternalcontactMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactexternalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
