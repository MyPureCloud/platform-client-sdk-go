package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastplanninggroupdata
type Longtermforecastplanninggroupdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
	PlanningGroupId *string `json:"planningGroupId,omitempty"`

	// OfferedPerDay - Forecast offered counts per day for this planning group
	OfferedPerDay *[]float64 `json:"offeredPerDay,omitempty"`

	// AverageHandleTimeSecondsPerDay - Forecast average handle time per day in seconds
	AverageHandleTimeSecondsPerDay *[]float64 `json:"averageHandleTimeSecondsPerDay,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Longtermforecastplanninggroupdata) SetField(field string, fieldValue interface{}) {
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

func (o Longtermforecastplanninggroupdata) MarshalJSON() ([]byte, error) {
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
	type Alias Longtermforecastplanninggroupdata
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		OfferedPerDay *[]float64 `json:"offeredPerDay,omitempty"`
		
		AverageHandleTimeSecondsPerDay *[]float64 `json:"averageHandleTimeSecondsPerDay,omitempty"`
		Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		OfferedPerDay: o.OfferedPerDay,
		
		AverageHandleTimeSecondsPerDay: o.AverageHandleTimeSecondsPerDay,
		Alias:    (Alias)(o),
	})
}

func (o *Longtermforecastplanninggroupdata) UnmarshalJSON(b []byte) error {
	var LongtermforecastplanninggroupdataMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermforecastplanninggroupdataMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := LongtermforecastplanninggroupdataMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    
	if OfferedPerDay, ok := LongtermforecastplanninggroupdataMap["offeredPerDay"].([]interface{}); ok {
		OfferedPerDayString, _ := json.Marshal(OfferedPerDay)
		json.Unmarshal(OfferedPerDayString, &o.OfferedPerDay)
	}
	
	if AverageHandleTimeSecondsPerDay, ok := LongtermforecastplanninggroupdataMap["averageHandleTimeSecondsPerDay"].([]interface{}); ok {
		AverageHandleTimeSecondsPerDayString, _ := json.Marshal(AverageHandleTimeSecondsPerDay)
		json.Unmarshal(AverageHandleTimeSecondsPerDayString, &o.AverageHandleTimeSecondsPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermforecastplanninggroupdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
