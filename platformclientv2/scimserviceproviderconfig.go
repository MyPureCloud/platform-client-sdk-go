package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfig - Defines a SCIM service provider's configuration.
type Scimserviceproviderconfig struct { 
	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// DocumentationUri - The HTTP-addressable URL that points to the service provider's documentation.
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// Patch - The \"patch\" configuration options.
	Patch *Scimserviceproviderconfigsimplefeature `json:"patch,omitempty"`


	// Filter - The \"filter\" configuration options.
	Filter *Scimserviceproviderconfigfilterfeature `json:"filter,omitempty"`


	// Etag - The \"etag\" configuration options.
	Etag *Scimserviceproviderconfigsimplefeature `json:"etag,omitempty"`


	// Sort - The \"sort\" configuration options.
	Sort *Scimserviceproviderconfigsimplefeature `json:"sort,omitempty"`


	// Bulk - The \"bulk\" configuration options.
	Bulk *Scimserviceproviderconfigbulkfeature `json:"bulk,omitempty"`


	// ChangePassword - The \"changePassword\" configuration options.
	ChangePassword *Scimserviceproviderconfigsimplefeature `json:"changePassword,omitempty"`


	// AuthenticationSchemes - The list of supported authentication schemes.
	AuthenticationSchemes *[]Scimserviceproviderconfigauthenticationscheme `json:"authenticationSchemes,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

func (o *Scimserviceproviderconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfig
	
	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		Patch *Scimserviceproviderconfigsimplefeature `json:"patch,omitempty"`
		
		Filter *Scimserviceproviderconfigfilterfeature `json:"filter,omitempty"`
		
		Etag *Scimserviceproviderconfigsimplefeature `json:"etag,omitempty"`
		
		Sort *Scimserviceproviderconfigsimplefeature `json:"sort,omitempty"`
		
		Bulk *Scimserviceproviderconfigbulkfeature `json:"bulk,omitempty"`
		
		ChangePassword *Scimserviceproviderconfigsimplefeature `json:"changePassword,omitempty"`
		
		AuthenticationSchemes *[]Scimserviceproviderconfigauthenticationscheme `json:"authenticationSchemes,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Schemas: o.Schemas,
		
		DocumentationUri: o.DocumentationUri,
		
		Patch: o.Patch,
		
		Filter: o.Filter,
		
		Etag: o.Etag,
		
		Sort: o.Sort,
		
		Bulk: o.Bulk,
		
		ChangePassword: o.ChangePassword,
		
		AuthenticationSchemes: o.AuthenticationSchemes,
		
		Meta: o.Meta,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimserviceproviderconfig) UnmarshalJSON(b []byte) error {
	var ScimserviceproviderconfigMap map[string]interface{}
	err := json.Unmarshal(b, &ScimserviceproviderconfigMap)
	if err != nil {
		return err
	}
	
	if Schemas, ok := ScimserviceproviderconfigMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if DocumentationUri, ok := ScimserviceproviderconfigMap["documentationUri"].(string); ok {
		o.DocumentationUri = &DocumentationUri
	}
    
	if Patch, ok := ScimserviceproviderconfigMap["patch"].(map[string]interface{}); ok {
		PatchString, _ := json.Marshal(Patch)
		json.Unmarshal(PatchString, &o.Patch)
	}
	
	if Filter, ok := ScimserviceproviderconfigMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Etag, ok := ScimserviceproviderconfigMap["etag"].(map[string]interface{}); ok {
		EtagString, _ := json.Marshal(Etag)
		json.Unmarshal(EtagString, &o.Etag)
	}
	
	if Sort, ok := ScimserviceproviderconfigMap["sort"].(map[string]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Bulk, ok := ScimserviceproviderconfigMap["bulk"].(map[string]interface{}); ok {
		BulkString, _ := json.Marshal(Bulk)
		json.Unmarshal(BulkString, &o.Bulk)
	}
	
	if ChangePassword, ok := ScimserviceproviderconfigMap["changePassword"].(map[string]interface{}); ok {
		ChangePasswordString, _ := json.Marshal(ChangePassword)
		json.Unmarshal(ChangePasswordString, &o.ChangePassword)
	}
	
	if AuthenticationSchemes, ok := ScimserviceproviderconfigMap["authenticationSchemes"].([]interface{}); ok {
		AuthenticationSchemesString, _ := json.Marshal(AuthenticationSchemes)
		json.Unmarshal(AuthenticationSchemesString, &o.AuthenticationSchemes)
	}
	
	if Meta, ok := ScimserviceproviderconfigMap["meta"].(map[string]interface{}); ok {
		MetaString, _ := json.Marshal(Meta)
		json.Unmarshal(MetaString, &o.Meta)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
