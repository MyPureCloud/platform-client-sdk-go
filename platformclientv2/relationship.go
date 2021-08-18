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

func (u *Relationship) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		User: u.User,
		
		ExternalOrganization: u.ExternalOrganization,
		
		Relationship: u.Relationship,
		
		ExternalDataSources: u.ExternalDataSources,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Relationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
