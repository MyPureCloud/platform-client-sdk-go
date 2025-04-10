package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkentityerrorcontactenrichrequest
type Bulkentityerrorcontactenrichrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Code - An error code for the specific error condition.
	Code *string `json:"code,omitempty"`

	// Message - A short error message.
	Message *string `json:"message,omitempty"`

	// Status - The HTTP Status Code for the error.
	Status *int `json:"status,omitempty"`

	// Retryable - Whether this particular error should be retried.
	Retryable *bool `json:"retryable,omitempty"`

	// Details - Additional error details for specific fields.
	Details *[]Bulkerrordetail `json:"details,omitempty"`

	// Entity - The entity body specified in the Bulk request operation that caused this error.
	Entity *Contactenrichrequest `json:"entity,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bulkentityerrorcontactenrichrequest) SetField(field string, fieldValue interface{}) {
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

func (o Bulkentityerrorcontactenrichrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Bulkentityerrorcontactenrichrequest
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		Retryable *bool `json:"retryable,omitempty"`
		
		Details *[]Bulkerrordetail `json:"details,omitempty"`
		
		Entity *Contactenrichrequest `json:"entity,omitempty"`
		Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		
		Status: o.Status,
		
		Retryable: o.Retryable,
		
		Details: o.Details,
		
		Entity: o.Entity,
		Alias:    (Alias)(o),
	})
}

func (o *Bulkentityerrorcontactenrichrequest) UnmarshalJSON(b []byte) error {
	var BulkentityerrorcontactenrichrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkentityerrorcontactenrichrequestMap)
	if err != nil {
		return err
	}
	
	if Code, ok := BulkentityerrorcontactenrichrequestMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Message, ok := BulkentityerrorcontactenrichrequestMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Status, ok := BulkentityerrorcontactenrichrequestMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Retryable, ok := BulkentityerrorcontactenrichrequestMap["retryable"].(bool); ok {
		o.Retryable = &Retryable
	}
    
	if Details, ok := BulkentityerrorcontactenrichrequestMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	
	if Entity, ok := BulkentityerrorcontactenrichrequestMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkentityerrorcontactenrichrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
