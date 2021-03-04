package platformclientv2
import (
	"time"
	"encoding/json"
)

// Program
type Program struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// Topics
	Topics *[]Basetopicentitiy `json:"topics,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// PublishedBy
	PublishedBy *Addressableentityref `json:"publishedBy,omitempty"`


	// DatePublished - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Program) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
