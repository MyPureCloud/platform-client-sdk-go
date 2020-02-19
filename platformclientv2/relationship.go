package platformclientv2
import (
	"encoding/json"
)

// Relationship
type Relationship struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User - The user associated with the external organization
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

// String returns a JSON representation of the model
func (o *Relationship) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
