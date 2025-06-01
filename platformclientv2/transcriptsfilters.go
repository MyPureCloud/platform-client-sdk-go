package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptsfilters
type Transcriptsfilters struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MediaType - The media type of the transcripts, default value is all 
	MediaType *string `json:"mediaType,omitempty"`

	// StartTimeMs - start time to filter by, default value is 7 days into the past
	StartTimeMs *int `json:"startTimeMs,omitempty"`

	// EndTimeMs - end time to filter by, default value is current time
	EndTimeMs *int `json:"endTimeMs,omitempty"`

	// Queues - list of queues ids to filter by
	Queues *[]string `json:"queues,omitempty"`

	// Flows - list of flows ids to filter by
	Flows *[]string `json:"flows,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transcriptsfilters) SetField(field string, fieldValue interface{}) {
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

func (o Transcriptsfilters) MarshalJSON() ([]byte, error) {
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
	type Alias Transcriptsfilters
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		StartTimeMs *int `json:"startTimeMs,omitempty"`
		
		EndTimeMs *int `json:"endTimeMs,omitempty"`
		
		Queues *[]string `json:"queues,omitempty"`
		
		Flows *[]string `json:"flows,omitempty"`
		Alias
	}{ 
		MediaType: o.MediaType,
		
		StartTimeMs: o.StartTimeMs,
		
		EndTimeMs: o.EndTimeMs,
		
		Queues: o.Queues,
		
		Flows: o.Flows,
		Alias:    (Alias)(o),
	})
}

func (o *Transcriptsfilters) UnmarshalJSON(b []byte) error {
	var TranscriptsfiltersMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptsfiltersMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := TranscriptsfiltersMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if StartTimeMs, ok := TranscriptsfiltersMap["startTimeMs"].(float64); ok {
		StartTimeMsInt := int(StartTimeMs)
		o.StartTimeMs = &StartTimeMsInt
	}
	
	if EndTimeMs, ok := TranscriptsfiltersMap["endTimeMs"].(float64); ok {
		EndTimeMsInt := int(EndTimeMs)
		o.EndTimeMs = &EndTimeMsInt
	}
	
	if Queues, ok := TranscriptsfiltersMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Flows, ok := TranscriptsfiltersMap["flows"].([]interface{}); ok {
		FlowsString, _ := json.Marshal(Flows)
		json.Unmarshal(FlowsString, &o.Flows)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptsfilters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
