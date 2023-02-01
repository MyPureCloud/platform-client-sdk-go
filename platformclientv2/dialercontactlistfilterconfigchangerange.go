package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangerange - FilterRange is one of the attributes of a FilterPredicate
type Dialercontactlistfilterconfigchangerange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Min - Minimum end of the range
	Min *string `json:"min,omitempty"`

	// Max - Maximum end of the range
	Max *string `json:"max,omitempty"`

	// MinInclusive - Whether or not to include the minimum in the range
	MinInclusive *bool `json:"minInclusive,omitempty"`

	// MaxInclusive - Whether or not to include the maximum in the range
	MaxInclusive *bool `json:"maxInclusive,omitempty"`

	// InSet - Elements that apply to the IN operator
	InSet *[]string `json:"inSet,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercontactlistfilterconfigchangerange) SetField(field string, fieldValue interface{}) {
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

func (o Dialercontactlistfilterconfigchangerange) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercontactlistfilterconfigchangerange
	
	return json.Marshal(&struct { 
		Min *string `json:"min,omitempty"`
		
		Max *string `json:"max,omitempty"`
		
		MinInclusive *bool `json:"minInclusive,omitempty"`
		
		MaxInclusive *bool `json:"maxInclusive,omitempty"`
		
		InSet *[]string `json:"inSet,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		Alias
	}{ 
		Min: o.Min,
		
		Max: o.Max,
		
		MinInclusive: o.MinInclusive,
		
		MaxInclusive: o.MaxInclusive,
		
		InSet: o.InSet,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangerange) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangerangeMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangerangeMap)
	if err != nil {
		return err
	}
	
	if Min, ok := DialercontactlistfilterconfigchangerangeMap["min"].(string); ok {
		o.Min = &Min
	}
    
	if Max, ok := DialercontactlistfilterconfigchangerangeMap["max"].(string); ok {
		o.Max = &Max
	}
    
	if MinInclusive, ok := DialercontactlistfilterconfigchangerangeMap["minInclusive"].(bool); ok {
		o.MinInclusive = &MinInclusive
	}
    
	if MaxInclusive, ok := DialercontactlistfilterconfigchangerangeMap["maxInclusive"].(bool); ok {
		o.MaxInclusive = &MaxInclusive
	}
    
	if InSet, ok := DialercontactlistfilterconfigchangerangeMap["inSet"].([]interface{}); ok {
		InSetString, _ := json.Marshal(InSet)
		json.Unmarshal(InSetString, &o.InSet)
	}
	
	if AdditionalProperties, ok := DialercontactlistfilterconfigchangerangeMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangerange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
