package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignedwrapupcode
type Assignedwrapupcode struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`

	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`

	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`

	// DurationSeconds - The duration in seconds of the wrap-up segment.
	DurationSeconds *int `json:"durationSeconds,omitempty"`

	// EndTime - The timestamp when the wrap-up segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assignedwrapupcode) SetField(field string, fieldValue interface{}) {
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

func (o Assignedwrapupcode) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EndTime", }
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
	type Alias Assignedwrapupcode
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		Alias
	}{ 
		Code: o.Code,
		
		Notes: o.Notes,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		EndTime: EndTime,
		Alias:    (Alias)(o),
	})
}

func (o *Assignedwrapupcode) UnmarshalJSON(b []byte) error {
	var AssignedwrapupcodeMap map[string]interface{}
	err := json.Unmarshal(b, &AssignedwrapupcodeMap)
	if err != nil {
		return err
	}
	
	if Code, ok := AssignedwrapupcodeMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Notes, ok := AssignedwrapupcodeMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if Tags, ok := AssignedwrapupcodeMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := AssignedwrapupcodeMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if endTimeString, ok := AssignedwrapupcodeMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignedwrapupcode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
