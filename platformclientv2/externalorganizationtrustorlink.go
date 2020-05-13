package platformclientv2
import (
	"time"
	"encoding/json"
)

// Externalorganizationtrustorlink
type Externalorganizationtrustorlink struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ExternalOrganizationId - The id of a PureCloud External Organization entity in the External Contacts system that will be used to represent the trustor org
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// TrustorOrgId - The id of a PureCloud organization that has granted trust to this PureCloud organization
	TrustorOrgId *string `json:"trustorOrgId,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
