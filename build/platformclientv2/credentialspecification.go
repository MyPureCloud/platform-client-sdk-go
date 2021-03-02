package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Credentialspecification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
