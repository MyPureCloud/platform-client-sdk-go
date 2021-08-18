package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimconfigresourcetype - Defines a SCIM resource.
type Scimconfigresourcetype struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// Name - The name of the resource type.
	Name *string `json:"name,omitempty"`


	// Description - The description of the resource type.
	Description *string `json:"description,omitempty"`


	// Schema - The URI of the primary or base schema for the resource type.
	Schema *string `json:"schema,omitempty"`


	// SchemaExtensions - The list of schema extensions for the resource type.
	SchemaExtensions *[]Scimconfigresourcetypeschemaextension `json:"schemaExtensions,omitempty"`


	// Endpoint - The HTTP-addressable endpoint of the resource type. Appears after the base URL.
	Endpoint *string `json:"endpoint,omitempty"`


	// Meta - The metadata of the SCIM resource. Only \"location\" and \"resourceType\" are set for \"ResourceType\" resources.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

func (u *Scimconfigresourcetype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimconfigresourcetype

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schemas *[]string `json:"schemas,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Schema *string `json:"schema,omitempty"`
		
		SchemaExtensions *[]Scimconfigresourcetypeschemaextension `json:"schemaExtensions,omitempty"`
		
		Endpoint *string `json:"endpoint,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Schemas: u.Schemas,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Schema: u.Schema,
		
		SchemaExtensions: u.SchemaExtensions,
		
		Endpoint: u.Endpoint,
		
		Meta: u.Meta,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
