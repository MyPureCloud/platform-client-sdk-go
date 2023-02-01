package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluesmetricitem
type Workdayvaluesmetricitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metric - Gamification metric for the average and the trend
	Metric *Addressableentityref `json:"metric,omitempty"`

	// MetricDefinition - Gamification metric definition for the average and the trend
	MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`

	// Average - The average value of the metric
	Average *float64 `json:"average,omitempty"`

	// UnitType - The unit type of the metric value
	UnitType *string `json:"unitType,omitempty"`

	// Trend - The metric value trend
	Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workdayvaluesmetricitem) SetField(field string, fieldValue interface{}) {
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

func (o Workdayvaluesmetricitem) MarshalJSON() ([]byte, error) {
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
	type Alias Workdayvaluesmetricitem
	
	return json.Marshal(&struct { 
		Metric *Addressableentityref `json:"metric,omitempty"`
		
		MetricDefinition *Domainentityref `json:"metricDefinition,omitempty"`
		
		Average *float64 `json:"average,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		Trend *[]Workdayvaluestrenditem `json:"trend,omitempty"`
		Alias
	}{ 
		Metric: o.Metric,
		
		MetricDefinition: o.MetricDefinition,
		
		Average: o.Average,
		
		UnitType: o.UnitType,
		
		Trend: o.Trend,
		Alias:    (Alias)(o),
	})
}

func (o *Workdayvaluesmetricitem) UnmarshalJSON(b []byte) error {
	var WorkdayvaluesmetricitemMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdayvaluesmetricitemMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := WorkdayvaluesmetricitemMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	
	if MetricDefinition, ok := WorkdayvaluesmetricitemMap["metricDefinition"].(map[string]interface{}); ok {
		MetricDefinitionString, _ := json.Marshal(MetricDefinition)
		json.Unmarshal(MetricDefinitionString, &o.MetricDefinition)
	}
	
	if Average, ok := WorkdayvaluesmetricitemMap["average"].(float64); ok {
		o.Average = &Average
	}
    
	if UnitType, ok := WorkdayvaluesmetricitemMap["unitType"].(string); ok {
		o.UnitType = &UnitType
	}
    
	if Trend, ok := WorkdayvaluesmetricitemMap["trend"].([]interface{}); ok {
		TrendString, _ := json.Marshal(Trend)
		json.Unmarshal(TrendString, &o.Trend)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdayvaluesmetricitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
