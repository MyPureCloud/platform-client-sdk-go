package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detail
type Detail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Detail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detail
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		FieldName *string `json:"fieldName,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		*Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		FieldName: o.FieldName,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		Alias:    (*Alias)(o),
	})
}

func (o *Detail) UnmarshalJSON(b []byte) error {
	var DetailMap map[string]interface{}
	err := json.Unmarshal(b, &DetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := DetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if FieldName, ok := DetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
	
	if EntityId, ok := DetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
	
	if EntityName, ok := DetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Detail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
