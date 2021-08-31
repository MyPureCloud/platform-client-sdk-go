package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicdetail
type Queueconversationcobrowseeventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Queueconversationcobrowseeventtopicdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicdetail
	
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

func (o *Queueconversationcobrowseeventtopicdetail) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicdetailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicdetailMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := QueueconversationcobrowseeventtopicdetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if FieldName, ok := QueueconversationcobrowseeventtopicdetailMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
	
	if EntityId, ok := QueueconversationcobrowseeventtopicdetailMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
	
	if EntityName, ok := QueueconversationcobrowseeventtopicdetailMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
