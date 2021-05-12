package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Integration - Details for an Integration
type Integration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Type of the integration
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// IntendedState - Configured state of the integration.
	IntendedState *string `json:"intendedState,omitempty"`


	// Config - Configuration information for the integration.
	Config *Integrationconfigurationinfo `json:"config,omitempty"`


	// ReportedState - Last reported status of the integration.
	ReportedState *Integrationstatusinfo `json:"reportedState,omitempty"`


	// Attributes - Read-only attributes for the integration.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
