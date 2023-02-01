package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buimportshorttermforecastschema
type Buimportshorttermforecastschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`

	// WeekCount - The number of weeks covered by the forecast
	WeekCount *int `json:"weekCount,omitempty"`

	// PlanningGroups - The short term planning group data
	PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`

	// LongTermPlanningGroups - The long term planning group data
	LongTermPlanningGroups *[]Longtermforecastplanninggroupdata `json:"longTermPlanningGroups,omitempty"`

	// CanUseForScheduling - Whether this forecast can be used for scheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buimportshorttermforecastschema) SetField(field string, fieldValue interface{}) {
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

func (o Buimportshorttermforecastschema) MarshalJSON() ([]byte, error) {
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
	type Alias Buimportshorttermforecastschema
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`
		
		LongTermPlanningGroups *[]Longtermforecastplanninggroupdata `json:"longTermPlanningGroups,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		Alias
	}{ 
		Description: o.Description,
		
		WeekCount: o.WeekCount,
		
		PlanningGroups: o.PlanningGroups,
		
		LongTermPlanningGroups: o.LongTermPlanningGroups,
		
		CanUseForScheduling: o.CanUseForScheduling,
		Alias:    (Alias)(o),
	})
}

func (o *Buimportshorttermforecastschema) UnmarshalJSON(b []byte) error {
	var BuimportshorttermforecastschemaMap map[string]interface{}
	err := json.Unmarshal(b, &BuimportshorttermforecastschemaMap)
	if err != nil {
		return err
	}
	
	if Description, ok := BuimportshorttermforecastschemaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if WeekCount, ok := BuimportshorttermforecastschemaMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if PlanningGroups, ok := BuimportshorttermforecastschemaMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if LongTermPlanningGroups, ok := BuimportshorttermforecastschemaMap["longTermPlanningGroups"].([]interface{}); ok {
		LongTermPlanningGroupsString, _ := json.Marshal(LongTermPlanningGroups)
		json.Unmarshal(LongTermPlanningGroupsString, &o.LongTermPlanningGroups)
	}
	
	if CanUseForScheduling, ok := BuimportshorttermforecastschemaMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buimportshorttermforecastschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
