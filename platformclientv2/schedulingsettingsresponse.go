package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingsettingsresponse
type Schedulingsettingsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MaxOccupancyPercentForDeferredWork - Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`

	// DefaultShrinkagePercent - Default shrinkage percent for scheduling
	DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`

	// ShrinkageOverrides - Shrinkage overrides for scheduling
	ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`

	// PlanningPeriod - Planning period settings for scheduling
	PlanningPeriod *Planningperiodsettings `json:"planningPeriod,omitempty"`

	// StartDayOfWeekend - Start day of weekend for scheduling
	StartDayOfWeekend *string `json:"startDayOfWeekend,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulingsettingsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Schedulingsettingsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulingsettingsresponse
	
	return json.Marshal(&struct { 
		MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`
		
		DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`
		
		ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`
		
		PlanningPeriod *Planningperiodsettings `json:"planningPeriod,omitempty"`
		
		StartDayOfWeekend *string `json:"startDayOfWeekend,omitempty"`
		Alias
	}{ 
		MaxOccupancyPercentForDeferredWork: o.MaxOccupancyPercentForDeferredWork,
		
		DefaultShrinkagePercent: o.DefaultShrinkagePercent,
		
		ShrinkageOverrides: o.ShrinkageOverrides,
		
		PlanningPeriod: o.PlanningPeriod,
		
		StartDayOfWeekend: o.StartDayOfWeekend,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulingsettingsresponse) UnmarshalJSON(b []byte) error {
	var SchedulingsettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingsettingsresponseMap)
	if err != nil {
		return err
	}
	
	if MaxOccupancyPercentForDeferredWork, ok := SchedulingsettingsresponseMap["maxOccupancyPercentForDeferredWork"].(float64); ok {
		MaxOccupancyPercentForDeferredWorkInt := int(MaxOccupancyPercentForDeferredWork)
		o.MaxOccupancyPercentForDeferredWork = &MaxOccupancyPercentForDeferredWorkInt
	}
	
	if DefaultShrinkagePercent, ok := SchedulingsettingsresponseMap["defaultShrinkagePercent"].(float64); ok {
		o.DefaultShrinkagePercent = &DefaultShrinkagePercent
	}
    
	if ShrinkageOverrides, ok := SchedulingsettingsresponseMap["shrinkageOverrides"].(map[string]interface{}); ok {
		ShrinkageOverridesString, _ := json.Marshal(ShrinkageOverrides)
		json.Unmarshal(ShrinkageOverridesString, &o.ShrinkageOverrides)
	}
	
	if PlanningPeriod, ok := SchedulingsettingsresponseMap["planningPeriod"].(map[string]interface{}); ok {
		PlanningPeriodString, _ := json.Marshal(PlanningPeriod)
		json.Unmarshal(PlanningPeriodString, &o.PlanningPeriod)
	}
	
	if StartDayOfWeekend, ok := SchedulingsettingsresponseMap["startDayOfWeekend"].(string); ok {
		o.StartDayOfWeekend = &StartDayOfWeekend
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
