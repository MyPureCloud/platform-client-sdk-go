package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundsettings
type Outboundsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// MaxCallsPerAgent - The maximum number of calls that can be placed per agent on any campaign
	MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`

	// MaxConfigurableCallsPerAgent - The maximum number of calls that can be configured to be placed per agent on any campaign
	MaxConfigurableCallsPerAgent *int `json:"maxConfigurableCallsPerAgent,omitempty"`

	// MaxLineUtilization - The maximum percentage of lines that should be used for Outbound, expressed as a decimal in the range [0.0, 1.0]
	MaxLineUtilization *float64 `json:"maxLineUtilization,omitempty"`

	// AbandonSeconds - The number of seconds used to determine if a call is abandoned
	AbandonSeconds *float64 `json:"abandonSeconds,omitempty"`

	// ComplianceAbandonRateDenominator - The denominator to be used in determining the compliance abandon rate
	ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`

	// AutomaticTimeZoneMapping - The settings for automatic time zone mapping. Note that changing these settings will change them for both voice and messaging campaigns.
	AutomaticTimeZoneMapping *Automatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundsettings) SetField(field string, fieldValue interface{}) {
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

func (o Outboundsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundsettings
	
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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`
		
		MaxConfigurableCallsPerAgent *int `json:"maxConfigurableCallsPerAgent,omitempty"`
		
		MaxLineUtilization *float64 `json:"maxLineUtilization,omitempty"`
		
		AbandonSeconds *float64 `json:"abandonSeconds,omitempty"`
		
		ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`
		
		AutomaticTimeZoneMapping *Automatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		MaxCallsPerAgent: o.MaxCallsPerAgent,
		
		MaxConfigurableCallsPerAgent: o.MaxConfigurableCallsPerAgent,
		
		MaxLineUtilization: o.MaxLineUtilization,
		
		AbandonSeconds: o.AbandonSeconds,
		
		ComplianceAbandonRateDenominator: o.ComplianceAbandonRateDenominator,
		
		AutomaticTimeZoneMapping: o.AutomaticTimeZoneMapping,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundsettings) UnmarshalJSON(b []byte) error {
	var OutboundsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundsettingsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutboundsettingsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OutboundsettingsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := OutboundsettingsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OutboundsettingsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := OutboundsettingsMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MaxCallsPerAgent, ok := OutboundsettingsMap["maxCallsPerAgent"].(float64); ok {
		MaxCallsPerAgentInt := int(MaxCallsPerAgent)
		o.MaxCallsPerAgent = &MaxCallsPerAgentInt
	}
	
	if MaxConfigurableCallsPerAgent, ok := OutboundsettingsMap["maxConfigurableCallsPerAgent"].(float64); ok {
		MaxConfigurableCallsPerAgentInt := int(MaxConfigurableCallsPerAgent)
		o.MaxConfigurableCallsPerAgent = &MaxConfigurableCallsPerAgentInt
	}
	
	if MaxLineUtilization, ok := OutboundsettingsMap["maxLineUtilization"].(float64); ok {
		o.MaxLineUtilization = &MaxLineUtilization
	}
    
	if AbandonSeconds, ok := OutboundsettingsMap["abandonSeconds"].(float64); ok {
		o.AbandonSeconds = &AbandonSeconds
	}
    
	if ComplianceAbandonRateDenominator, ok := OutboundsettingsMap["complianceAbandonRateDenominator"].(string); ok {
		o.ComplianceAbandonRateDenominator = &ComplianceAbandonRateDenominator
	}
    
	if AutomaticTimeZoneMapping, ok := OutboundsettingsMap["automaticTimeZoneMapping"].(map[string]interface{}); ok {
		AutomaticTimeZoneMappingString, _ := json.Marshal(AutomaticTimeZoneMapping)
		json.Unmarshal(AutomaticTimeZoneMappingString, &o.AutomaticTimeZoneMapping)
	}
	
	if SelfUri, ok := OutboundsettingsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
