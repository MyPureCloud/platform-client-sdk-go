package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradaydatagroup
type Wfmintradaydataupdatetopicintradaydatagroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// ForecastDataPerInterval
	ForecastDataPerInterval *[]Wfmintradaydataupdatetopicintradayforecastdata `json:"forecastDataPerInterval,omitempty"`

	// ScheduleDataPerInterval
	ScheduleDataPerInterval *[]Wfmintradaydataupdatetopicintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`

	// HistoricalAgentDataPerInterval
	HistoricalAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalagentdata `json:"historicalAgentDataPerInterval,omitempty"`

	// HistoricalQueueDataPerInterval
	HistoricalQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalqueuedata `json:"historicalQueueDataPerInterval,omitempty"`

	// PerformancePredictionAgentDataPerInterval
	PerformancePredictionAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionagentdata `json:"performancePredictionAgentDataPerInterval,omitempty"`

	// PerformancePredictionQueueDataPerInterval
	PerformancePredictionQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata `json:"performancePredictionQueueDataPerInterval,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmintradaydataupdatetopicintradaydatagroup) SetField(field string, fieldValue interface{}) {
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

func (o Wfmintradaydataupdatetopicintradaydatagroup) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmintradaydataupdatetopicintradaydatagroup
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ForecastDataPerInterval *[]Wfmintradaydataupdatetopicintradayforecastdata `json:"forecastDataPerInterval,omitempty"`
		
		ScheduleDataPerInterval *[]Wfmintradaydataupdatetopicintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`
		
		HistoricalAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalagentdata `json:"historicalAgentDataPerInterval,omitempty"`
		
		HistoricalQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalqueuedata `json:"historicalQueueDataPerInterval,omitempty"`
		
		PerformancePredictionAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionagentdata `json:"performancePredictionAgentDataPerInterval,omitempty"`
		
		PerformancePredictionQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata `json:"performancePredictionQueueDataPerInterval,omitempty"`
		Alias
	}{ 
		MediaType: o.MediaType,
		
		ForecastDataPerInterval: o.ForecastDataPerInterval,
		
		ScheduleDataPerInterval: o.ScheduleDataPerInterval,
		
		HistoricalAgentDataPerInterval: o.HistoricalAgentDataPerInterval,
		
		HistoricalQueueDataPerInterval: o.HistoricalQueueDataPerInterval,
		
		PerformancePredictionAgentDataPerInterval: o.PerformancePredictionAgentDataPerInterval,
		
		PerformancePredictionQueueDataPerInterval: o.PerformancePredictionQueueDataPerInterval,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradaydatagroup) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradaydatagroupMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradaydatagroupMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := WfmintradaydataupdatetopicintradaydatagroupMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ForecastDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["forecastDataPerInterval"].([]interface{}); ok {
		ForecastDataPerIntervalString, _ := json.Marshal(ForecastDataPerInterval)
		json.Unmarshal(ForecastDataPerIntervalString, &o.ForecastDataPerInterval)
	}
	
	if ScheduleDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["scheduleDataPerInterval"].([]interface{}); ok {
		ScheduleDataPerIntervalString, _ := json.Marshal(ScheduleDataPerInterval)
		json.Unmarshal(ScheduleDataPerIntervalString, &o.ScheduleDataPerInterval)
	}
	
	if HistoricalAgentDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["historicalAgentDataPerInterval"].([]interface{}); ok {
		HistoricalAgentDataPerIntervalString, _ := json.Marshal(HistoricalAgentDataPerInterval)
		json.Unmarshal(HistoricalAgentDataPerIntervalString, &o.HistoricalAgentDataPerInterval)
	}
	
	if HistoricalQueueDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["historicalQueueDataPerInterval"].([]interface{}); ok {
		HistoricalQueueDataPerIntervalString, _ := json.Marshal(HistoricalQueueDataPerInterval)
		json.Unmarshal(HistoricalQueueDataPerIntervalString, &o.HistoricalQueueDataPerInterval)
	}
	
	if PerformancePredictionAgentDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["performancePredictionAgentDataPerInterval"].([]interface{}); ok {
		PerformancePredictionAgentDataPerIntervalString, _ := json.Marshal(PerformancePredictionAgentDataPerInterval)
		json.Unmarshal(PerformancePredictionAgentDataPerIntervalString, &o.PerformancePredictionAgentDataPerInterval)
	}
	
	if PerformancePredictionQueueDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["performancePredictionQueueDataPerInterval"].([]interface{}); ok {
		PerformancePredictionQueueDataPerIntervalString, _ := json.Marshal(PerformancePredictionQueueDataPerInterval)
		json.Unmarshal(PerformancePredictionQueueDataPerIntervalString, &o.PerformancePredictionQueueDataPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
