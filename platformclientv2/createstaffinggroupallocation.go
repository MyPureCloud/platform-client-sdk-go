package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createstaffinggroupallocation
type Createstaffinggroupallocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StaffingGroupId - The ID of the staffing group used in this capacity plan
	StaffingGroupId *string `json:"staffingGroupId,omitempty"`

	// InitialShrinkagePercentage - The shrinkage percentage of the staffing group that can be used for all weeks in the planning period, in the scale of 0 - 100
	InitialShrinkagePercentage *float64 `json:"initialShrinkagePercentage,omitempty"`

	// InitialAttritionPercentage - The attrition percentage of the staffing group that can be used for all weeks in the planning period, in the scale of 0 - 100
	InitialAttritionPercentage *float64 `json:"initialAttritionPercentage,omitempty"`

	// StartingWeeklyFullTimeEquivalentCount - The weekly count of full time equivalent agents in the staffing group that can be used for the first week of the planning period
	StartingWeeklyFullTimeEquivalentCount *float64 `json:"startingWeeklyFullTimeEquivalentCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createstaffinggroupallocation) SetField(field string, fieldValue interface{}) {
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

func (o Createstaffinggroupallocation) MarshalJSON() ([]byte, error) {
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
	type Alias Createstaffinggroupallocation
	
	return json.Marshal(&struct { 
		StaffingGroupId *string `json:"staffingGroupId,omitempty"`
		
		InitialShrinkagePercentage *float64 `json:"initialShrinkagePercentage,omitempty"`
		
		InitialAttritionPercentage *float64 `json:"initialAttritionPercentage,omitempty"`
		
		StartingWeeklyFullTimeEquivalentCount *float64 `json:"startingWeeklyFullTimeEquivalentCount,omitempty"`
		Alias
	}{ 
		StaffingGroupId: o.StaffingGroupId,
		
		InitialShrinkagePercentage: o.InitialShrinkagePercentage,
		
		InitialAttritionPercentage: o.InitialAttritionPercentage,
		
		StartingWeeklyFullTimeEquivalentCount: o.StartingWeeklyFullTimeEquivalentCount,
		Alias:    (Alias)(o),
	})
}

func (o *Createstaffinggroupallocation) UnmarshalJSON(b []byte) error {
	var CreatestaffinggroupallocationMap map[string]interface{}
	err := json.Unmarshal(b, &CreatestaffinggroupallocationMap)
	if err != nil {
		return err
	}
	
	if StaffingGroupId, ok := CreatestaffinggroupallocationMap["staffingGroupId"].(string); ok {
		o.StaffingGroupId = &StaffingGroupId
	}
    
	if InitialShrinkagePercentage, ok := CreatestaffinggroupallocationMap["initialShrinkagePercentage"].(float64); ok {
		o.InitialShrinkagePercentage = &InitialShrinkagePercentage
	}
    
	if InitialAttritionPercentage, ok := CreatestaffinggroupallocationMap["initialAttritionPercentage"].(float64); ok {
		o.InitialAttritionPercentage = &InitialAttritionPercentage
	}
    
	if StartingWeeklyFullTimeEquivalentCount, ok := CreatestaffinggroupallocationMap["startingWeeklyFullTimeEquivalentCount"].(float64); ok {
		o.StartingWeeklyFullTimeEquivalentCount = &StartingWeeklyFullTimeEquivalentCount
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createstaffinggroupallocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
