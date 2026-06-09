package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontableexportjoberror - Error details when a decision table export job fails
type Decisiontableexportjoberror struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ErrorCode - The error code for this job failure.
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessage - A human-readable error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// MessageWithParams - Parameterized message template for the aggregate failure (when applicable)
	MessageWithParams *string `json:"messageWithParams,omitempty"`

	// MessageParams - Parameters for messageWithParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`

	// ValidationErrors - Validation failures for the export job
	ValidationErrors *[]Decisiontablejobvalidationerror `json:"validationErrors,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontableexportjoberror) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontableexportjoberror) MarshalJSON() ([]byte, error) {
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
	type Alias Decisiontableexportjoberror
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		MessageWithParams *string `json:"messageWithParams,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		
		ValidationErrors *[]Decisiontablejobvalidationerror `json:"validationErrors,omitempty"`
		Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		
		MessageWithParams: o.MessageWithParams,
		
		MessageParams: o.MessageParams,
		
		ValidationErrors: o.ValidationErrors,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontableexportjoberror) UnmarshalJSON(b []byte) error {
	var DecisiontableexportjoberrorMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontableexportjoberrorMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := DecisiontableexportjoberrorMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := DecisiontableexportjoberrorMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if MessageWithParams, ok := DecisiontableexportjoberrorMap["messageWithParams"].(string); ok {
		o.MessageWithParams = &MessageWithParams
	}
    
	if MessageParams, ok := DecisiontableexportjoberrorMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	
	if ValidationErrors, ok := DecisiontableexportjoberrorMap["validationErrors"].([]interface{}); ok {
		ValidationErrorsString, _ := json.Marshal(ValidationErrors)
		json.Unmarshal(ValidationErrorsString, &o.ValidationErrors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontableexportjoberror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
