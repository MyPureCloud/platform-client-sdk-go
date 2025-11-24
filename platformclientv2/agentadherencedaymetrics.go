package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentadherencedaymetrics
type Agentadherencedaymetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DayStartOffsetSeconds - Start of day offset in seconds relative to query start time
	DayStartOffsetSeconds *int `json:"dayStartOffsetSeconds,omitempty"`

	// AdherenceScheduleSeconds - Duration of schedule in seconds included for adherence percentage calculation
	AdherenceScheduleSeconds *int `json:"adherenceScheduleSeconds,omitempty"`

	// ConformanceScheduleSeconds - Total scheduled duration in seconds for OnQueue activities
	ConformanceScheduleSeconds *int `json:"conformanceScheduleSeconds,omitempty"`

	// ConformanceActualSeconds - Total actually worked duration in seconds for OnQueue activities
	ConformanceActualSeconds *int `json:"conformanceActualSeconds,omitempty"`

	// ExceptionCount - Total number of adherence exceptions for this user
	ExceptionCount *int `json:"exceptionCount,omitempty"`

	// ExceptionDurationSeconds - Total duration in seconds of adherence exceptions for this user
	ExceptionDurationSeconds *int `json:"exceptionDurationSeconds,omitempty"`

	// ImpactSeconds - The impact duration in seconds of current adherence state for this user
	ImpactSeconds *int `json:"impactSeconds,omitempty"`

	// ScheduleLengthSeconds - Total duration in seconds for all scheduled activities
	ScheduleLengthSeconds *int `json:"scheduleLengthSeconds,omitempty"`

	// ActualLengthSeconds - Total duration in seconds for all actually worked activities
	ActualLengthSeconds *int `json:"actualLengthSeconds,omitempty"`

	// AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`

	// ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentadherencedaymetrics) SetField(field string, fieldValue interface{}) {
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

func (o Agentadherencedaymetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Agentadherencedaymetrics
	
	return json.Marshal(&struct { 
		DayStartOffsetSeconds *int `json:"dayStartOffsetSeconds,omitempty"`
		
		AdherenceScheduleSeconds *int `json:"adherenceScheduleSeconds,omitempty"`
		
		ConformanceScheduleSeconds *int `json:"conformanceScheduleSeconds,omitempty"`
		
		ConformanceActualSeconds *int `json:"conformanceActualSeconds,omitempty"`
		
		ExceptionCount *int `json:"exceptionCount,omitempty"`
		
		ExceptionDurationSeconds *int `json:"exceptionDurationSeconds,omitempty"`
		
		ImpactSeconds *int `json:"impactSeconds,omitempty"`
		
		ScheduleLengthSeconds *int `json:"scheduleLengthSeconds,omitempty"`
		
		ActualLengthSeconds *int `json:"actualLengthSeconds,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		Alias
	}{ 
		DayStartOffsetSeconds: o.DayStartOffsetSeconds,
		
		AdherenceScheduleSeconds: o.AdherenceScheduleSeconds,
		
		ConformanceScheduleSeconds: o.ConformanceScheduleSeconds,
		
		ConformanceActualSeconds: o.ConformanceActualSeconds,
		
		ExceptionCount: o.ExceptionCount,
		
		ExceptionDurationSeconds: o.ExceptionDurationSeconds,
		
		ImpactSeconds: o.ImpactSeconds,
		
		ScheduleLengthSeconds: o.ScheduleLengthSeconds,
		
		ActualLengthSeconds: o.ActualLengthSeconds,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		Alias:    (Alias)(o),
	})
}

func (o *Agentadherencedaymetrics) UnmarshalJSON(b []byte) error {
	var AgentadherencedaymetricsMap map[string]interface{}
	err := json.Unmarshal(b, &AgentadherencedaymetricsMap)
	if err != nil {
		return err
	}
	
	if DayStartOffsetSeconds, ok := AgentadherencedaymetricsMap["dayStartOffsetSeconds"].(float64); ok {
		DayStartOffsetSecondsInt := int(DayStartOffsetSeconds)
		o.DayStartOffsetSeconds = &DayStartOffsetSecondsInt
	}
	
	if AdherenceScheduleSeconds, ok := AgentadherencedaymetricsMap["adherenceScheduleSeconds"].(float64); ok {
		AdherenceScheduleSecondsInt := int(AdherenceScheduleSeconds)
		o.AdherenceScheduleSeconds = &AdherenceScheduleSecondsInt
	}
	
	if ConformanceScheduleSeconds, ok := AgentadherencedaymetricsMap["conformanceScheduleSeconds"].(float64); ok {
		ConformanceScheduleSecondsInt := int(ConformanceScheduleSeconds)
		o.ConformanceScheduleSeconds = &ConformanceScheduleSecondsInt
	}
	
	if ConformanceActualSeconds, ok := AgentadherencedaymetricsMap["conformanceActualSeconds"].(float64); ok {
		ConformanceActualSecondsInt := int(ConformanceActualSeconds)
		o.ConformanceActualSeconds = &ConformanceActualSecondsInt
	}
	
	if ExceptionCount, ok := AgentadherencedaymetricsMap["exceptionCount"].(float64); ok {
		ExceptionCountInt := int(ExceptionCount)
		o.ExceptionCount = &ExceptionCountInt
	}
	
	if ExceptionDurationSeconds, ok := AgentadherencedaymetricsMap["exceptionDurationSeconds"].(float64); ok {
		ExceptionDurationSecondsInt := int(ExceptionDurationSeconds)
		o.ExceptionDurationSeconds = &ExceptionDurationSecondsInt
	}
	
	if ImpactSeconds, ok := AgentadherencedaymetricsMap["impactSeconds"].(float64); ok {
		ImpactSecondsInt := int(ImpactSeconds)
		o.ImpactSeconds = &ImpactSecondsInt
	}
	
	if ScheduleLengthSeconds, ok := AgentadherencedaymetricsMap["scheduleLengthSeconds"].(float64); ok {
		ScheduleLengthSecondsInt := int(ScheduleLengthSeconds)
		o.ScheduleLengthSeconds = &ScheduleLengthSecondsInt
	}
	
	if ActualLengthSeconds, ok := AgentadherencedaymetricsMap["actualLengthSeconds"].(float64); ok {
		ActualLengthSecondsInt := int(ActualLengthSeconds)
		o.ActualLengthSeconds = &ActualLengthSecondsInt
	}
	
	if AdherencePercentage, ok := AgentadherencedaymetricsMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := AgentadherencedaymetricsMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentadherencedaymetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
