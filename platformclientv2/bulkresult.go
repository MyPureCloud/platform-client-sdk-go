package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresult
type Bulkresult struct { 
	// VarError - Error details if the operation failed.
	VarError *Bulkerror `json:"error,omitempty"`


	// Entity - The result of the operation if it succeeded.
	Entity *interface{} `json:"entity,omitempty"`

}

func (o *Bulkresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresult
	
	return json.Marshal(&struct { 
		VarError *Bulkerror `json:"error,omitempty"`
		
		Entity *interface{} `json:"entity,omitempty"`
		*Alias
	}{ 
		VarError: o.VarError,
		
		Entity: o.Entity,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkresult) UnmarshalJSON(b []byte) error {
	var BulkresultMap map[string]interface{}
	err := json.Unmarshal(b, &BulkresultMap)
	if err != nil {
		return err
	}
	
	if VarError, ok := BulkresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if Entity, ok := BulkresultMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
