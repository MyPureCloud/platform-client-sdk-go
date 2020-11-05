package platformclientv2
import (
	"time"
	"encoding/json"
)

// Knowledgebase
type Knowledgebase struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - Knowledge base description
	Description *string `json:"description,omitempty"`


	// CoreLanguage - Core language for knowledge base in which initial content must be created first
	CoreLanguage *string `json:"coreLanguage,omitempty"`


	// DateCreated - Knowledge base creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Knowledge base last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// FaqCount - The count representing the number of documents of type FAQ per KnowledgeBase
	FaqCount *int32 `json:"faqCount,omitempty"`


	// DateDocumentLastModified - The date representing when the last document is modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDocumentLastModified *time.Time `json:"dateDocumentLastModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgebase) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
