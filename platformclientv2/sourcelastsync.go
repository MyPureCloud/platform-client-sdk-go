package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sourcelastsync
type Sourcelastsync struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - State of the last synchronization.
	State *string `json:"state,omitempty"`

	// DateStarted - Last synchronization start-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// DateEnded - Last synchronization end-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnded *time.Time `json:"dateEnded,omitempty"`

	// VarError - Optional error message of the last synchronization.
	VarError *Errorbody `json:"error,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sourcelastsync) SetField(field string, fieldValue interface{}) {
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

func (o Sourcelastsync) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStarted","DateEnded", }
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
	type Alias Sourcelastsync
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateEnded := new(string)
	if o.DateEnded != nil {
		
		*DateEnded = timeutil.Strftime(o.DateEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnded = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateEnded *string `json:"dateEnded,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		DateStarted: DateStarted,
		
		DateEnded: DateEnded,
		
		VarError: o.VarError,
		Alias:    (Alias)(o),
	})
}

func (o *Sourcelastsync) UnmarshalJSON(b []byte) error {
	var SourcelastsyncMap map[string]interface{}
	err := json.Unmarshal(b, &SourcelastsyncMap)
	if err != nil {
		return err
	}
	
	if State, ok := SourcelastsyncMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateStartedString, ok := SourcelastsyncMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateEndedString, ok := SourcelastsyncMap["dateEnded"].(string); ok {
		DateEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndedString)
		o.DateEnded = &DateEnded
	}
	
	if VarError, ok := SourcelastsyncMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sourcelastsync) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
