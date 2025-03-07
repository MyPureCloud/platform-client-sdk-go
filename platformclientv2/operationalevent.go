package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Operationalevent
type Operationalevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventDefinition - The event that occurred.
	EventDefinition *Addressableentityref `json:"eventDefinition,omitempty"`

	// EntityId - The unique identifier for the entity
	EntityId *string `json:"entityId,omitempty"`

	// EntityToken - A token representing the entity
	EntityToken *string `json:"entityToken,omitempty"`

	// EntityName - The name for the entity
	EntityName *string `json:"entityName,omitempty"`

	// PreviousValue - The value prior to the event
	PreviousValue *string `json:"previousValue,omitempty"`

	// CurrentValue - The changed value following the event
	CurrentValue *string `json:"currentValue,omitempty"`

	// ErrorCode - The error code of the entity in the providing service associated to the event
	ErrorCode *string `json:"errorCode,omitempty"`

	// ParentEntityId - The unique identifier for the parent of the entity
	ParentEntityId *string `json:"parentEntityId,omitempty"`

	// Conversation - The link to a conversation
	Conversation *Addressableentityref `json:"conversation,omitempty"`

	// DateCreated - The date when the event created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// EntityVersion - The version of the entity in the providing service
	EntityVersion *string `json:"entityVersion,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Operationalevent) SetField(field string, fieldValue interface{}) {
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

func (o Operationalevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Operationalevent
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		EventDefinition *Addressableentityref `json:"eventDefinition,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityToken *string `json:"entityToken,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		PreviousValue *string `json:"previousValue,omitempty"`
		
		CurrentValue *string `json:"currentValue,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ParentEntityId *string `json:"parentEntityId,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		EntityVersion *string `json:"entityVersion,omitempty"`
		Alias
	}{ 
		EventDefinition: o.EventDefinition,
		
		EntityId: o.EntityId,
		
		EntityToken: o.EntityToken,
		
		EntityName: o.EntityName,
		
		PreviousValue: o.PreviousValue,
		
		CurrentValue: o.CurrentValue,
		
		ErrorCode: o.ErrorCode,
		
		ParentEntityId: o.ParentEntityId,
		
		Conversation: o.Conversation,
		
		DateCreated: DateCreated,
		
		EntityVersion: o.EntityVersion,
		Alias:    (Alias)(o),
	})
}

func (o *Operationalevent) UnmarshalJSON(b []byte) error {
	var OperationaleventMap map[string]interface{}
	err := json.Unmarshal(b, &OperationaleventMap)
	if err != nil {
		return err
	}
	
	if EventDefinition, ok := OperationaleventMap["eventDefinition"].(map[string]interface{}); ok {
		EventDefinitionString, _ := json.Marshal(EventDefinition)
		json.Unmarshal(EventDefinitionString, &o.EventDefinition)
	}
	
	if EntityId, ok := OperationaleventMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityToken, ok := OperationaleventMap["entityToken"].(string); ok {
		o.EntityToken = &EntityToken
	}
    
	if EntityName, ok := OperationaleventMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if PreviousValue, ok := OperationaleventMap["previousValue"].(string); ok {
		o.PreviousValue = &PreviousValue
	}
    
	if CurrentValue, ok := OperationaleventMap["currentValue"].(string); ok {
		o.CurrentValue = &CurrentValue
	}
    
	if ErrorCode, ok := OperationaleventMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ParentEntityId, ok := OperationaleventMap["parentEntityId"].(string); ok {
		o.ParentEntityId = &ParentEntityId
	}
    
	if Conversation, ok := OperationaleventMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if dateCreatedString, ok := OperationaleventMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if EntityVersion, ok := OperationaleventMap["entityVersion"].(string); ok {
		o.EntityVersion = &EntityVersion
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Operationalevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
