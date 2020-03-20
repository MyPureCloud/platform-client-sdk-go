package platformclientv2
import (
	"time"
	"encoding/json"
)

// Workspace
type Workspace struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The current name of the workspace.
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// IsCurrentUserWorkspace
	IsCurrentUserWorkspace *bool `json:"isCurrentUserWorkspace,omitempty"`


	// User
	User *Domainentityref `json:"user,omitempty"`


	// Bucket
	Bucket *string `json:"bucket,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Summary
	Summary *Workspacesummary `json:"summary,omitempty"`


	// Acl
	Acl *[]string `json:"acl,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workspace) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
