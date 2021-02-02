package platformclientv2
import (
	"time"
	"encoding/json"
)

// Userrecording
type Userrecording struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userrecording) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
