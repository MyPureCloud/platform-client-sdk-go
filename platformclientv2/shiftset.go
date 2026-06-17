package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shiftset
type Shiftset struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the shift set
	Id *string `json:"id,omitempty"`

	// Name - The name given for the shift set
	Name *string `json:"name,omitempty"`

	// EffectiveWorkPlan - The work plan or work plan rotation used for generating the shift set
	EffectiveWorkPlan *Shiftseteffectiveworkplan `json:"effectiveWorkPlan,omitempty"`

	// Shifts - The scheduled shifts
	Shifts *[]Schedulebidscheduledshift `json:"shifts,omitempty"`

	// SuggestedAgentCount - The suggested agent count
	SuggestedAgentCount *int `json:"suggestedAgentCount,omitempty"`

	// OverrideAgentCount - The override agent count. If it is null, it falls back to using the suggestedAgentCount
	OverrideAgentCount *int `json:"overrideAgentCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shiftset) SetField(field string, fieldValue interface{}) {
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

func (o Shiftset) MarshalJSON() ([]byte, error) {
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
	type Alias Shiftset
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EffectiveWorkPlan *Shiftseteffectiveworkplan `json:"effectiveWorkPlan,omitempty"`
		
		Shifts *[]Schedulebidscheduledshift `json:"shifts,omitempty"`
		
		SuggestedAgentCount *int `json:"suggestedAgentCount,omitempty"`
		
		OverrideAgentCount *int `json:"overrideAgentCount,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		EffectiveWorkPlan: o.EffectiveWorkPlan,
		
		Shifts: o.Shifts,
		
		SuggestedAgentCount: o.SuggestedAgentCount,
		
		OverrideAgentCount: o.OverrideAgentCount,
		Alias:    (Alias)(o),
	})
}

func (o *Shiftset) UnmarshalJSON(b []byte) error {
	var ShiftsetMap map[string]interface{}
	err := json.Unmarshal(b, &ShiftsetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShiftsetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ShiftsetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if EffectiveWorkPlan, ok := ShiftsetMap["effectiveWorkPlan"].(map[string]interface{}); ok {
		EffectiveWorkPlanString, _ := json.Marshal(EffectiveWorkPlan)
		json.Unmarshal(EffectiveWorkPlanString, &o.EffectiveWorkPlan)
	}
	
	if Shifts, ok := ShiftsetMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if SuggestedAgentCount, ok := ShiftsetMap["suggestedAgentCount"].(float64); ok {
		SuggestedAgentCountInt := int(SuggestedAgentCount)
		o.SuggestedAgentCount = &SuggestedAgentCountInt
	}
	
	if OverrideAgentCount, ok := ShiftsetMap["overrideAgentCount"].(float64); ok {
		OverrideAgentCountInt := int(OverrideAgentCount)
		o.OverrideAgentCount = &OverrideAgentCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shiftset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
