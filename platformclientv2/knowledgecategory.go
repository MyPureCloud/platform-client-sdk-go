package platformclientv2
import (
	"time"
	"encoding/json"
)

// Knowledgecategory
type Knowledgecategory struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Category name
	Name *string `json:"name,omitempty"`


	// Description - Category description
	Description *string `json:"description,omitempty"`


	// KnowledgeBase - Knowledge base which category does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// LanguageCode - Actual language of the category
	LanguageCode *string `json:"languageCode,omitempty"`


	// DateCreated - Category creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Category last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgecategory) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
