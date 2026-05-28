package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externaleventchange - A change in an external event definition
type Externaleventchange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ChangeCategory - The category of the change
	ChangeCategory *string `json:"changeCategory,omitempty"`

	// SchemaId - The unique identifier for the schema
	SchemaId *string `json:"schemaId,omitempty"`

	// EventName - The name of the event
	EventName *string `json:"eventName,omitempty"`

	// DateDetected - The timestamp when the change was detected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDetected *time.Time `json:"dateDetected,omitempty"`

	// SystemStatus - The status of the change
	SystemStatus *string `json:"systemStatus,omitempty"`

	// ErrorCode - A code representing the error, only present for ERROR category changes
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorDescription - A description of the error, only present for ERROR category changes
	ErrorDescription *string `json:"errorDescription,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externaleventchange) SetField(field string, fieldValue interface{}) {
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

func (o Externaleventchange) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDetected", }
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
	type Alias Externaleventchange
	
	DateDetected := new(string)
	if o.DateDetected != nil {
		
		*DateDetected = timeutil.Strftime(o.DateDetected, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDetected = nil
	}
	
	return json.Marshal(&struct { 
		ChangeCategory *string `json:"changeCategory,omitempty"`
		
		SchemaId *string `json:"schemaId,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		DateDetected *string `json:"dateDetected,omitempty"`
		
		SystemStatus *string `json:"systemStatus,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorDescription *string `json:"errorDescription,omitempty"`
		Alias
	}{ 
		ChangeCategory: o.ChangeCategory,
		
		SchemaId: o.SchemaId,
		
		EventName: o.EventName,
		
		DateDetected: DateDetected,
		
		SystemStatus: o.SystemStatus,
		
		ErrorCode: o.ErrorCode,
		
		ErrorDescription: o.ErrorDescription,
		Alias:    (Alias)(o),
	})
}

func (o *Externaleventchange) UnmarshalJSON(b []byte) error {
	var ExternaleventchangeMap map[string]interface{}
	err := json.Unmarshal(b, &ExternaleventchangeMap)
	if err != nil {
		return err
	}
	
	if ChangeCategory, ok := ExternaleventchangeMap["changeCategory"].(string); ok {
		o.ChangeCategory = &ChangeCategory
	}
    
	if SchemaId, ok := ExternaleventchangeMap["schemaId"].(string); ok {
		o.SchemaId = &SchemaId
	}
    
	if EventName, ok := ExternaleventchangeMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if dateDetectedString, ok := ExternaleventchangeMap["dateDetected"].(string); ok {
		DateDetected, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDetectedString)
		o.DateDetected = &DateDetected
	}
	
	if SystemStatus, ok := ExternaleventchangeMap["systemStatus"].(string); ok {
		o.SystemStatus = &SystemStatus
	}
    
	if ErrorCode, ok := ExternaleventchangeMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorDescription, ok := ExternaleventchangeMap["errorDescription"].(string); ok {
		o.ErrorDescription = &ErrorDescription
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externaleventchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
