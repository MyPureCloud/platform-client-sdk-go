package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Organization
type Organization struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DefaultLanguage - The default language for this organization. Example: 'en'
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`


	// DefaultCountryCode - The default country code for this organization. Example: 'US'
	DefaultCountryCode *string `json:"defaultCountryCode,omitempty"`


	// ThirdPartyOrgName - The short name for the organization. This field is globally unique and cannot be changed.
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`


	// ThirdPartyURI
	ThirdPartyURI *string `json:"thirdPartyURI,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// Version - The current version of the organization.
	Version *int `json:"version,omitempty"`


	// State - The current state. Examples are active, inactive, deleted.
	State *string `json:"state,omitempty"`


	// DefaultSiteId
	DefaultSiteId *string `json:"defaultSiteId,omitempty"`


	// SupportURI - Email address where support tickets are sent to.
	SupportURI *string `json:"supportURI,omitempty"`


	// VoicemailEnabled
	VoicemailEnabled *bool `json:"voicemailEnabled,omitempty"`


	// ProductPlatform - Organizations Originating Platform.
	ProductPlatform *string `json:"productPlatform,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Features - The state of features available for the organization.
	Features *map[string]bool `json:"features,omitempty"`

}

func (u *Organization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Organization

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DefaultLanguage *string `json:"defaultLanguage,omitempty"`
		
		DefaultCountryCode *string `json:"defaultCountryCode,omitempty"`
		
		ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`
		
		ThirdPartyURI *string `json:"thirdPartyURI,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DefaultSiteId *string `json:"defaultSiteId,omitempty"`
		
		SupportURI *string `json:"supportURI,omitempty"`
		
		VoicemailEnabled *bool `json:"voicemailEnabled,omitempty"`
		
		ProductPlatform *string `json:"productPlatform,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Features *map[string]bool `json:"features,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DefaultLanguage: u.DefaultLanguage,
		
		DefaultCountryCode: u.DefaultCountryCode,
		
		ThirdPartyOrgName: u.ThirdPartyOrgName,
		
		ThirdPartyURI: u.ThirdPartyURI,
		
		Domain: u.Domain,
		
		Version: u.Version,
		
		State: u.State,
		
		DefaultSiteId: u.DefaultSiteId,
		
		SupportURI: u.SupportURI,
		
		VoicemailEnabled: u.VoicemailEnabled,
		
		ProductPlatform: u.ProductPlatform,
		
		SelfUri: u.SelfUri,
		
		Features: u.Features,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Organization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
