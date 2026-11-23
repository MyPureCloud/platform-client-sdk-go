package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueanomalyexport
type Queueanomalyexport struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the queue anomaly
	Id *string `json:"id,omitempty"`

	// QueueId - The ID of the queue for which the anomaly was detected
	QueueId *string `json:"queueId,omitempty"`

	// MediaType - The media type associated with the detected anomaly
	MediaType *string `json:"mediaType,omitempty"`

	// DivisionId - The ID of the division in which the anomaly was detected
	DivisionId *string `json:"divisionId,omitempty"`

	// FalsePositive - Indicates whether the anomaly has been flagged as a false positive
	FalsePositive *bool `json:"falsePositive,omitempty"`

	// Metric - The name of the metric for which the anomaly was detected
	Metric *string `json:"metric,omitempty"`

	// DateAnomalyDetected - The date and time at which the anomaly was detected, in ISO-8601 format
	DateAnomalyDetected *time.Time `json:"dateAnomalyDetected,omitempty"`

	// DateIntervalStart - The beginning of the measured time period during which the anomaly was detected, in ISO-8601 format
	DateIntervalStart *time.Time `json:"dateIntervalStart,omitempty"`

	// DateIntervalEnd - The end of the measured time period during which the anomaly was detected, in ISO-8601 format
	DateIntervalEnd *time.Time `json:"dateIntervalEnd,omitempty"`

	// ConsensusAnalyses - The consensus methods used for analysis and their metadata
	ConsensusAnalyses *[]Anomalyconsensusanalysis `json:"consensusAnalyses,omitempty"`

	// AnomalyScorePercent - Percentage value (0-100) representing confidence that the detected value is anomalous
	AnomalyScorePercent *float64 `json:"anomalyScorePercent,omitempty"`

	// DetectedValue - The historical value that was analyzed
	DetectedValue *float64 `json:"detectedValue,omitempty"`

	// LowerBound - The lower bound of the normal range
	LowerBound *float64 `json:"lowerBound,omitempty"`

	// UpperBound - The upper bound of the normal range
	UpperBound *float64 `json:"upperBound,omitempty"`

	// ExpectedValue - The expected value based on historical data
	ExpectedValue *float64 `json:"expectedValue,omitempty"`

	// Deviation - Percentage deviation from the expected value
	Deviation *float64 `json:"deviation,omitempty"`

	// DateExpires - The date and time when this anomaly expires, in ISO-8601 format
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// DateModified - The date and time at which this anomaly was last modified, in ISO-8601 format
	DateModified *time.Time `json:"dateModified,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueanomalyexport) SetField(field string, fieldValue interface{}) {
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

func (o Queueanomalyexport) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAnomalyDetected","DateIntervalStart","DateIntervalEnd","DateExpires","DateModified", }
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
	type Alias Queueanomalyexport
	
	DateAnomalyDetected := new(string)
	if o.DateAnomalyDetected != nil {
		
		*DateAnomalyDetected = timeutil.Strftime(o.DateAnomalyDetected, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAnomalyDetected = nil
	}
	
	DateIntervalStart := new(string)
	if o.DateIntervalStart != nil {
		
		*DateIntervalStart = timeutil.Strftime(o.DateIntervalStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIntervalStart = nil
	}
	
	DateIntervalEnd := new(string)
	if o.DateIntervalEnd != nil {
		
		*DateIntervalEnd = timeutil.Strftime(o.DateIntervalEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIntervalEnd = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		FalsePositive *bool `json:"falsePositive,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		DateAnomalyDetected *string `json:"dateAnomalyDetected,omitempty"`
		
		DateIntervalStart *string `json:"dateIntervalStart,omitempty"`
		
		DateIntervalEnd *string `json:"dateIntervalEnd,omitempty"`
		
		ConsensusAnalyses *[]Anomalyconsensusanalysis `json:"consensusAnalyses,omitempty"`
		
		AnomalyScorePercent *float64 `json:"anomalyScorePercent,omitempty"`
		
		DetectedValue *float64 `json:"detectedValue,omitempty"`
		
		LowerBound *float64 `json:"lowerBound,omitempty"`
		
		UpperBound *float64 `json:"upperBound,omitempty"`
		
		ExpectedValue *float64 `json:"expectedValue,omitempty"`
		
		Deviation *float64 `json:"deviation,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		QueueId: o.QueueId,
		
		MediaType: o.MediaType,
		
		DivisionId: o.DivisionId,
		
		FalsePositive: o.FalsePositive,
		
		Metric: o.Metric,
		
		DateAnomalyDetected: DateAnomalyDetected,
		
		DateIntervalStart: DateIntervalStart,
		
		DateIntervalEnd: DateIntervalEnd,
		
		ConsensusAnalyses: o.ConsensusAnalyses,
		
		AnomalyScorePercent: o.AnomalyScorePercent,
		
		DetectedValue: o.DetectedValue,
		
		LowerBound: o.LowerBound,
		
		UpperBound: o.UpperBound,
		
		ExpectedValue: o.ExpectedValue,
		
		Deviation: o.Deviation,
		
		DateExpires: DateExpires,
		
		DateModified: DateModified,
		Alias:    (Alias)(o),
	})
}

func (o *Queueanomalyexport) UnmarshalJSON(b []byte) error {
	var QueueanomalyexportMap map[string]interface{}
	err := json.Unmarshal(b, &QueueanomalyexportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueanomalyexportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if QueueId, ok := QueueanomalyexportMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if MediaType, ok := QueueanomalyexportMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if DivisionId, ok := QueueanomalyexportMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if FalsePositive, ok := QueueanomalyexportMap["falsePositive"].(bool); ok {
		o.FalsePositive = &FalsePositive
	}
    
	if Metric, ok := QueueanomalyexportMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if dateAnomalyDetectedString, ok := QueueanomalyexportMap["dateAnomalyDetected"].(string); ok {
		DateAnomalyDetected, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAnomalyDetectedString)
		o.DateAnomalyDetected = &DateAnomalyDetected
	}
	
	if dateIntervalStartString, ok := QueueanomalyexportMap["dateIntervalStart"].(string); ok {
		DateIntervalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalStartString)
		o.DateIntervalStart = &DateIntervalStart
	}
	
	if dateIntervalEndString, ok := QueueanomalyexportMap["dateIntervalEnd"].(string); ok {
		DateIntervalEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalEndString)
		o.DateIntervalEnd = &DateIntervalEnd
	}
	
	if ConsensusAnalyses, ok := QueueanomalyexportMap["consensusAnalyses"].([]interface{}); ok {
		ConsensusAnalysesString, _ := json.Marshal(ConsensusAnalyses)
		json.Unmarshal(ConsensusAnalysesString, &o.ConsensusAnalyses)
	}
	
	if AnomalyScorePercent, ok := QueueanomalyexportMap["anomalyScorePercent"].(float64); ok {
		o.AnomalyScorePercent = &AnomalyScorePercent
	}
    
	if DetectedValue, ok := QueueanomalyexportMap["detectedValue"].(float64); ok {
		o.DetectedValue = &DetectedValue
	}
    
	if LowerBound, ok := QueueanomalyexportMap["lowerBound"].(float64); ok {
		o.LowerBound = &LowerBound
	}
    
	if UpperBound, ok := QueueanomalyexportMap["upperBound"].(float64); ok {
		o.UpperBound = &UpperBound
	}
    
	if ExpectedValue, ok := QueueanomalyexportMap["expectedValue"].(float64); ok {
		o.ExpectedValue = &ExpectedValue
	}
    
	if Deviation, ok := QueueanomalyexportMap["deviation"].(float64); ok {
		o.Deviation = &Deviation
	}
    
	if dateExpiresString, ok := QueueanomalyexportMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if dateModifiedString, ok := QueueanomalyexportMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueanomalyexport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
