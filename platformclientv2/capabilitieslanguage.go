package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitieslanguage
type Capabilitieslanguage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsAvailable
	IsAvailable *bool `json:"isAvailable,omitempty"`

	// UnavailableReason
	UnavailableReason *Capabilitiesunavailablereason `json:"unavailableReason,omitempty"`

	// Contexts
	Contexts *Capabilitieslanguagecontextsmodel `json:"contexts,omitempty"`

	// Fallback
	Fallback *string `json:"fallback,omitempty"`

	// Features
	Features *Capabilitieslanguagefeatures `json:"features,omitempty"`

	// Functions
	Functions *map[string]Capabilitiesfunction `json:"functions,omitempty"`

	// IsDefault
	IsDefault *bool `json:"isDefault,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitieslanguage) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitieslanguage) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitieslanguage
	
	return json.Marshal(&struct { 
		IsAvailable *bool `json:"isAvailable,omitempty"`
		
		UnavailableReason *Capabilitiesunavailablereason `json:"unavailableReason,omitempty"`
		
		Contexts *Capabilitieslanguagecontextsmodel `json:"contexts,omitempty"`
		
		Fallback *string `json:"fallback,omitempty"`
		
		Features *Capabilitieslanguagefeatures `json:"features,omitempty"`
		
		Functions *map[string]Capabilitiesfunction `json:"functions,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		Alias
	}{ 
		IsAvailable: o.IsAvailable,
		
		UnavailableReason: o.UnavailableReason,
		
		Contexts: o.Contexts,
		
		Fallback: o.Fallback,
		
		Features: o.Features,
		
		Functions: o.Functions,
		
		IsDefault: o.IsDefault,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitieslanguage) UnmarshalJSON(b []byte) error {
	var CapabilitieslanguageMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitieslanguageMap)
	if err != nil {
		return err
	}
	
	if IsAvailable, ok := CapabilitieslanguageMap["isAvailable"].(bool); ok {
		o.IsAvailable = &IsAvailable
	}
    
	if UnavailableReason, ok := CapabilitieslanguageMap["unavailableReason"].(map[string]interface{}); ok {
		UnavailableReasonString, _ := json.Marshal(UnavailableReason)
		json.Unmarshal(UnavailableReasonString, &o.UnavailableReason)
	}
	
	if Contexts, ok := CapabilitieslanguageMap["contexts"].(map[string]interface{}); ok {
		ContextsString, _ := json.Marshal(Contexts)
		json.Unmarshal(ContextsString, &o.Contexts)
	}
	
	if Fallback, ok := CapabilitieslanguageMap["fallback"].(string); ok {
		o.Fallback = &Fallback
	}
    
	if Features, ok := CapabilitieslanguageMap["features"].(map[string]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	
	if Functions, ok := CapabilitieslanguageMap["functions"].(map[string]interface{}); ok {
		FunctionsString, _ := json.Marshal(Functions)
		json.Unmarshal(FunctionsString, &o.Functions)
	}
	
	if IsDefault, ok := CapabilitieslanguageMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitieslanguage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
