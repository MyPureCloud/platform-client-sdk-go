package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupsettings
type Groupsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MinimumGroupSize - The minimum size of a group for a session
	MinimumGroupSize *int `json:"minimumGroupSize,omitempty"`

	// MaximumGroupSize - The maximum size of a group for a session
	MaximumGroupSize *int `json:"maximumGroupSize,omitempty"`

	// MaximumTotalSessions - The maximum total number of sessions
	MaximumTotalSessions *int `json:"maximumTotalSessions,omitempty"`

	// MaximumConcurrentSessions - The maximum number of sessions that can be scheduled concurrently
	MaximumConcurrentSessions *int `json:"maximumConcurrentSessions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Groupsettings) SetField(field string, fieldValue interface{}) {
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

func (o Groupsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Groupsettings
	
	return json.Marshal(&struct { 
		MinimumGroupSize *int `json:"minimumGroupSize,omitempty"`
		
		MaximumGroupSize *int `json:"maximumGroupSize,omitempty"`
		
		MaximumTotalSessions *int `json:"maximumTotalSessions,omitempty"`
		
		MaximumConcurrentSessions *int `json:"maximumConcurrentSessions,omitempty"`
		Alias
	}{ 
		MinimumGroupSize: o.MinimumGroupSize,
		
		MaximumGroupSize: o.MaximumGroupSize,
		
		MaximumTotalSessions: o.MaximumTotalSessions,
		
		MaximumConcurrentSessions: o.MaximumConcurrentSessions,
		Alias:    (Alias)(o),
	})
}

func (o *Groupsettings) UnmarshalJSON(b []byte) error {
	var GroupsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &GroupsettingsMap)
	if err != nil {
		return err
	}
	
	if MinimumGroupSize, ok := GroupsettingsMap["minimumGroupSize"].(float64); ok {
		MinimumGroupSizeInt := int(MinimumGroupSize)
		o.MinimumGroupSize = &MinimumGroupSizeInt
	}
	
	if MaximumGroupSize, ok := GroupsettingsMap["maximumGroupSize"].(float64); ok {
		MaximumGroupSizeInt := int(MaximumGroupSize)
		o.MaximumGroupSize = &MaximumGroupSizeInt
	}
	
	if MaximumTotalSessions, ok := GroupsettingsMap["maximumTotalSessions"].(float64); ok {
		MaximumTotalSessionsInt := int(MaximumTotalSessions)
		o.MaximumTotalSessions = &MaximumTotalSessionsInt
	}
	
	if MaximumConcurrentSessions, ok := GroupsettingsMap["maximumConcurrentSessions"].(float64); ok {
		MaximumConcurrentSessionsInt := int(MaximumConcurrentSessions)
		o.MaximumConcurrentSessions = &MaximumConcurrentSessionsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Groupsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
