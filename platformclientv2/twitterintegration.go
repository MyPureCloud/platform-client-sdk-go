package platformclientv2
import (
	"time"
	"encoding/json"
)

// Twitterintegration
type Twitterintegration struct { 
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`


	// Name - The name of the Twitter Integration
	Name *string `json:"name,omitempty"`


	// AccessTokenKey - The Access Token Key from Twitter messenger
	AccessTokenKey *string `json:"accessTokenKey,omitempty"`


	// ConsumerKey - The Consumer Key from Twitter messenger
	ConsumerKey *string `json:"consumerKey,omitempty"`


	// Username - The Username from Twitter
	Username *string `json:"username,omitempty"`


	// UserId - The UserId from Twitter
	UserId *string `json:"userId,omitempty"`


	// Status - The status of the Twitter Integration
	Status *string `json:"status,omitempty"`


	// Tier - The type of twitter account to be used for the integration
	Tier *string `json:"tier,omitempty"`


	// EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
	EnvName *string `json:"envName,omitempty"`


	// Recipient - The recipient associated to the Twitter Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int32 `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Twitterintegration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
