package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2schemadefinition - Defines a SCIM schema.
type Scimv2schemadefinition struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Name - The name of the schema.
	Name *string `json:"name,omitempty"`


	// Description - The description of the schema.
	Description *string `json:"description,omitempty"`


	// Attributes - The list of service provider attributes.
	Attributes *[]Scimv2schemaattribute `json:"attributes,omitempty"`


	// Meta - The metadata of the SCIM resource. Only \"location\" and \"resourceType\" are set for \"Schema\" resources.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

func (u *Scimv2schemadefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2schemadefinition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Attributes *[]Scimv2schemaattribute `json:"attributes,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Attributes: u.Attributes,
		
		Meta: u.Meta,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2schemadefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
