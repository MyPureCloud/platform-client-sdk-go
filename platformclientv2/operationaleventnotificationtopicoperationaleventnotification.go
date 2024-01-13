package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Operationaleventnotificationtopicoperationaleventnotification
type Operationaleventnotificationtopicoperationaleventnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventEntity
	EventEntity *Operationaleventnotificationtopicevententity `json:"eventEntity,omitempty"`

	// EntityId
	EntityId *string `json:"entityId,omitempty"`

	// EntityName
	EntityName *string `json:"entityName,omitempty"`

	// PreviousValue
	PreviousValue *string `json:"previousValue,omitempty"`

	// CurrentValue
	CurrentValue *string `json:"currentValue,omitempty"`

	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

	// Version
	Version *string `json:"version,omitempty"`

	// ParentEntity
	ParentEntity *string `json:"parentEntity,omitempty"`

	// EntityType
	EntityType *string `json:"entityType,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// Timestamp
	Timestamp *int `json:"timestamp,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Operationaleventnotificationtopicoperationaleventnotification) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Operationaleventnotificationtopicoperationaleventnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Operationaleventnotificationtopicoperationaleventnotification
	
	return json.Marshal(&struct { 
		EventEntity *Operationaleventnotificationtopicevententity `json:"eventEntity,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		PreviousValue *string `json:"previousValue,omitempty"`
		
		CurrentValue *string `json:"currentValue,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		ParentEntity *string `json:"parentEntity,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		Timestamp *int `json:"timestamp,omitempty"`
		Alias
	}{ 
		EventEntity: o.EventEntity,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		
		PreviousValue: o.PreviousValue,
		
		CurrentValue: o.CurrentValue,
		
		ErrorCode: o.ErrorCode,
		
		Version: o.Version,
		
		ParentEntity: o.ParentEntity,
		
		EntityType: o.EntityType,
		
		ConversationId: o.ConversationId,
		
		Timestamp: o.Timestamp,
		Alias:    (Alias)(o),
	})
}

func (o *Operationaleventnotificationtopicoperationaleventnotification) UnmarshalJSON(b []byte) error {
	var OperationaleventnotificationtopicoperationaleventnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &OperationaleventnotificationtopicoperationaleventnotificationMap)
	if err != nil {
		return err
	}
	
	if EventEntity, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["eventEntity"].(map[string]interface{}); ok {
		EventEntityString, _ := json.Marshal(EventEntity)
		json.Unmarshal(EventEntityString, &o.EventEntity)
	}
	
	if EntityId, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if PreviousValue, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["previousValue"].(string); ok {
		o.PreviousValue = &PreviousValue
	}
    
	if CurrentValue, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["currentValue"].(string); ok {
		o.CurrentValue = &CurrentValue
	}
    
	if ErrorCode, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if Version, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if ParentEntity, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["parentEntity"].(string); ok {
		o.ParentEntity = &ParentEntity
	}
    
	if EntityType, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if ConversationId, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if Timestamp, ok := OperationaleventnotificationtopicoperationaleventnotificationMap["timestamp"].(float64); ok {
		TimestampInt := int(Timestamp)
		o.Timestamp = &TimestampInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Operationaleventnotificationtopicoperationaleventnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
