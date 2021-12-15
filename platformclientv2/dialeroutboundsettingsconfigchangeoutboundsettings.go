package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeoutboundsettings
type Dialeroutboundsettingsconfigchangeoutboundsettings struct { 
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

func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
