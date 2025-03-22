package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Aiscoring
type Aiscoring struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FailureType - The type of error that occurred while processing AI scores. It is null where there is no error.
	FailureType *string `json:"failureType,omitempty"`

	// Pending - Indicates whether AI scoring is currently processing the evaluation.
	Pending *bool `json:"pending,omitempty"`

	// DateLastChanged - The date when the AI scores were last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastChanged *time.Time `json:"dateLastChanged,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Aiscoring) SetField(field string, fieldValue interface{}) {
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

func (o Aiscoring) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastChanged", }
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
	type Alias Aiscoring
	
	DateLastChanged := new(string)
	if o.DateLastChanged != nil {
		
		*DateLastChanged = timeutil.Strftime(o.DateLastChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastChanged = nil
	}
	
	return json.Marshal(&struct { 
		FailureType *string `json:"failureType,omitempty"`
		
		Pending *bool `json:"pending,omitempty"`
		
		DateLastChanged *string `json:"dateLastChanged,omitempty"`
		Alias
	}{ 
		FailureType: o.FailureType,
		
		Pending: o.Pending,
		
		DateLastChanged: DateLastChanged,
		Alias:    (Alias)(o),
	})
}

func (o *Aiscoring) UnmarshalJSON(b []byte) error {
	var AiscoringMap map[string]interface{}
	err := json.Unmarshal(b, &AiscoringMap)
	if err != nil {
		return err
	}
	
	if FailureType, ok := AiscoringMap["failureType"].(string); ok {
		o.FailureType = &FailureType
	}
    
	if Pending, ok := AiscoringMap["pending"].(bool); ok {
		o.Pending = &Pending
	}
    
	if dateLastChangedString, ok := AiscoringMap["dateLastChanged"].(string); ok {
		DateLastChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastChangedString)
		o.DateLastChanged = &DateLastChanged
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Aiscoring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
