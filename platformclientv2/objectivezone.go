package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Objectivezone
type Objectivezone struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Label - label
	Label *string `json:"label,omitempty"`

	// DirectionType - direction type
	DirectionType *string `json:"directionType,omitempty"`

	// ZoneType - zone type
	ZoneType *string `json:"zoneType,omitempty"`

	// UpperLimitPoints - upper limit points
	UpperLimitPoints *int `json:"upperLimitPoints,omitempty"`

	// LowerLimitPoints - lower limit points
	LowerLimitPoints *int `json:"lowerLimitPoints,omitempty"`

	// UpperLimitValue - upper limit value
	UpperLimitValue *int `json:"upperLimitValue,omitempty"`

	// LowerLimitValue - lower limit value
	LowerLimitValue *int `json:"lowerLimitValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Objectivezone) SetField(field string, fieldValue interface{}) {
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

func (o Objectivezone) MarshalJSON() ([]byte, error) {
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
	type Alias Objectivezone
	
	return json.Marshal(&struct { 
		Label *string `json:"label,omitempty"`
		
		DirectionType *string `json:"directionType,omitempty"`
		
		ZoneType *string `json:"zoneType,omitempty"`
		
		UpperLimitPoints *int `json:"upperLimitPoints,omitempty"`
		
		LowerLimitPoints *int `json:"lowerLimitPoints,omitempty"`
		
		UpperLimitValue *int `json:"upperLimitValue,omitempty"`
		
		LowerLimitValue *int `json:"lowerLimitValue,omitempty"`
		Alias
	}{ 
		Label: o.Label,
		
		DirectionType: o.DirectionType,
		
		ZoneType: o.ZoneType,
		
		UpperLimitPoints: o.UpperLimitPoints,
		
		LowerLimitPoints: o.LowerLimitPoints,
		
		UpperLimitValue: o.UpperLimitValue,
		
		LowerLimitValue: o.LowerLimitValue,
		Alias:    (Alias)(o),
	})
}

func (o *Objectivezone) UnmarshalJSON(b []byte) error {
	var ObjectivezoneMap map[string]interface{}
	err := json.Unmarshal(b, &ObjectivezoneMap)
	if err != nil {
		return err
	}
	
	if Label, ok := ObjectivezoneMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if DirectionType, ok := ObjectivezoneMap["directionType"].(string); ok {
		o.DirectionType = &DirectionType
	}
    
	if ZoneType, ok := ObjectivezoneMap["zoneType"].(string); ok {
		o.ZoneType = &ZoneType
	}
    
	if UpperLimitPoints, ok := ObjectivezoneMap["upperLimitPoints"].(float64); ok {
		UpperLimitPointsInt := int(UpperLimitPoints)
		o.UpperLimitPoints = &UpperLimitPointsInt
	}
	
	if LowerLimitPoints, ok := ObjectivezoneMap["lowerLimitPoints"].(float64); ok {
		LowerLimitPointsInt := int(LowerLimitPoints)
		o.LowerLimitPoints = &LowerLimitPointsInt
	}
	
	if UpperLimitValue, ok := ObjectivezoneMap["upperLimitValue"].(float64); ok {
		UpperLimitValueInt := int(UpperLimitValue)
		o.UpperLimitValue = &UpperLimitValueInt
	}
	
	if LowerLimitValue, ok := ObjectivezoneMap["lowerLimitValue"].(float64); ok {
		LowerLimitValueInt := int(LowerLimitValue)
		o.LowerLimitValue = &LowerLimitValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Objectivezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
