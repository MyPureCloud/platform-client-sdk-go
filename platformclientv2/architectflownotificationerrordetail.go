package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationerrordetail
type Architectflownotificationerrordetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`

}

func (o *Architectflownotificationerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationerrordetail
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		FieldName *string `json:"fieldName,omitempty"`
		*Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		
		FieldName: o.FieldName,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationerrordetail) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationerrordetailMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationerrordetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := ArchitectflownotificationerrordetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if EntityId, ok := ArchitectflownotificationerrordetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := ArchitectflownotificationerrordetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if FieldName, ok := ArchitectflownotificationerrordetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
