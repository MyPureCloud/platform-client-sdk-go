package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeoutboundsettings
type Dialeroutboundsettingsconfigchangeoutboundsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MaxCallsPerAgent - The maximum number of calls that can be placed per agent on any campaign
	MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`

	// MaxLineUtilization - The maximum percentage of lines that should be used for Outbound, expressed as a decimal in the range [0.0, 1.0]
	MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`

	// AbandonSeconds - The number of seconds used to determine if a call is abandoned
	AbandonSeconds *float32 `json:"abandonSeconds,omitempty"`

	// ComplianceAbandonRateDenominator - The denominator to be used in determining the compliance abandon rate
	ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`

	// AutomaticTimeZoneMapping
	AutomaticTimeZoneMapping *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`

	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) SetField(field string, fieldValue interface{}) {
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

func (o Dialeroutboundsettingsconfigchangeoutboundsettings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Dialeroutboundsettingsconfigchangeoutboundsettings
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`
		
		MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`
		
		AbandonSeconds *float32 `json:"abandonSeconds,omitempty"`
		
		ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`
		
		AutomaticTimeZoneMapping *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		MaxCallsPerAgent: o.MaxCallsPerAgent,
		
		MaxLineUtilization: o.MaxLineUtilization,
		
		AbandonSeconds: o.AbandonSeconds,
		
		ComplianceAbandonRateDenominator: o.ComplianceAbandonRateDenominator,
		
		AutomaticTimeZoneMapping: o.AutomaticTimeZoneMapping,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeoutboundsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeoutboundsettingsMap)
	if err != nil {
		return err
	}
	
	if MaxCallsPerAgent, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["maxCallsPerAgent"].(float64); ok {
		MaxCallsPerAgentInt := int(MaxCallsPerAgent)
		o.MaxCallsPerAgent = &MaxCallsPerAgentInt
	}
	
	if MaxLineUtilization, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["maxLineUtilization"].(float64); ok {
		MaxLineUtilizationFloat32 := float32(MaxLineUtilization)
		o.MaxLineUtilization = &MaxLineUtilizationFloat32
	}
    
	if AbandonSeconds, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["abandonSeconds"].(float64); ok {
		AbandonSecondsFloat32 := float32(AbandonSeconds)
		o.AbandonSeconds = &AbandonSecondsFloat32
	}
    
	if ComplianceAbandonRateDenominator, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["complianceAbandonRateDenominator"].(string); ok {
		o.ComplianceAbandonRateDenominator = &ComplianceAbandonRateDenominator
	}
    
	if AutomaticTimeZoneMapping, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["automaticTimeZoneMapping"].(map[string]interface{}); ok {
		AutomaticTimeZoneMappingString, _ := json.Marshal(AutomaticTimeZoneMapping)
		json.Unmarshal(AutomaticTimeZoneMappingString, &o.AutomaticTimeZoneMapping)
	}
	
	if Id, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialeroutboundsettingsconfigchangeoutboundsettingsMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
