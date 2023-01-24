package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationresultsummary
type Schedulegenerationresultsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Failed - Whether the schedule generation run failed
	Failed *bool `json:"failed,omitempty"`

	// RunId - The ID of the schedule generation run. Reference this when requesting support
	RunId *string `json:"runId,omitempty"`

	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`

	// MessageSeverityCounts - The list of schedule generation message counts by severity for this schedule generation run
	MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulegenerationresultsummary) SetField(field string, fieldValue interface{}) {
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

func (o Schedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulegenerationresultsummary
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		
		MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
		Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		MessageCount: o.MessageCount,
		
		MessageSeverityCounts: o.MessageSeverityCounts,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulegenerationresultsummary) UnmarshalJSON(b []byte) error {
	var SchedulegenerationresultsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegenerationresultsummaryMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := SchedulegenerationresultsummaryMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
    
	if RunId, ok := SchedulegenerationresultsummaryMap["runId"].(string); ok {
		o.RunId = &RunId
	}
    
	if MessageCount, ok := SchedulegenerationresultsummaryMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if MessageSeverityCounts, ok := SchedulegenerationresultsummaryMap["messageSeverityCounts"].([]interface{}); ok {
		MessageSeverityCountsString, _ := json.Marshal(MessageSeverityCounts)
		json.Unmarshal(MessageSeverityCountsString, &o.MessageSeverityCounts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
