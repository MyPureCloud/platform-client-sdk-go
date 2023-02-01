package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationtype - Descriptor for a type of Integration.
type Integrationtype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Integrationtype) SetField(field string, fieldValue interface{}) {
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

func (o Integrationtype) MarshalJSON() ([]byte, error) {
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
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Provider: o.Provider,
		
		Category: o.Category,
		
		Images: o.Images,
		
		ConfigPropertiesSchemaUri: o.ConfigPropertiesSchemaUri,
		
		ConfigAdvancedSchemaUri: o.ConfigAdvancedSchemaUri,
		
		HelpUri: o.HelpUri,
		
		TermsOfServiceUri: o.TermsOfServiceUri,
		
		VendorName: o.VendorName,
		
		VendorWebsiteUri: o.VendorWebsiteUri,
		
		MarketplaceUri: o.MarketplaceUri,
		
		FaqUri: o.FaqUri,
		
		PrivacyPolicyUri: o.PrivacyPolicyUri,
		
		SupportContactUri: o.SupportContactUri,
		
		SalesContactUri: o.SalesContactUri,
		
		HelpLinks: o.HelpLinks,
		
		Credentials: o.Credentials,
		
		NonInstallable: o.NonInstallable,
		
		MaxInstances: o.MaxInstances,
		
		UserPermissions: o.UserPermissions,
		
		VendorOAuthClientIds: o.VendorOAuthClientIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Integrationtype) UnmarshalJSON(b []byte) error {
	var IntegrationtypeMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationtypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := IntegrationtypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := IntegrationtypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := IntegrationtypeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Provider, ok := IntegrationtypeMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Category, ok := IntegrationtypeMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Images, ok := IntegrationtypeMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if ConfigPropertiesSchemaUri, ok := IntegrationtypeMap["configPropertiesSchemaUri"].(string); ok {
		o.ConfigPropertiesSchemaUri = &ConfigPropertiesSchemaUri
	}
    
	if ConfigAdvancedSchemaUri, ok := IntegrationtypeMap["configAdvancedSchemaUri"].(string); ok {
		o.ConfigAdvancedSchemaUri = &ConfigAdvancedSchemaUri
	}
    
	if HelpUri, ok := IntegrationtypeMap["helpUri"].(string); ok {
		o.HelpUri = &HelpUri
	}
    
	if TermsOfServiceUri, ok := IntegrationtypeMap["termsOfServiceUri"].(string); ok {
		o.TermsOfServiceUri = &TermsOfServiceUri
	}
    
	if VendorName, ok := IntegrationtypeMap["vendorName"].(string); ok {
		o.VendorName = &VendorName
	}
    
	if VendorWebsiteUri, ok := IntegrationtypeMap["vendorWebsiteUri"].(string); ok {
		o.VendorWebsiteUri = &VendorWebsiteUri
	}
    
	if MarketplaceUri, ok := IntegrationtypeMap["marketplaceUri"].(string); ok {
		o.MarketplaceUri = &MarketplaceUri
	}
    
	if FaqUri, ok := IntegrationtypeMap["faqUri"].(string); ok {
		o.FaqUri = &FaqUri
	}
    
	if PrivacyPolicyUri, ok := IntegrationtypeMap["privacyPolicyUri"].(string); ok {
		o.PrivacyPolicyUri = &PrivacyPolicyUri
	}
    
	if SupportContactUri, ok := IntegrationtypeMap["supportContactUri"].(string); ok {
		o.SupportContactUri = &SupportContactUri
	}
    
	if SalesContactUri, ok := IntegrationtypeMap["salesContactUri"].(string); ok {
		o.SalesContactUri = &SalesContactUri
	}
    
	if HelpLinks, ok := IntegrationtypeMap["helpLinks"].([]interface{}); ok {
		HelpLinksString, _ := json.Marshal(HelpLinks)
		json.Unmarshal(HelpLinksString, &o.HelpLinks)
	}
	
	if Credentials, ok := IntegrationtypeMap["credentials"].(map[string]interface{}); ok {
		CredentialsString, _ := json.Marshal(Credentials)
		json.Unmarshal(CredentialsString, &o.Credentials)
	}
	
	if NonInstallable, ok := IntegrationtypeMap["nonInstallable"].(bool); ok {
		o.NonInstallable = &NonInstallable
	}
    
	if MaxInstances, ok := IntegrationtypeMap["maxInstances"].(float64); ok {
		MaxInstancesInt := int(MaxInstances)
		o.MaxInstances = &MaxInstancesInt
	}
	
	if UserPermissions, ok := IntegrationtypeMap["userPermissions"].([]interface{}); ok {
		UserPermissionsString, _ := json.Marshal(UserPermissions)
		json.Unmarshal(UserPermissionsString, &o.UserPermissions)
	}
	
	if VendorOAuthClientIds, ok := IntegrationtypeMap["vendorOAuthClientIds"].([]interface{}); ok {
		VendorOAuthClientIdsString, _ := json.Marshal(VendorOAuthClientIds)
		json.Unmarshal(VendorOAuthClientIdsString, &o.VendorOAuthClientIds)
	}
	
	if SelfUri, ok := IntegrationtypeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
