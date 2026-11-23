package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Anomalysummary
type Anomalysummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateIntervalStart - The start of the time interval for the anomaly summary count, in ISO-8601 format
	DateIntervalStart *time.Time `json:"dateIntervalStart,omitempty"`

	// DateIntervalEnd - The end of the time interval for the anomaly summary count, in ISO-8601 format
	DateIntervalEnd *time.Time `json:"dateIntervalEnd,omitempty"`

	// AnomalyCount - The number of anomalies detected in the time window
	AnomalyCount *int `json:"anomalyCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Anomalysummary) SetField(field string, fieldValue interface{}) {
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

func (o Anomalysummary) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateIntervalStart","DateIntervalEnd", }
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
	type Alias Anomalysummary
	
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
	
	return json.Marshal(&struct { 
		DateIntervalStart *string `json:"dateIntervalStart,omitempty"`
		
		DateIntervalEnd *string `json:"dateIntervalEnd,omitempty"`
		
		AnomalyCount *int `json:"anomalyCount,omitempty"`
		Alias
	}{ 
		DateIntervalStart: DateIntervalStart,
		
		DateIntervalEnd: DateIntervalEnd,
		
		AnomalyCount: o.AnomalyCount,
		Alias:    (Alias)(o),
	})
}

func (o *Anomalysummary) UnmarshalJSON(b []byte) error {
	var AnomalysummaryMap map[string]interface{}
	err := json.Unmarshal(b, &AnomalysummaryMap)
	if err != nil {
		return err
	}
	
	if dateIntervalStartString, ok := AnomalysummaryMap["dateIntervalStart"].(string); ok {
		DateIntervalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalStartString)
		o.DateIntervalStart = &DateIntervalStart
	}
	
	if dateIntervalEndString, ok := AnomalysummaryMap["dateIntervalEnd"].(string); ok {
		DateIntervalEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalEndString)
		o.DateIntervalEnd = &DateIntervalEnd
	}
	
	if AnomalyCount, ok := AnomalysummaryMap["anomalyCount"].(float64); ok {
		AnomalyCountInt := int(AnomalyCount)
		o.AnomalyCount = &AnomalyCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Anomalysummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
