package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultnotenote
type Bulkresponseresultnotenote struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Note `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrornote `json:"error,omitempty"`

}

func (o *Bulkresponseresultnotenote) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultnotenote
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Note `json:"entity,omitempty"`
		
		VarError *Bulkerrornote `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Success: o.Success,
		
		Entity: o.Entity,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresponseresultnotenote) UnmarshalJSON(b []byte) error {
	var BulkresponseresultnotenoteMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresponseresultnotenoteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkresponseresultnotenoteMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Success, ok := BulkresponseresultnotenoteMap["success"].(bool); ok {
		o.Success = &Success
	}
	
	if Entity, ok := BulkresponseresultnotenoteMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if VarError, ok := BulkresponseresultnotenoteMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultnotenote) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
