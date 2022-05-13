package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigauthenticationscheme - Defines an authentication scheme in the SCIM service provider's configuration.
type Scimserviceproviderconfigauthenticationscheme struct { 
	// Name - The name of the authentication scheme, for example, HTTP Basic.
	Name *string `json:"name,omitempty"`


	// Description - The description of the authentication scheme.
	Description *string `json:"description,omitempty"`


	// SpecUri - The HTTP-addressable URL that points to the authentication scheme's specification.
	SpecUri *string `json:"specUri,omitempty"`


	// DocumentationUri - The HTTP-addressable URL that points to the authentication scheme's usage documentation.
	DocumentationUri *string `json:"documentationUri,omitempty"`


	// VarType - The type of authentication scheme.
	VarType *string `json:"type,omitempty"`


	// Primary - Indicates whether this authentication scheme is the primary method of authentication.
	Primary *bool `json:"primary,omitempty"`

}

func (o *Scimserviceproviderconfigauthenticationscheme) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigauthenticationscheme
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SpecUri *string `json:"specUri,omitempty"`
		
		DocumentationUri *string `json:"documentationUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		SpecUri: o.SpecUri,
		
		DocumentationUri: o.DocumentationUri,
		
		VarType: o.VarType,
		
		Primary: o.Primary,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimserviceproviderconfigauthenticationscheme) UnmarshalJSON(b []byte) error {
	var ScimserviceproviderconfigauthenticationschemeMap map[string]interface{}
	err := json.Unmarshal(b, &ScimserviceproviderconfigauthenticationschemeMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ScimserviceproviderconfigauthenticationschemeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ScimserviceproviderconfigauthenticationschemeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SpecUri, ok := ScimserviceproviderconfigauthenticationschemeMap["specUri"].(string); ok {
		o.SpecUri = &SpecUri
	}
    
	if DocumentationUri, ok := ScimserviceproviderconfigauthenticationschemeMap["documentationUri"].(string); ok {
		o.DocumentationUri = &DocumentationUri
	}
    
	if VarType, ok := ScimserviceproviderconfigauthenticationschemeMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Primary, ok := ScimserviceproviderconfigauthenticationschemeMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigauthenticationscheme) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
