package platformclientv2
import (
	"time"
	"encoding/json"
)

// Oauthauthorization
type Oauthauthorization struct { 
	// Client
	Client *Oauthclient `json:"client,omitempty"`


	// Scope
	Scope *[]string `json:"scope,omitempty"`


	// ResourceOwner
	ResourceOwner *Domainentityref `json:"resourceOwner,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Pending
	Pending *bool `json:"pending,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthauthorization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
