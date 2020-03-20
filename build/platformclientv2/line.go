package platformclientv2
import (
	"time"
	"encoding/json"
)

// Line
type Line struct { 
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


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// EdgeGroup
	EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`


	// Template
	Template *Domainentityref `json:"template,omitempty"`


	// Site
	Site *Domainentityref `json:"site,omitempty"`


	// LineBaseSettings
	LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`


	// PrimaryEdge - The primary edge associated to the line. (Deprecated)
	PrimaryEdge *Edge `json:"primaryEdge,omitempty"`


	// SecondaryEdge - The secondary edge associated to the line. (Deprecated)
	SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`


	// LoggedInUser
	LoggedInUser *Domainentityref `json:"loggedInUser,omitempty"`


	// DefaultForUser
	DefaultForUser *Domainentityref `json:"defaultForUser,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Line) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
