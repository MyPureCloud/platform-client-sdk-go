package platformclientv2
import (
	"encoding/json"
)

// Scimv2group - Defines a SCIM group.
type Scimv2group struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"Mutability\" is set to \"readOnly\". \"Returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// DisplayName - The display name of the group.
	DisplayName *string `json:"displayName,omitempty"`


	// ExternalId - The external ID of the group. Set by the provisioning client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
	ExternalId *string `json:"externalId,omitempty"`


	// Members - The list of members in the group.
	Members *[]Scimv2memberreference `json:"members,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2group) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
