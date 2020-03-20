package platformclientv2
import (
	"encoding/json"
)

// Scimv2user - Defines a SCIM user.
type Scimv2user struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"Mutability\" is set to \"readOnly\". \"Returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// Active - Indicates whether the user's administrative status is active.
	Active *bool `json:"active,omitempty"`


	// UserName - The user's PureCloud email address. Must be unique.
	UserName *string `json:"userName,omitempty"`


	// DisplayName - The display name of the user.
	DisplayName *string `json:"displayName,omitempty"`


	// Password - The new password for the PureCloud user. Does not return an existing password.
	Password *string `json:"password,omitempty"`


	// Title - The user's title.
	Title *string `json:"title,omitempty"`


	// PhoneNumbers - The list of the user's phone numbers.
	PhoneNumbers *[]Scimphonenumber `json:"phoneNumbers,omitempty"`


	// Emails - The list of the user's email addresses.
	Emails *[]Scimemail `json:"emails,omitempty"`


	// ExternalId - The external ID of the user. Set by the provisioning client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
	ExternalId *string `json:"externalId,omitempty"`


	// Groups - The list of groups that the user is a member of.
	Groups *[]Scimv2groupreference `json:"groups,omitempty"`


	// Roles - The list of roles assigned to the user.
	Roles *[]Scimuserrole `json:"roles,omitempty"`


	// UrnIetfParamsScimSchemasExtensionEnterprise20User - The URI of the schema for the enterprise user.
	UrnIetfParamsScimSchemasExtensionEnterprise20User *Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`


	// UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User
	UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2user) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
