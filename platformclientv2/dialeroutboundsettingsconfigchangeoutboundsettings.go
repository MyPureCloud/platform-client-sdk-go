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

func (u *Dialeroutboundsettingsconfigchangeoutboundsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeoutboundsettings

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		MaxCallsPerAgent: u.MaxCallsPerAgent,
		
		MaxLineUtilization: u.MaxLineUtilization,
		
		AbandonSeconds: u.AbandonSeconds,
		
		ComplianceAbandonRateDenominator: u.ComplianceAbandonRateDenominator,
		
		AutomaticTimeZoneMapping: u.AutomaticTimeZoneMapping,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
