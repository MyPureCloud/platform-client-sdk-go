package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicdetail
type Conversationcallbackeventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Conversationcallbackeventtopicdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicdetail
	
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

func (o *Conversationcallbackeventtopicdetail) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicdetailMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicdetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := ConversationcallbackeventtopicdetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if FieldName, ok := ConversationcallbackeventtopicdetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
	
	if EntityId, ok := ConversationcallbackeventtopicdetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
	
	if EntityName, ok := ConversationcallbackeventtopicdetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
