package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2group - Defines a SCIM group.
type Scimv2group struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
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

func (u *Scimv2group) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2group

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schemas *[]string `json:"schemas,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Members *[]Scimv2memberreference `json:"members,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Schemas: u.Schemas,
		
		DisplayName: u.DisplayName,
		
		ExternalId: u.ExternalId,
		
		Members: u.Members,
		
		Meta: u.Meta,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2group) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
