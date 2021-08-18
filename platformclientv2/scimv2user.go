package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2user - Defines a SCIM user.
type Scimv2user struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// Active - Indicates whether the user's administrative status is active.
	Active *bool `json:"active,omitempty"`


	// UserName - The user's Genesys Cloud email address. Must be unique.
	UserName *string `json:"userName,omitempty"`


	// DisplayName - The display name of the user.
	DisplayName *string `json:"displayName,omitempty"`


	// Password - The new password for the Genesys Cloud user. Does not return an existing password. When creating a user, if a password is not supplied, then a password will be randomly generated that is 40 characters in length and contains five characters from each of the password policy groups.
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


	// UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User - The URI of the schema for the Genesys Cloud user.
	UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

func (u *Scimv2user) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2user

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schemas *[]string `json:"schemas,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		UserName *string `json:"userName,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Password *string `json:"password,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		PhoneNumbers *[]Scimphonenumber `json:"phoneNumbers,omitempty"`
		
		Emails *[]Scimemail `json:"emails,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Groups *[]Scimv2groupreference `json:"groups,omitempty"`
		
		Roles *[]Scimuserrole `json:"roles,omitempty"`
		
		UrnIetfParamsScimSchemasExtensionEnterprise20User *Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
		
		UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Schemas: u.Schemas,
		
		Active: u.Active,
		
		UserName: u.UserName,
		
		DisplayName: u.DisplayName,
		
		Password: u.Password,
		
		Title: u.Title,
		
		PhoneNumbers: u.PhoneNumbers,
		
		Emails: u.Emails,
		
		ExternalId: u.ExternalId,
		
		Groups: u.Groups,
		
		Roles: u.Roles,
		
		UrnIetfParamsScimSchemasExtensionEnterprise20User: u.UrnIetfParamsScimSchemasExtensionEnterprise20User,
		
		UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User: u.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User,
		
		Meta: u.Meta,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2user) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
