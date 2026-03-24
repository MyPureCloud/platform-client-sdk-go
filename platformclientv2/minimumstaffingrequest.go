package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Minimumstaffingrequest
type Minimumstaffingrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether the setting is turned on or off
	Enabled *bool `json:"enabled,omitempty"`

	// MinimumValue - Default minimum staff value to be applied to all planning groups
	MinimumValue *float64 `json:"minimumValue,omitempty"`

	// PlanningGroupOverrides - List of planning groups with their minimum staff value overrides and the days to which the overrides apply. Setting the enclosed list to empty will clear out all existing overrides
	PlanningGroupOverrides *Listwrapperplanninggroupminimumsrequest `json:"planningGroupOverrides,omitempty"`

	// ApplicableIntervals - The intervals to which the minimum staff values will apply
	ApplicableIntervals *string `json:"applicableIntervals,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Minimumstaffingrequest) SetField(field string, fieldValue interface{}) {
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

func (o Minimumstaffingrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Minimumstaffingrequest
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		MinimumValue *float64 `json:"minimumValue,omitempty"`
		
		PlanningGroupOverrides *Listwrapperplanninggroupminimumsrequest `json:"planningGroupOverrides,omitempty"`
		
		ApplicableIntervals *string `json:"applicableIntervals,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		MinimumValue: o.MinimumValue,
		
		PlanningGroupOverrides: o.PlanningGroupOverrides,
		
		ApplicableIntervals: o.ApplicableIntervals,
		Alias:    (Alias)(o),
	})
}

func (o *Minimumstaffingrequest) UnmarshalJSON(b []byte) error {
	var MinimumstaffingrequestMap map[string]interface{}
	err := json.Unmarshal(b, &MinimumstaffingrequestMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := MinimumstaffingrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MinimumValue, ok := MinimumstaffingrequestMap["minimumValue"].(float64); ok {
		o.MinimumValue = &MinimumValue
	}
    
	if PlanningGroupOverrides, ok := MinimumstaffingrequestMap["planningGroupOverrides"].(map[string]interface{}); ok {
		PlanningGroupOverridesString, _ := json.Marshal(PlanningGroupOverrides)
		json.Unmarshal(PlanningGroupOverridesString, &o.PlanningGroupOverrides)
	}
	
	if ApplicableIntervals, ok := MinimumstaffingrequestMap["applicableIntervals"].(string); ok {
		o.ApplicableIntervals = &ApplicableIntervals
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Minimumstaffingrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
