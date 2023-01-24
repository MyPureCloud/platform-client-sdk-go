package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherencedaymetrics
type Historicaladherencedaymetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DayStartOffsetSecs - Start of day offset in seconds relative to query start time
	DayStartOffsetSecs *int `json:"dayStartOffsetSecs,omitempty"`

	// AdherenceScheduleSecs - Duration of schedule in seconds included for adherence percentage calculation
	AdherenceScheduleSecs *int `json:"adherenceScheduleSecs,omitempty"`

	// ConformanceScheduleSecs - Total scheduled duration in seconds for OnQueue activities
	ConformanceScheduleSecs *int `json:"conformanceScheduleSecs,omitempty"`

	// ConformanceActualSecs - Total actually worked duration in seconds for OnQueue activities
	ConformanceActualSecs *int `json:"conformanceActualSecs,omitempty"`

	// ExceptionCount - Total number of adherence exceptions for this user
	ExceptionCount *int `json:"exceptionCount,omitempty"`

	// ExceptionDurationSecs - Total duration in seconds of adherence exceptions for this user
	ExceptionDurationSecs *int `json:"exceptionDurationSecs,omitempty"`

	// ImpactSeconds - The impact duration in seconds of current adherence state for this user
	ImpactSeconds *int `json:"impactSeconds,omitempty"`

	// ScheduleLengthSecs - Total duration in seconds for all scheduled activities
	ScheduleLengthSecs *int `json:"scheduleLengthSecs,omitempty"`

	// ActualLengthSecs - Total duration in seconds for all actually worked activities
	ActualLengthSecs *int `json:"actualLengthSecs,omitempty"`

	// AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`

	// ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicaladherencedaymetrics) SetField(field string, fieldValue interface{}) {
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

func (o Historicaladherencedaymetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Historicaladherencedaymetrics
	
	return json.Marshal(&struct { 
		DayStartOffsetSecs *int `json:"dayStartOffsetSecs,omitempty"`
		
		AdherenceScheduleSecs *int `json:"adherenceScheduleSecs,omitempty"`
		
		ConformanceScheduleSecs *int `json:"conformanceScheduleSecs,omitempty"`
		
		ConformanceActualSecs *int `json:"conformanceActualSecs,omitempty"`
		
		ExceptionCount *int `json:"exceptionCount,omitempty"`
		
		ExceptionDurationSecs *int `json:"exceptionDurationSecs,omitempty"`
		
		ImpactSeconds *int `json:"impactSeconds,omitempty"`
		
		ScheduleLengthSecs *int `json:"scheduleLengthSecs,omitempty"`
		
		ActualLengthSecs *int `json:"actualLengthSecs,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		Alias
	}{ 
		DayStartOffsetSecs: o.DayStartOffsetSecs,
		
		AdherenceScheduleSecs: o.AdherenceScheduleSecs,
		
		ConformanceScheduleSecs: o.ConformanceScheduleSecs,
		
		ConformanceActualSecs: o.ConformanceActualSecs,
		
		ExceptionCount: o.ExceptionCount,
		
		ExceptionDurationSecs: o.ExceptionDurationSecs,
		
		ImpactSeconds: o.ImpactSeconds,
		
		ScheduleLengthSecs: o.ScheduleLengthSecs,
		
		ActualLengthSecs: o.ActualLengthSecs,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		Alias:    (Alias)(o),
	})
}

func (o *Historicaladherencedaymetrics) UnmarshalJSON(b []byte) error {
	var HistoricaladherencedaymetricsMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricaladherencedaymetricsMap)
	if err != nil {
		return err
	}
	
	if DayStartOffsetSecs, ok := HistoricaladherencedaymetricsMap["dayStartOffsetSecs"].(float64); ok {
		DayStartOffsetSecsInt := int(DayStartOffsetSecs)
		o.DayStartOffsetSecs = &DayStartOffsetSecsInt
	}
	
	if AdherenceScheduleSecs, ok := HistoricaladherencedaymetricsMap["adherenceScheduleSecs"].(float64); ok {
		AdherenceScheduleSecsInt := int(AdherenceScheduleSecs)
		o.AdherenceScheduleSecs = &AdherenceScheduleSecsInt
	}
	
	if ConformanceScheduleSecs, ok := HistoricaladherencedaymetricsMap["conformanceScheduleSecs"].(float64); ok {
		ConformanceScheduleSecsInt := int(ConformanceScheduleSecs)
		o.ConformanceScheduleSecs = &ConformanceScheduleSecsInt
	}
	
	if ConformanceActualSecs, ok := HistoricaladherencedaymetricsMap["conformanceActualSecs"].(float64); ok {
		ConformanceActualSecsInt := int(ConformanceActualSecs)
		o.ConformanceActualSecs = &ConformanceActualSecsInt
	}
	
	if ExceptionCount, ok := HistoricaladherencedaymetricsMap["exceptionCount"].(float64); ok {
		ExceptionCountInt := int(ExceptionCount)
		o.ExceptionCount = &ExceptionCountInt
	}
	
	if ExceptionDurationSecs, ok := HistoricaladherencedaymetricsMap["exceptionDurationSecs"].(float64); ok {
		ExceptionDurationSecsInt := int(ExceptionDurationSecs)
		o.ExceptionDurationSecs = &ExceptionDurationSecsInt
	}
	
	if ImpactSeconds, ok := HistoricaladherencedaymetricsMap["impactSeconds"].(float64); ok {
		ImpactSecondsInt := int(ImpactSeconds)
		o.ImpactSeconds = &ImpactSecondsInt
	}
	
	if ScheduleLengthSecs, ok := HistoricaladherencedaymetricsMap["scheduleLengthSecs"].(float64); ok {
		ScheduleLengthSecsInt := int(ScheduleLengthSecs)
		o.ScheduleLengthSecs = &ScheduleLengthSecsInt
	}
	
	if ActualLengthSecs, ok := HistoricaladherencedaymetricsMap["actualLengthSecs"].(float64); ok {
		ActualLengthSecsInt := int(ActualLengthSecs)
		o.ActualLengthSecs = &ActualLengthSecsInt
	}
	
	if AdherencePercentage, ok := HistoricaladherencedaymetricsMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := HistoricaladherencedaymetricsMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Historicaladherencedaymetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
