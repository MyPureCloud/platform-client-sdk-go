package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesdatatypefacetsmodel
type Capabilitiesdatatypefacetsmodel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MaxItems
	MaxItems *int `json:"maxItems,omitempty"`

	// MinItems
	MinItems *int `json:"minItems,omitempty"`

	// Minimum
	Minimum *Capabilitiesdatatypefacetslimit `json:"minimum,omitempty"`

	// MinimumExclusive
	MinimumExclusive *bool `json:"minimumExclusive,omitempty"`

	// Maximum
	Maximum *Capabilitiesdatatypefacetslimit `json:"maximum,omitempty"`

	// MaximumExclusive
	MaximumExclusive *bool `json:"maximumExclusive,omitempty"`

	// MinLength
	MinLength *int `json:"minLength,omitempty"`

	// MaxLength
	MaxLength *int `json:"maxLength,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesdatatypefacetsmodel) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesdatatypefacetsmodel) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesdatatypefacetsmodel
	
	return json.Marshal(&struct { 
		MaxItems *int `json:"maxItems,omitempty"`
		
		MinItems *int `json:"minItems,omitempty"`
		
		Minimum *Capabilitiesdatatypefacetslimit `json:"minimum,omitempty"`
		
		MinimumExclusive *bool `json:"minimumExclusive,omitempty"`
		
		Maximum *Capabilitiesdatatypefacetslimit `json:"maximum,omitempty"`
		
		MaximumExclusive *bool `json:"maximumExclusive,omitempty"`
		
		MinLength *int `json:"minLength,omitempty"`
		
		MaxLength *int `json:"maxLength,omitempty"`
		Alias
	}{ 
		MaxItems: o.MaxItems,
		
		MinItems: o.MinItems,
		
		Minimum: o.Minimum,
		
		MinimumExclusive: o.MinimumExclusive,
		
		Maximum: o.Maximum,
		
		MaximumExclusive: o.MaximumExclusive,
		
		MinLength: o.MinLength,
		
		MaxLength: o.MaxLength,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesdatatypefacetsmodel) UnmarshalJSON(b []byte) error {
	var CapabilitiesdatatypefacetsmodelMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesdatatypefacetsmodelMap)
	if err != nil {
		return err
	}
	
	if MaxItems, ok := CapabilitiesdatatypefacetsmodelMap["maxItems"].(float64); ok {
		MaxItemsInt := int(MaxItems)
		o.MaxItems = &MaxItemsInt
	}
	
	if MinItems, ok := CapabilitiesdatatypefacetsmodelMap["minItems"].(float64); ok {
		MinItemsInt := int(MinItems)
		o.MinItems = &MinItemsInt
	}
	
	if Minimum, ok := CapabilitiesdatatypefacetsmodelMap["minimum"].(map[string]interface{}); ok {
		MinimumString, _ := json.Marshal(Minimum)
		json.Unmarshal(MinimumString, &o.Minimum)
	}
	
	if MinimumExclusive, ok := CapabilitiesdatatypefacetsmodelMap["minimumExclusive"].(bool); ok {
		o.MinimumExclusive = &MinimumExclusive
	}
    
	if Maximum, ok := CapabilitiesdatatypefacetsmodelMap["maximum"].(map[string]interface{}); ok {
		MaximumString, _ := json.Marshal(Maximum)
		json.Unmarshal(MaximumString, &o.Maximum)
	}
	
	if MaximumExclusive, ok := CapabilitiesdatatypefacetsmodelMap["maximumExclusive"].(bool); ok {
		o.MaximumExclusive = &MaximumExclusive
	}
    
	if MinLength, ok := CapabilitiesdatatypefacetsmodelMap["minLength"].(float64); ok {
		MinLengthInt := int(MinLength)
		o.MinLength = &MinLengthInt
	}
	
	if MaxLength, ok := CapabilitiesdatatypefacetsmodelMap["maxLength"].(float64); ok {
		MaxLengthInt := int(MaxLength)
		o.MaxLength = &MaxLengthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypefacetsmodel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
