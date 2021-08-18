package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationtype - Descriptor for a type of Integration.
type Integrationtype struct { 
	// Id - The ID of the integration type.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - Description of the integration type.
	Description *string `json:"description,omitempty"`


	// Provider - PureCloud provider of the integration type.
	Provider *string `json:"provider,omitempty"`


	// Category - Category describing the integration type.
	Category *string `json:"category,omitempty"`


	// Images - Collection of logos.
	Images *[]Userimage `json:"images,omitempty"`


	// ConfigPropertiesSchemaUri - URI of the schema describing the key-value properties needed to configure an integration of this type.
	ConfigPropertiesSchemaUri *string `json:"configPropertiesSchemaUri,omitempty"`


	// ConfigAdvancedSchemaUri - URI of the schema describing the advanced JSON document needed to configure an integration of this type.
	ConfigAdvancedSchemaUri *string `json:"configAdvancedSchemaUri,omitempty"`


	// HelpUri - URI of a page with more information about the integration type
	HelpUri *string `json:"helpUri,omitempty"`


	// TermsOfServiceUri - URI of a page with terms and conditions for the integration type
	TermsOfServiceUri *string `json:"termsOfServiceUri,omitempty"`


	// VendorName - Name of the vendor of this integration type
	VendorName *string `json:"vendorName,omitempty"`


	// VendorWebsiteUri - URI of the vendor's website
	VendorWebsiteUri *string `json:"vendorWebsiteUri,omitempty"`


	// MarketplaceUri - URI of the marketplace listing for this integration type
	MarketplaceUri *string `json:"marketplaceUri,omitempty"`


	// FaqUri - URI of frequently asked questions about the integration type
	FaqUri *string `json:"faqUri,omitempty"`


	// PrivacyPolicyUri - URI of a privacy policy for users of the integration type
	PrivacyPolicyUri *string `json:"privacyPolicyUri,omitempty"`


	// SupportContactUri - URI for vendor support
	SupportContactUri *string `json:"supportContactUri,omitempty"`


	// SalesContactUri - URI for vendor sales information
	SalesContactUri *string `json:"salesContactUri,omitempty"`


	// HelpLinks - List of links to additional help resources
	HelpLinks *[]Helplink `json:"helpLinks,omitempty"`


	// Credentials - Map of credentials for integrations of this type. The key is the name of a credential that can be provided in the credentials property of the integration configuration.
	Credentials *map[string]Credentialspecification `json:"credentials,omitempty"`


	// NonInstallable - Indicates if the integration type is installable or not.
	NonInstallable *bool `json:"nonInstallable,omitempty"`


	// MaxInstances - The maximum number of integration instances allowable for this integration type
	MaxInstances *int `json:"maxInstances,omitempty"`


	// UserPermissions - List of permissions required to permit user access to the integration type.
	UserPermissions *[]string `json:"userPermissions,omitempty"`


	// VendorOAuthClientIds - List of OAuth Client IDs that must be authorized when the integration is created.
	VendorOAuthClientIds *[]string `json:"vendorOAuthClientIds,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Integrationtype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationtype

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		ConfigPropertiesSchemaUri *string `json:"configPropertiesSchemaUri,omitempty"`
		
		ConfigAdvancedSchemaUri *string `json:"configAdvancedSchemaUri,omitempty"`
		
		HelpUri *string `json:"helpUri,omitempty"`
		
		TermsOfServiceUri *string `json:"termsOfServiceUri,omitempty"`
		
		VendorName *string `json:"vendorName,omitempty"`
		
		VendorWebsiteUri *string `json:"vendorWebsiteUri,omitempty"`
		
		MarketplaceUri *string `json:"marketplaceUri,omitempty"`
		
		FaqUri *string `json:"faqUri,omitempty"`
		
		PrivacyPolicyUri *string `json:"privacyPolicyUri,omitempty"`
		
		SupportContactUri *string `json:"supportContactUri,omitempty"`
		
		SalesContactUri *string `json:"salesContactUri,omitempty"`
		
		HelpLinks *[]Helplink `json:"helpLinks,omitempty"`
		
		Credentials *map[string]Credentialspecification `json:"credentials,omitempty"`
		
		NonInstallable *bool `json:"nonInstallable,omitempty"`
		
		MaxInstances *int `json:"maxInstances,omitempty"`
		
		UserPermissions *[]string `json:"userPermissions,omitempty"`
		
		VendorOAuthClientIds *[]string `json:"vendorOAuthClientIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Provider: u.Provider,
		
		Category: u.Category,
		
		Images: u.Images,
		
		ConfigPropertiesSchemaUri: u.ConfigPropertiesSchemaUri,
		
		ConfigAdvancedSchemaUri: u.ConfigAdvancedSchemaUri,
		
		HelpUri: u.HelpUri,
		
		TermsOfServiceUri: u.TermsOfServiceUri,
		
		VendorName: u.VendorName,
		
		VendorWebsiteUri: u.VendorWebsiteUri,
		
		MarketplaceUri: u.MarketplaceUri,
		
		FaqUri: u.FaqUri,
		
		PrivacyPolicyUri: u.PrivacyPolicyUri,
		
		SupportContactUri: u.SupportContactUri,
		
		SalesContactUri: u.SalesContactUri,
		
		HelpLinks: u.HelpLinks,
		
		Credentials: u.Credentials,
		
		NonInstallable: u.NonInstallable,
		
		MaxInstances: u.MaxInstances,
		
		UserPermissions: u.UserPermissions,
		
		VendorOAuthClientIds: u.VendorOAuthClientIds,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Integrationtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
