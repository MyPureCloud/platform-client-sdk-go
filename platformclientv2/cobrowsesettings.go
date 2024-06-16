package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Cobrowsesettings - Settings concerning cobrowse
type Cobrowsesettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether or not cobrowse is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// AllowAgentControl - Whether the viewer should have option to request control
	AllowAgentControl *bool `json:"allowAgentControl,omitempty"`

	// AllowAgentNavigation - Whether the viewer should have option to request navigation
	AllowAgentNavigation *bool `json:"allowAgentNavigation,omitempty"`

	// MaskSelectors - Mask patterns that will apply to pages being shared
	MaskSelectors *[]string `json:"maskSelectors,omitempty"`

	// Channels - Cobrowse channels for web messenger
	Channels *[]string `json:"channels,omitempty"`

	// ReadonlySelectors - Readonly patterns that will apply to pages being shared
	ReadonlySelectors *[]string `json:"readonlySelectors,omitempty"`

	// PauseCriteria - Pause criteria that will pause cobrowse if some of them are met in the user's URL
	PauseCriteria *[]Pausecriteria `json:"pauseCriteria,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Cobrowsesettings) SetField(field string, fieldValue interface{}) {
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

func (o Cobrowsesettings) MarshalJSON() ([]byte, error) {
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
	type Alias Cobrowsesettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AllowAgentControl *bool `json:"allowAgentControl,omitempty"`
		
		AllowAgentNavigation *bool `json:"allowAgentNavigation,omitempty"`
		
		MaskSelectors *[]string `json:"maskSelectors,omitempty"`
		
		Channels *[]string `json:"channels,omitempty"`
		
		ReadonlySelectors *[]string `json:"readonlySelectors,omitempty"`
		
		PauseCriteria *[]Pausecriteria `json:"pauseCriteria,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		AllowAgentControl: o.AllowAgentControl,
		
		AllowAgentNavigation: o.AllowAgentNavigation,
		
		MaskSelectors: o.MaskSelectors,
		
		Channels: o.Channels,
		
		ReadonlySelectors: o.ReadonlySelectors,
		
		PauseCriteria: o.PauseCriteria,
		Alias:    (Alias)(o),
	})
}

func (o *Cobrowsesettings) UnmarshalJSON(b []byte) error {
	var CobrowsesettingsMap map[string]interface{}
	err := json.Unmarshal(b, &CobrowsesettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := CobrowsesettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if AllowAgentControl, ok := CobrowsesettingsMap["allowAgentControl"].(bool); ok {
		o.AllowAgentControl = &AllowAgentControl
	}
    
	if AllowAgentNavigation, ok := CobrowsesettingsMap["allowAgentNavigation"].(bool); ok {
		o.AllowAgentNavigation = &AllowAgentNavigation
	}
    
	if MaskSelectors, ok := CobrowsesettingsMap["maskSelectors"].([]interface{}); ok {
		MaskSelectorsString, _ := json.Marshal(MaskSelectors)
		json.Unmarshal(MaskSelectorsString, &o.MaskSelectors)
	}
	
	if Channels, ok := CobrowsesettingsMap["channels"].([]interface{}); ok {
		ChannelsString, _ := json.Marshal(Channels)
		json.Unmarshal(ChannelsString, &o.Channels)
	}
	
	if ReadonlySelectors, ok := CobrowsesettingsMap["readonlySelectors"].([]interface{}); ok {
		ReadonlySelectorsString, _ := json.Marshal(ReadonlySelectors)
		json.Unmarshal(ReadonlySelectorsString, &o.ReadonlySelectors)
	}
	
	if PauseCriteria, ok := CobrowsesettingsMap["pauseCriteria"].([]interface{}); ok {
		PauseCriteriaString, _ := json.Marshal(PauseCriteria)
		json.Unmarshal(PauseCriteriaString, &o.PauseCriteria)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Cobrowsesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
