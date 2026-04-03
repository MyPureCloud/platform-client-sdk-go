package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Unprocessedexternalevent
type Unprocessedexternalevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Event - The event that failed processing.
	Event *Externalevent `json:"event,omitempty"`

	// OriginalRequestIndex - The index of the event in the original request.
	OriginalRequestIndex *int `json:"originalRequestIndex,omitempty"`

	// IsRetryable - Whether the error is retryable.
	IsRetryable *bool `json:"isRetryable,omitempty"`

	// ErrorMessage - The error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// StatusCode - The HTTP status code associated with the error.
	StatusCode *int `json:"statusCode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Unprocessedexternalevent) SetField(field string, fieldValue interface{}) {
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

func (o Unprocessedexternalevent) MarshalJSON() ([]byte, error) {
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
	type Alias Unprocessedexternalevent
	
	return json.Marshal(&struct { 
		Event *Externalevent `json:"event,omitempty"`
		
		OriginalRequestIndex *int `json:"originalRequestIndex,omitempty"`
		
		IsRetryable *bool `json:"isRetryable,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		StatusCode *int `json:"statusCode,omitempty"`
		Alias
	}{ 
		Event: o.Event,
		
		OriginalRequestIndex: o.OriginalRequestIndex,
		
		IsRetryable: o.IsRetryable,
		
		ErrorMessage: o.ErrorMessage,
		
		StatusCode: o.StatusCode,
		Alias:    (Alias)(o),
	})
}

func (o *Unprocessedexternalevent) UnmarshalJSON(b []byte) error {
	var UnprocessedexternaleventMap map[string]interface{}
	err := json.Unmarshal(b, &UnprocessedexternaleventMap)
	if err != nil {
		return err
	}
	
	if Event, ok := UnprocessedexternaleventMap["event"].(map[string]interface{}); ok {
		EventString, _ := json.Marshal(Event)
		json.Unmarshal(EventString, &o.Event)
	}
	
	if OriginalRequestIndex, ok := UnprocessedexternaleventMap["originalRequestIndex"].(float64); ok {
		OriginalRequestIndexInt := int(OriginalRequestIndex)
		o.OriginalRequestIndex = &OriginalRequestIndexInt
	}
	
	if IsRetryable, ok := UnprocessedexternaleventMap["isRetryable"].(bool); ok {
		o.IsRetryable = &IsRetryable
	}
    
	if ErrorMessage, ok := UnprocessedexternaleventMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if StatusCode, ok := UnprocessedexternaleventMap["statusCode"].(float64); ok {
		StatusCodeInt := int(StatusCode)
		o.StatusCode = &StatusCodeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unprocessedexternalevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
