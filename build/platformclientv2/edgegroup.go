package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edgegroup
type Edgegroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int32 `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
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


	// Managed - Is this edge group being managed remotely.
	Managed *bool `json:"managed,omitempty"`


	// EdgeTrunkBaseAssignment - A trunk base settings assignment of trunkType \"EDGE\" to use for edge-to-edge communication.
	EdgeTrunkBaseAssignment *Trunkbaseassignment `json:"edgeTrunkBaseAssignment,omitempty"`


	// PhoneTrunkBases - Trunk base settings of trunkType \"PHONE\" to inherit to edge logical interface for phone communication.
	PhoneTrunkBases *[]Trunkbase `json:"phoneTrunkBases,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgegroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
