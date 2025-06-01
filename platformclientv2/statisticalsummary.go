package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Statisticalsummary
type Statisticalsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Max
	Max *float32 `json:"max,omitempty"`

	// Min
	Min *float32 `json:"min,omitempty"`

	// Count
	Count *int `json:"count,omitempty"`

	// CountNegative
	CountNegative *int `json:"countNegative,omitempty"`

	// CountPositive
	CountPositive *int `json:"countPositive,omitempty"`

	// Sum
	Sum *float32 `json:"sum,omitempty"`

	// Current
	Current *float32 `json:"current,omitempty"`

	// Ratio
	Ratio *float32 `json:"ratio,omitempty"`

	// Numerator
	Numerator *float32 `json:"numerator,omitempty"`

	// Denominator
	Denominator *float32 `json:"denominator,omitempty"`

	// Target
	Target *float32 `json:"target,omitempty"`

	// P95
	P95 *int `json:"p95,omitempty"`

	// P99
	P99 *int `json:"p99,omitempty"`

	// CalculatedMetricValue
	CalculatedMetricValue *int `json:"calculatedMetricValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Statisticalsummary) SetField(field string, fieldValue interface{}) {
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

func (o Statisticalsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Statisticalsummary
	
	return json.Marshal(&struct { 
		Max *float32 `json:"max,omitempty"`
		
		Min *float32 `json:"min,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		CountNegative *int `json:"countNegative,omitempty"`
		
		CountPositive *int `json:"countPositive,omitempty"`
		
		Sum *float32 `json:"sum,omitempty"`
		
		Current *float32 `json:"current,omitempty"`
		
		Ratio *float32 `json:"ratio,omitempty"`
		
		Numerator *float32 `json:"numerator,omitempty"`
		
		Denominator *float32 `json:"denominator,omitempty"`
		
		Target *float32 `json:"target,omitempty"`
		
		P95 *int `json:"p95,omitempty"`
		
		P99 *int `json:"p99,omitempty"`
		
		CalculatedMetricValue *int `json:"calculatedMetricValue,omitempty"`
		Alias
	}{ 
		Max: o.Max,
		
		Min: o.Min,
		
		Count: o.Count,
		
		CountNegative: o.CountNegative,
		
		CountPositive: o.CountPositive,
		
		Sum: o.Sum,
		
		Current: o.Current,
		
		Ratio: o.Ratio,
		
		Numerator: o.Numerator,
		
		Denominator: o.Denominator,
		
		Target: o.Target,
		
		P95: o.P95,
		
		P99: o.P99,
		
		CalculatedMetricValue: o.CalculatedMetricValue,
		Alias:    (Alias)(o),
	})
}

func (o *Statisticalsummary) UnmarshalJSON(b []byte) error {
	var StatisticalsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &StatisticalsummaryMap)
	if err != nil {
		return err
	}
	
	if Max, ok := StatisticalsummaryMap["max"].(float64); ok {
		MaxFloat32 := float32(Max)
		o.Max = &MaxFloat32
	}
    
	if Min, ok := StatisticalsummaryMap["min"].(float64); ok {
		MinFloat32 := float32(Min)
		o.Min = &MinFloat32
	}
    
	if Count, ok := StatisticalsummaryMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if CountNegative, ok := StatisticalsummaryMap["countNegative"].(float64); ok {
		CountNegativeInt := int(CountNegative)
		o.CountNegative = &CountNegativeInt
	}
	
	if CountPositive, ok := StatisticalsummaryMap["countPositive"].(float64); ok {
		CountPositiveInt := int(CountPositive)
		o.CountPositive = &CountPositiveInt
	}
	
	if Sum, ok := StatisticalsummaryMap["sum"].(float64); ok {
		SumFloat32 := float32(Sum)
		o.Sum = &SumFloat32
	}
    
	if Current, ok := StatisticalsummaryMap["current"].(float64); ok {
		CurrentFloat32 := float32(Current)
		o.Current = &CurrentFloat32
	}
    
	if Ratio, ok := StatisticalsummaryMap["ratio"].(float64); ok {
		RatioFloat32 := float32(Ratio)
		o.Ratio = &RatioFloat32
	}
    
	if Numerator, ok := StatisticalsummaryMap["numerator"].(float64); ok {
		NumeratorFloat32 := float32(Numerator)
		o.Numerator = &NumeratorFloat32
	}
    
	if Denominator, ok := StatisticalsummaryMap["denominator"].(float64); ok {
		DenominatorFloat32 := float32(Denominator)
		o.Denominator = &DenominatorFloat32
	}
    
	if Target, ok := StatisticalsummaryMap["target"].(float64); ok {
		TargetFloat32 := float32(Target)
		o.Target = &TargetFloat32
	}
    
	if P95, ok := StatisticalsummaryMap["p95"].(float64); ok {
		P95Int := int(P95)
		o.P95 = &P95Int
	}
	
	if P99, ok := StatisticalsummaryMap["p99"].(float64); ok {
		P99Int := int(P99)
		o.P99 = &P99Int
	}
	
	if CalculatedMetricValue, ok := StatisticalsummaryMap["calculatedMetricValue"].(float64); ok {
		CalculatedMetricValueInt := int(CalculatedMetricValue)
		o.CalculatedMetricValue = &CalculatedMetricValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Statisticalsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
