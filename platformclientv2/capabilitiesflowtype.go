package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesflowtype
type Capabilitiesflowtype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsAvailable
	IsAvailable *bool `json:"isAvailable,omitempty"`

	// UnavailableReason
	UnavailableReason *Capabilitiesunavailablereason `json:"unavailableReason,omitempty"`

	// Components
	Components *map[string]Capabilitiesflowcomponent `json:"components,omitempty"`

	// IsSecure
	IsSecure *bool `json:"isSecure,omitempty"`

	// SupportsLanguages
	SupportsLanguages *Capabilitiesflowtypelanguagesupport `json:"supportsLanguages,omitempty"`

	// UserCreatable
	UserCreatable *Capabilitiesavailability `json:"userCreatable,omitempty"`

	// MaxSize
	MaxSize *int `json:"maxSize,omitempty"`

	// UserVariableLimits - Flow type user variable limits
	UserVariableLimits *Capabilitiesflowtypeuservariablelimits `json:"userVariableLimits,omitempty"`

	// Limits - Limit information for this flow type
	Limits *Capabilitiesflowtypelimitinfo `json:"limits,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesflowtype) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesflowtype) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesflowtype
	
	return json.Marshal(&struct { 
		IsAvailable *bool `json:"isAvailable,omitempty"`
		
		UnavailableReason *Capabilitiesunavailablereason `json:"unavailableReason,omitempty"`
		
		Components *map[string]Capabilitiesflowcomponent `json:"components,omitempty"`
		
		IsSecure *bool `json:"isSecure,omitempty"`
		
		SupportsLanguages *Capabilitiesflowtypelanguagesupport `json:"supportsLanguages,omitempty"`
		
		UserCreatable *Capabilitiesavailability `json:"userCreatable,omitempty"`
		
		MaxSize *int `json:"maxSize,omitempty"`
		
		UserVariableLimits *Capabilitiesflowtypeuservariablelimits `json:"userVariableLimits,omitempty"`
		
		Limits *Capabilitiesflowtypelimitinfo `json:"limits,omitempty"`
		Alias
	}{ 
		IsAvailable: o.IsAvailable,
		
		UnavailableReason: o.UnavailableReason,
		
		Components: o.Components,
		
		IsSecure: o.IsSecure,
		
		SupportsLanguages: o.SupportsLanguages,
		
		UserCreatable: o.UserCreatable,
		
		MaxSize: o.MaxSize,
		
		UserVariableLimits: o.UserVariableLimits,
		
		Limits: o.Limits,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesflowtype) UnmarshalJSON(b []byte) error {
	var CapabilitiesflowtypeMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesflowtypeMap)
	if err != nil {
		return err
	}
	
	if IsAvailable, ok := CapabilitiesflowtypeMap["isAvailable"].(bool); ok {
		o.IsAvailable = &IsAvailable
	}
    
	if UnavailableReason, ok := CapabilitiesflowtypeMap["unavailableReason"].(map[string]interface{}); ok {
		UnavailableReasonString, _ := json.Marshal(UnavailableReason)
		json.Unmarshal(UnavailableReasonString, &o.UnavailableReason)
	}
	
	if Components, ok := CapabilitiesflowtypeMap["components"].(map[string]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	
	if IsSecure, ok := CapabilitiesflowtypeMap["isSecure"].(bool); ok {
		o.IsSecure = &IsSecure
	}
    
	if SupportsLanguages, ok := CapabilitiesflowtypeMap["supportsLanguages"].(map[string]interface{}); ok {
		SupportsLanguagesString, _ := json.Marshal(SupportsLanguages)
		json.Unmarshal(SupportsLanguagesString, &o.SupportsLanguages)
	}
	
	if UserCreatable, ok := CapabilitiesflowtypeMap["userCreatable"].(map[string]interface{}); ok {
		UserCreatableString, _ := json.Marshal(UserCreatable)
		json.Unmarshal(UserCreatableString, &o.UserCreatable)
	}
	
	if MaxSize, ok := CapabilitiesflowtypeMap["maxSize"].(float64); ok {
		MaxSizeInt := int(MaxSize)
		o.MaxSize = &MaxSizeInt
	}
	
	if UserVariableLimits, ok := CapabilitiesflowtypeMap["userVariableLimits"].(map[string]interface{}); ok {
		UserVariableLimitsString, _ := json.Marshal(UserVariableLimits)
		json.Unmarshal(UserVariableLimitsString, &o.UserVariableLimits)
	}
	
	if Limits, ok := CapabilitiesflowtypeMap["limits"].(map[string]interface{}); ok {
		LimitsString, _ := json.Marshal(Limits)
		json.Unmarshal(LimitsString, &o.Limits)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
