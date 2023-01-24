package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrorrelationship
type Bulkerrorrelationship struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Code
	Code *string `json:"code,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// Status
	Status *int `json:"status,omitempty"`

	// Retryable
	Retryable *bool `json:"retryable,omitempty"`

	// Entity
	Entity *Relationship `json:"entity,omitempty"`

	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bulkerrorrelationship) SetField(field string, fieldValue interface{}) {
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

func (o Bulkerrorrelationship) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Bulkerrorrelationship
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		Retryable *bool `json:"retryable,omitempty"`
		
		Entity *Relationship `json:"entity,omitempty"`
		
		Details *[]Bulkerrordetail `json:"details,omitempty"`
		Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		
		Status: o.Status,
		
		Retryable: o.Retryable,
		
		Entity: o.Entity,
		
		Details: o.Details,
		Alias:    (Alias)(o),
	})
}

func (o *Bulkerrorrelationship) UnmarshalJSON(b []byte) error {
	var BulkerrorrelationshipMap map[string]interface{}
	err := json.Unmarshal(b, &BulkerrorrelationshipMap)
	if err != nil {
		return err
	}
	
	if Code, ok := BulkerrorrelationshipMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Message, ok := BulkerrorrelationshipMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Status, ok := BulkerrorrelationshipMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Retryable, ok := BulkerrorrelationshipMap["retryable"].(bool); ok {
		o.Retryable = &Retryable
	}
    
	if Entity, ok := BulkerrorrelationshipMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Details, ok := BulkerrorrelationshipMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkerrorrelationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
