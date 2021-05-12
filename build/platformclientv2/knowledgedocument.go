package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocument
type Knowledgedocument struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// LanguageCode - Language of the document
	LanguageCode *string `json:"languageCode,omitempty"`


	// VarType - Document type
	VarType *string `json:"type,omitempty"`


	// Faq - FAQ document details
	Faq *Documentfaq `json:"faq,omitempty"`


	// DateCreated - Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Categories - Document categories
	Categories *[]Knowledgecategory `json:"categories,omitempty"`


	// KnowledgeBase - Knowledge base which document does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// ExternalUrl - External URL to the document
	ExternalUrl *string `json:"externalUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
