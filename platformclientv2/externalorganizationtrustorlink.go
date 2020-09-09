package platformclientv2
import (
	"time"
	"encoding/json"
)

// Externalorganizationtrustorlink
type Externalorganizationtrustorlink struct { 
	// ExternalOrganizationId - The id of a PureCloud External Organization entity in the External Contacts system that will be used to represent the trustor org
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// TrustorOrgId - The id of a PureCloud organization that has granted trust to this PureCloud organization
	TrustorOrgId *string `json:"trustorOrgId,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ExternalOrganizationUri - The URI for the External Organization that is linked to the trustor org
	ExternalOrganizationUri *string `json:"externalOrganizationUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
