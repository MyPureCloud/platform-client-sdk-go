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

func (o *Scimv2group) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Schemas: o.Schemas,
		
		DisplayName: o.DisplayName,
		
		ExternalId: o.ExternalId,
		
		Members: o.Members,
		
		Meta: o.Meta,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimv2group) UnmarshalJSON(b []byte) error {
	var Scimv2groupMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2groupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := Scimv2groupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schemas, ok := Scimv2groupMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if DisplayName, ok := Scimv2groupMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if ExternalId, ok := Scimv2groupMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Members, ok := Scimv2groupMap["members"].([]interface{}); ok {
		MembersString, _ := json.Marshal(Members)
		json.Unmarshal(MembersString, &o.Members)
	}
	
	if Meta, ok := Scimv2groupMap["meta"].(map[string]interface{}); ok {
		MetaString, _ := json.Marshal(Meta)
		json.Unmarshal(MetaString, &o.Meta)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2group) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
