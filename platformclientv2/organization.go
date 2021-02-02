package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Organization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
