package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Performancepredictionoutputs
type Performancepredictionoutputs struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CalculationStartDate - Date as an ISO-8601 string, corresponding to the beginning of the performance prediction results
	CalculationStartDate *time.Time `json:"calculationStartDate,omitempty"`

	// CalculationIntervalLengthMinutes - Interval length of the response metrics
	CalculationIntervalLengthMinutes *int `json:"calculationIntervalLengthMinutes,omitempty"`

	// PlanningGroupResults - List of planning group level performance prediction results
	PlanningGroupResults *[]Planninggroupoutputs `json:"planningGroupResults,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Performancepredictionoutputs) SetField(field string, fieldValue interface{}) {
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

func (o Performancepredictionoutputs) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CalculationStartDate", }
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
	type Alias Performancepredictionoutputs
	
	CalculationStartDate := new(string)
	if o.CalculationStartDate != nil {
		
		*CalculationStartDate = timeutil.Strftime(o.CalculationStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CalculationStartDate = nil
	}
	
	return json.Marshal(&struct { 
		CalculationStartDate *string `json:"calculationStartDate,omitempty"`
		
		CalculationIntervalLengthMinutes *int `json:"calculationIntervalLengthMinutes,omitempty"`
		
		PlanningGroupResults *[]Planninggroupoutputs `json:"planningGroupResults,omitempty"`
		Alias
	}{ 
		CalculationStartDate: CalculationStartDate,
		
		CalculationIntervalLengthMinutes: o.CalculationIntervalLengthMinutes,
		
		PlanningGroupResults: o.PlanningGroupResults,
		Alias:    (Alias)(o),
	})
}

func (o *Performancepredictionoutputs) UnmarshalJSON(b []byte) error {
	var PerformancepredictionoutputsMap map[string]interface{}
	err := json.Unmarshal(b, &PerformancepredictionoutputsMap)
	if err != nil {
		return err
	}
	
	if calculationStartDateString, ok := PerformancepredictionoutputsMap["calculationStartDate"].(string); ok {
		CalculationStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", calculationStartDateString)
		o.CalculationStartDate = &CalculationStartDate
	}
	
	if CalculationIntervalLengthMinutes, ok := PerformancepredictionoutputsMap["calculationIntervalLengthMinutes"].(float64); ok {
		CalculationIntervalLengthMinutesInt := int(CalculationIntervalLengthMinutes)
		o.CalculationIntervalLengthMinutes = &CalculationIntervalLengthMinutesInt
	}
	
	if PlanningGroupResults, ok := PerformancepredictionoutputsMap["planningGroupResults"].([]interface{}); ok {
		PlanningGroupResultsString, _ := json.Marshal(PlanningGroupResults)
		json.Unmarshal(PlanningGroupResultsString, &o.PlanningGroupResults)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Performancepredictionoutputs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
