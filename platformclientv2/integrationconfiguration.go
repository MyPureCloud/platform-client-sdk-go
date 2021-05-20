package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationconfiguration - Configuration for an Integration
type Integrationconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
	Properties *interface{} `json:"properties,omitempty"`


	// Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
	Advanced *interface{} `json:"advanced,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
	Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
