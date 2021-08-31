package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Relationship
type Relationship struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// User - The user associated with the external organization. When creating or updating a relationship, only User.id is required. User object is fully populated when expanding a note.
	User *User `json:"user,omitempty"`


	// ExternalOrganization - The external organization this relationship is attached to
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`


	// Relationship - The relationship or role of the user to this external organization.Examples: Account Manager, Sales Engineer, Implementation Consultant
	Relationship *string `json:"relationship,omitempty"`


	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Relationship) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Relationship
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		Relationship *string `json:"relationship,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Relationship: o.Relationship,
		
		ExternalDataSources: o.ExternalDataSources,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Relationship) UnmarshalJSON(b []byte) error {
	var RelationshipMap map[string]interface{}
	err := json.Unmarshal(b, &RelationshipMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RelationshipMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := RelationshipMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ExternalOrganization, ok := RelationshipMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Relationship, ok := RelationshipMap["relationship"].(string); ok {
		o.Relationship = &Relationship
	}
	
	if ExternalDataSources, ok := RelationshipMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if SelfUri, ok := RelationshipMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Relationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
