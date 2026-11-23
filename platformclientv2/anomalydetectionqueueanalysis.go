package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Anomalydetectionqueueanalysis
type Anomalydetectionqueueanalysis struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metric - The name of the metric for which analysis was performed
	Metric *string `json:"metric,omitempty"`

	// MediaType - The media type that was analyzed
	MediaType *string `json:"mediaType,omitempty"`

	// DateLastUpdated - The timestamp when the model was last updated, in ISO-8601 format
	DateLastUpdated *time.Time `json:"dateLastUpdated,omitempty"`

	// AnalysisCode - When specified, indicates issues with the anomaly detection analysis
	AnalysisCode *string `json:"analysisCode,omitempty"`

	// TimeIntervalData - The time interval data with forecasted bounds. Not populated when analysisCode indicates an issue with the analysis
	TimeIntervalData *[]Anomalydetectiontimeintervaldata `json:"timeIntervalData,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Anomalydetectionqueueanalysis) SetField(field string, fieldValue interface{}) {
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

func (o Anomalydetectionqueueanalysis) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastUpdated", }
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
	type Alias Anomalydetectionqueueanalysis
	
	DateLastUpdated := new(string)
	if o.DateLastUpdated != nil {
		
		*DateLastUpdated = timeutil.Strftime(o.DateLastUpdated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastUpdated = nil
	}
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		DateLastUpdated *string `json:"dateLastUpdated,omitempty"`
		
		AnalysisCode *string `json:"analysisCode,omitempty"`
		
		TimeIntervalData *[]Anomalydetectiontimeintervaldata `json:"timeIntervalData,omitempty"`
		Alias
	}{ 
		Metric: o.Metric,
		
		MediaType: o.MediaType,
		
		DateLastUpdated: DateLastUpdated,
		
		AnalysisCode: o.AnalysisCode,
		
		TimeIntervalData: o.TimeIntervalData,
		Alias:    (Alias)(o),
	})
}

func (o *Anomalydetectionqueueanalysis) UnmarshalJSON(b []byte) error {
	var AnomalydetectionqueueanalysisMap map[string]interface{}
	err := json.Unmarshal(b, &AnomalydetectionqueueanalysisMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := AnomalydetectionqueueanalysisMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MediaType, ok := AnomalydetectionqueueanalysisMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if dateLastUpdatedString, ok := AnomalydetectionqueueanalysisMap["dateLastUpdated"].(string); ok {
		DateLastUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastUpdatedString)
		o.DateLastUpdated = &DateLastUpdated
	}
	
	if AnalysisCode, ok := AnomalydetectionqueueanalysisMap["analysisCode"].(string); ok {
		o.AnalysisCode = &AnalysisCode
	}
    
	if TimeIntervalData, ok := AnomalydetectionqueueanalysisMap["timeIntervalData"].([]interface{}); ok {
		TimeIntervalDataString, _ := json.Marshal(TimeIntervalData)
		json.Unmarshal(TimeIntervalDataString, &o.TimeIntervalData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Anomalydetectionqueueanalysis) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
