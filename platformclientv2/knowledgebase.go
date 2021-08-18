package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// FaqCount - The count representing the number of documents of type FAQ in the KnowledgeBase
	FaqCount *int `json:"faqCount,omitempty"`


	// DateDocumentLastModified - The date representing when the last document is modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDocumentLastModified *time.Time `json:"dateDocumentLastModified,omitempty"`


	// ArticleCount - The count representing the number of documents of type Article in the KnowledgeBase
	ArticleCount *int `json:"articleCount,omitempty"`


	// Published - Flag that indicates the knowledge base is published
	Published *bool `json:"published,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Knowledgebase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgebase

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDocumentLastModified := new(string)
	if u.DateDocumentLastModified != nil {
		
		*DateDocumentLastModified = timeutil.Strftime(u.DateDocumentLastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDocumentLastModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CoreLanguage *string `json:"coreLanguage,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		FaqCount *int `json:"faqCount,omitempty"`
		
		DateDocumentLastModified *string `json:"dateDocumentLastModified,omitempty"`
		
		ArticleCount *int `json:"articleCount,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		CoreLanguage: u.CoreLanguage,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		FaqCount: u.FaqCount,
		
		DateDocumentLastModified: DateDocumentLastModified,
		
		ArticleCount: u.ArticleCount,
		
		Published: u.Published,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgebase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
