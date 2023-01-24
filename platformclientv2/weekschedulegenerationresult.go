package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulegenerationresult
type Weekschedulegenerationresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Failed - Whether the schedule generation failed
	Failed *bool `json:"failed,omitempty"`

	// RunId - ID of the schedule run
	RunId *string `json:"runId,omitempty"`

	// AgentWarnings - Warning messages from the schedule run. This will be available only when requesting information for a single week schedule
	AgentWarnings *[]Schedulegenerationwarning `json:"agentWarnings,omitempty"`

	// AgentWarningCount - Count of warning messages from the schedule run. This will be available only when requesting multiple week schedules
	AgentWarningCount *int `json:"agentWarningCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Weekschedulegenerationresult) SetField(field string, fieldValue interface{}) {
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

func (o Weekschedulegenerationresult) MarshalJSON() ([]byte, error) {
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
	type Alias Weekschedulegenerationresult
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		AgentWarnings *[]Schedulegenerationwarning `json:"agentWarnings,omitempty"`
		
		AgentWarningCount *int `json:"agentWarningCount,omitempty"`
		Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		AgentWarnings: o.AgentWarnings,
		
		AgentWarningCount: o.AgentWarningCount,
		Alias:    (Alias)(o),
	})
}

func (o *Weekschedulegenerationresult) UnmarshalJSON(b []byte) error {
	var WeekschedulegenerationresultMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulegenerationresultMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := WeekschedulegenerationresultMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
    
	if RunId, ok := WeekschedulegenerationresultMap["runId"].(string); ok {
		o.RunId = &RunId
	}
    
	if AgentWarnings, ok := WeekschedulegenerationresultMap["agentWarnings"].([]interface{}); ok {
		AgentWarningsString, _ := json.Marshal(AgentWarnings)
		json.Unmarshal(AgentWarningsString, &o.AgentWarnings)
	}
	
	if AgentWarningCount, ok := WeekschedulegenerationresultMap["agentWarningCount"].(float64); ok {
		AgentWarningCountInt := int(AgentWarningCount)
		o.AgentWarningCount = &AgentWarningCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulegenerationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
