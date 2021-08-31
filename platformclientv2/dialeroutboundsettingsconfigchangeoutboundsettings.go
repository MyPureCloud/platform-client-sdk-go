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
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// MaxCallsPerAgent
	MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`


	// MaxLineUtilization
	MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`


	// AbandonSeconds
	AbandonSeconds *float32 `json:"abandonSeconds,omitempty"`


	// ComplianceAbandonRateDenominator
	ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`


	// AutomaticTimeZoneMapping
	AutomaticTimeZoneMapping *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`

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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`
		
		MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`
		
		AbandonSeconds *float32 `json:"abandonSeconds,omitempty"`
		
		ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`
		
		AutomaticTimeZoneMapping *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		MaxCallsPerAgent: o.MaxCallsPerAgent,
		
		MaxLineUtilization: o.MaxLineUtilization,
		
		AbandonSeconds: o.AbandonSeconds,
		
		ComplianceAbandonRateDenominator: o.ComplianceAbandonRateDenominator,
		
		AutomaticTimeZoneMapping: o.AutomaticTimeZoneMapping,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeoutboundsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeoutboundsettingsMap)
	if err != nil {
		return err
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
