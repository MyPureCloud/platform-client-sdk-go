package platformclientv2
import (
	"time"
	"encoding/json"
)

// Outboundroute
type Outboundroute struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// ClassificationTypes - The site associated to the outbound route.
	ClassificationTypes *[]string `json:"classificationTypes,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Distribution
	Distribution *string `json:"distribution,omitempty"`


	// ExternalTrunkBases - Trunk base settings of trunkType \"EXTERNAL\".  This base must also be set on an edge logical interface for correct routing.
	ExternalTrunkBases *[]Domainentityref `json:"externalTrunkBases,omitempty"`


	// Site - The site associated to the outbound route.
	Site *Site `json:"site,omitempty"`


	// Managed - Is this outbound route being managed remotely.
	Managed *bool `json:"managed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundroute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
