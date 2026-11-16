package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2queueanomalyeventqueueanomalynotification
type V2queueanomalyeventqueueanomalynotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Queue
	Queue *V2queueanomalyeventqueuereference `json:"queue,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Division
	Division *V2queueanomalyeventdivisionreference `json:"division,omitempty"`

	// FalsePositive
	FalsePositive *bool `json:"falsePositive,omitempty"`

	// Metric
	Metric *string `json:"metric,omitempty"`

	// DateAnomalyDetected
	DateAnomalyDetected *time.Time `json:"dateAnomalyDetected,omitempty"`

	// DateIntervalStart
	DateIntervalStart *time.Time `json:"dateIntervalStart,omitempty"`

	// DateIntervalEnd
	DateIntervalEnd *time.Time `json:"dateIntervalEnd,omitempty"`

	// ConsensusAnalyses
	ConsensusAnalyses *[]V2queueanomalyeventconsensusanalysis `json:"consensusAnalyses,omitempty"`

	// AnomalyScorePercent
	AnomalyScorePercent *float32 `json:"anomalyScorePercent,omitempty"`

	// DetectedValue
	DetectedValue *float32 `json:"detectedValue,omitempty"`

	// LowerBound
	LowerBound *float32 `json:"lowerBound,omitempty"`

	// UpperBound
	UpperBound *float32 `json:"upperBound,omitempty"`

	// ExpectedValue
	ExpectedValue *float32 `json:"expectedValue,omitempty"`

	// Deviation
	Deviation *float32 `json:"deviation,omitempty"`

	// DateExpires
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2queueanomalyeventqueueanomalynotification) SetField(field string, fieldValue interface{}) {
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

func (o V2queueanomalyeventqueueanomalynotification) MarshalJSON() ([]byte, error) {
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
	type Alias V2queueanomalyeventqueueanomalynotification
	
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
		
		Queue *V2queueanomalyeventqueuereference `json:"queue,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Division *V2queueanomalyeventdivisionreference `json:"division,omitempty"`
		
		FalsePositive *bool `json:"falsePositive,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		DateAnomalyDetected *string `json:"dateAnomalyDetected,omitempty"`
		
		DateIntervalStart *string `json:"dateIntervalStart,omitempty"`
		
		DateIntervalEnd *string `json:"dateIntervalEnd,omitempty"`
		
		ConsensusAnalyses *[]V2queueanomalyeventconsensusanalysis `json:"consensusAnalyses,omitempty"`
		
		AnomalyScorePercent *float32 `json:"anomalyScorePercent,omitempty"`
		
		DetectedValue *float32 `json:"detectedValue,omitempty"`
		
		LowerBound *float32 `json:"lowerBound,omitempty"`
		
		UpperBound *float32 `json:"upperBound,omitempty"`
		
		ExpectedValue *float32 `json:"expectedValue,omitempty"`
		
		Deviation *float32 `json:"deviation,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		Division: o.Division,
		
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

func (o *V2queueanomalyeventqueueanomalynotification) UnmarshalJSON(b []byte) error {
	var V2queueanomalyeventqueueanomalynotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2queueanomalyeventqueueanomalynotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2queueanomalyeventqueueanomalynotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Queue, ok := V2queueanomalyeventqueueanomalynotificationMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := V2queueanomalyeventqueueanomalynotificationMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Division, ok := V2queueanomalyeventqueueanomalynotificationMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if FalsePositive, ok := V2queueanomalyeventqueueanomalynotificationMap["falsePositive"].(bool); ok {
		o.FalsePositive = &FalsePositive
	}
    
	if Metric, ok := V2queueanomalyeventqueueanomalynotificationMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if dateAnomalyDetectedString, ok := V2queueanomalyeventqueueanomalynotificationMap["dateAnomalyDetected"].(string); ok {
		DateAnomalyDetected, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAnomalyDetectedString)
		o.DateAnomalyDetected = &DateAnomalyDetected
	}
	
	if dateIntervalStartString, ok := V2queueanomalyeventqueueanomalynotificationMap["dateIntervalStart"].(string); ok {
		DateIntervalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalStartString)
		o.DateIntervalStart = &DateIntervalStart
	}
	
	if dateIntervalEndString, ok := V2queueanomalyeventqueueanomalynotificationMap["dateIntervalEnd"].(string); ok {
		DateIntervalEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalEndString)
		o.DateIntervalEnd = &DateIntervalEnd
	}
	
	if ConsensusAnalyses, ok := V2queueanomalyeventqueueanomalynotificationMap["consensusAnalyses"].([]interface{}); ok {
		ConsensusAnalysesString, _ := json.Marshal(ConsensusAnalyses)
		json.Unmarshal(ConsensusAnalysesString, &o.ConsensusAnalyses)
	}
	
	if AnomalyScorePercent, ok := V2queueanomalyeventqueueanomalynotificationMap["anomalyScorePercent"].(float64); ok {
		AnomalyScorePercentFloat32 := float32(AnomalyScorePercent)
		o.AnomalyScorePercent = &AnomalyScorePercentFloat32
	}
    
	if DetectedValue, ok := V2queueanomalyeventqueueanomalynotificationMap["detectedValue"].(float64); ok {
		DetectedValueFloat32 := float32(DetectedValue)
		o.DetectedValue = &DetectedValueFloat32
	}
    
	if LowerBound, ok := V2queueanomalyeventqueueanomalynotificationMap["lowerBound"].(float64); ok {
		LowerBoundFloat32 := float32(LowerBound)
		o.LowerBound = &LowerBoundFloat32
	}
    
	if UpperBound, ok := V2queueanomalyeventqueueanomalynotificationMap["upperBound"].(float64); ok {
		UpperBoundFloat32 := float32(UpperBound)
		o.UpperBound = &UpperBoundFloat32
	}
    
	if ExpectedValue, ok := V2queueanomalyeventqueueanomalynotificationMap["expectedValue"].(float64); ok {
		ExpectedValueFloat32 := float32(ExpectedValue)
		o.ExpectedValue = &ExpectedValueFloat32
	}
    
	if Deviation, ok := V2queueanomalyeventqueueanomalynotificationMap["deviation"].(float64); ok {
		DeviationFloat32 := float32(Deviation)
		o.Deviation = &DeviationFloat32
	}
    
	if dateExpiresString, ok := V2queueanomalyeventqueueanomalynotificationMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if dateModifiedString, ok := V2queueanomalyeventqueueanomalynotificationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2queueanomalyeventqueueanomalynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
