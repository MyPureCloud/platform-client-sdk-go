package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Credentialspecification - Specifies the requirements for a credential that can be provided for configuring an integration
type Credentialspecification struct { 
	// Required - Indicates if the credential must be provided in order for the integration configuration to be valid.
	Required *bool `json:"required,omitempty"`


	// Title - Title describing the usage for this credential.
	Title *string `json:"title,omitempty"`


	// CredentialTypes - List of acceptable credential types that can be provided for this credential.
	CredentialTypes *[]string `json:"credentialTypes,omitempty"`

}

func (o *Credentialspecification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Credentialspecification
	
	return json.Marshal(&struct { 
		Required *bool `json:"required,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		CredentialTypes *[]string `json:"credentialTypes,omitempty"`
		*Alias
	}{ 
		Required: o.Required,
		
		Title: o.Title,
		
		CredentialTypes: o.CredentialTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Credentialspecification) UnmarshalJSON(b []byte) error {
	var CredentialspecificationMap map[string]interface{}
	err := json.Unmarshal(b, &CredentialspecificationMap)
	if err != nil {
		return err
	}
	
	if Required, ok := CredentialspecificationMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if Title, ok := CredentialspecificationMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if CredentialTypes, ok := CredentialspecificationMap["credentialTypes"].([]interface{}); ok {
		CredentialTypesString, _ := json.Marshal(CredentialTypes)
		json.Unmarshal(CredentialTypesString, &o.CredentialTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Credentialspecification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
