package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastplanninggroupdata
type Forecastplanninggroupdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
	PlanningGroupId *string `json:"planningGroupId,omitempty"`

	// OfferedPerInterval - Forecast offered counts per 15 minute interval for this week of the forecast
	OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`

	// AverageHandleTimeSecondsPerInterval - Forecast average handle time per 15 minute interval in seconds
	AverageHandleTimeSecondsPerInterval *[]float64 `json:"averageHandleTimeSecondsPerInterval,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Forecastplanninggroupdata) SetField(field string, fieldValue interface{}) {
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

func (o Forecastplanninggroupdata) MarshalJSON() ([]byte, error) {
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
	type Alias Forecastplanninggroupdata
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		OfferedPerInterval *[]float64 `json:"offeredPerInterval,omitempty"`
		
		AverageHandleTimeSecondsPerInterval *[]float64 `json:"averageHandleTimeSecondsPerInterval,omitempty"`
		Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		OfferedPerInterval: o.OfferedPerInterval,
		
		AverageHandleTimeSecondsPerInterval: o.AverageHandleTimeSecondsPerInterval,
		Alias:    (Alias)(o),
	})
}

func (o *Forecastplanninggroupdata) UnmarshalJSON(b []byte) error {
	var ForecastplanninggroupdataMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastplanninggroupdataMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := ForecastplanninggroupdataMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    
	if OfferedPerInterval, ok := ForecastplanninggroupdataMap["offeredPerInterval"].([]interface{}); ok {
		OfferedPerIntervalString, _ := json.Marshal(OfferedPerInterval)
		json.Unmarshal(OfferedPerIntervalString, &o.OfferedPerInterval)
	}
	
	if AverageHandleTimeSecondsPerInterval, ok := ForecastplanninggroupdataMap["averageHandleTimeSecondsPerInterval"].([]interface{}); ok {
		AverageHandleTimeSecondsPerIntervalString, _ := json.Marshal(AverageHandleTimeSecondsPerInterval)
		json.Unmarshal(AverageHandleTimeSecondsPerIntervalString, &o.AverageHandleTimeSecondsPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
