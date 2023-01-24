package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflowoutcome
type Analyticsflowoutcome struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FlowOutcome - Combination of unique flow outcome identifier and its value separated by colon
	FlowOutcome *string `json:"flowOutcome,omitempty"`

	// FlowOutcomeEndTimestamp - The outcome ending timestamp in ISO 8601 format. This may be null if the outcome did not succeed.
	FlowOutcomeEndTimestamp *time.Time `json:"flowOutcomeEndTimestamp,omitempty"`

	// FlowOutcomeId - Unique identifier of a flow outcome
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`

	// FlowOutcomeStartTimestamp - The outcome starting timestamp in ISO 8601 format
	FlowOutcomeStartTimestamp *time.Time `json:"flowOutcomeStartTimestamp,omitempty"`

	// FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsflowoutcome) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsflowoutcome) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "FlowOutcomeEndTimestamp","FlowOutcomeStartTimestamp", }
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
	type Alias Analyticsflowoutcome
	
	FlowOutcomeEndTimestamp := new(string)
	if o.FlowOutcomeEndTimestamp != nil {
		
		*FlowOutcomeEndTimestamp = timeutil.Strftime(o.FlowOutcomeEndTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		FlowOutcomeEndTimestamp = nil
	}
	
	FlowOutcomeStartTimestamp := new(string)
	if o.FlowOutcomeStartTimestamp != nil {
		
		*FlowOutcomeStartTimestamp = timeutil.Strftime(o.FlowOutcomeStartTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		FlowOutcomeStartTimestamp = nil
	}
	
	return json.Marshal(&struct { 
		FlowOutcome *string `json:"flowOutcome,omitempty"`
		
		FlowOutcomeEndTimestamp *string `json:"flowOutcomeEndTimestamp,omitempty"`
		
		FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`
		
		FlowOutcomeStartTimestamp *string `json:"flowOutcomeStartTimestamp,omitempty"`
		
		FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
		Alias
	}{ 
		FlowOutcome: o.FlowOutcome,
		
		FlowOutcomeEndTimestamp: FlowOutcomeEndTimestamp,
		
		FlowOutcomeId: o.FlowOutcomeId,
		
		FlowOutcomeStartTimestamp: FlowOutcomeStartTimestamp,
		
		FlowOutcomeValue: o.FlowOutcomeValue,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsflowoutcome) UnmarshalJSON(b []byte) error {
	var AnalyticsflowoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsflowoutcomeMap)
	if err != nil {
		return err
	}
	
	if FlowOutcome, ok := AnalyticsflowoutcomeMap["flowOutcome"].(string); ok {
		o.FlowOutcome = &FlowOutcome
	}
    
	if flowOutcomeEndTimestampString, ok := AnalyticsflowoutcomeMap["flowOutcomeEndTimestamp"].(string); ok {
		FlowOutcomeEndTimestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", flowOutcomeEndTimestampString)
		o.FlowOutcomeEndTimestamp = &FlowOutcomeEndTimestamp
	}
	
	if FlowOutcomeId, ok := AnalyticsflowoutcomeMap["flowOutcomeId"].(string); ok {
		o.FlowOutcomeId = &FlowOutcomeId
	}
    
	if flowOutcomeStartTimestampString, ok := AnalyticsflowoutcomeMap["flowOutcomeStartTimestamp"].(string); ok {
		FlowOutcomeStartTimestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", flowOutcomeStartTimestampString)
		o.FlowOutcomeStartTimestamp = &FlowOutcomeStartTimestamp
	}
	
	if FlowOutcomeValue, ok := AnalyticsflowoutcomeMap["flowOutcomeValue"].(string); ok {
		o.FlowOutcomeValue = &FlowOutcomeValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
