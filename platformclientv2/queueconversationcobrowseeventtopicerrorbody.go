package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicerrorbody
type Queueconversationcobrowseeventtopicerrorbody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Message
	Message *string `json:"message,omitempty"`

	// Code
	Code *string `json:"code,omitempty"`

	// Status
	Status *int `json:"status,omitempty"`

	// EntityId
	EntityId *string `json:"entityId,omitempty"`

	// EntityName
	EntityName *string `json:"entityName,omitempty"`

	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`

	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`

	// ContextId
	ContextId *string `json:"contextId,omitempty"`

	// Details
	Details *[]Queueconversationcobrowseeventtopicdetail `json:"details,omitempty"`

	// Errors
	Errors *[]Queueconversationcobrowseeventtopicerrorbody `json:"errors,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationcobrowseeventtopicerrorbody) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationcobrowseeventtopicerrorbody) MarshalJSON() ([]byte, error) {
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
	type Alias Queueconversationcobrowseeventtopicerrorbody
	
	return json.Marshal(&struct { 
		Message *string `json:"message,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Details *[]Queueconversationcobrowseeventtopicdetail `json:"details,omitempty"`
		
		Errors *[]Queueconversationcobrowseeventtopicerrorbody `json:"errors,omitempty"`
		Alias
	}{ 
		Message: o.Message,
		
		Code: o.Code,
		
		Status: o.Status,
		
		EntityId: o.EntityId,
		
		EntityName: o.EntityName,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		ContextId: o.ContextId,
		
		Details: o.Details,
		
		Errors: o.Errors,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationcobrowseeventtopicerrorbody) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicerrorbodyMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicerrorbodyMap)
	if err != nil {
		return err
	}
	
	if Message, ok := QueueconversationcobrowseeventtopicerrorbodyMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Code, ok := QueueconversationcobrowseeventtopicerrorbodyMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Status, ok := QueueconversationcobrowseeventtopicerrorbodyMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if EntityId, ok := QueueconversationcobrowseeventtopicerrorbodyMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityName, ok := QueueconversationcobrowseeventtopicerrorbodyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if MessageWithParams, ok := QueueconversationcobrowseeventtopicerrorbodyMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := QueueconversationcobrowseeventtopicerrorbodyMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ContextId, ok := QueueconversationcobrowseeventtopicerrorbodyMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Details, ok := QueueconversationcobrowseeventtopicerrorbodyMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	
	if Errors, ok := QueueconversationcobrowseeventtopicerrorbodyMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
