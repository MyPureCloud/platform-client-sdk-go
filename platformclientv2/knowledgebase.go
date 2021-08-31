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

func (o *Knowledgebase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgebase
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDocumentLastModified := new(string)
	if o.DateDocumentLastModified != nil {
		
		*DateDocumentLastModified = timeutil.Strftime(o.DateDocumentLastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CoreLanguage: o.CoreLanguage,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		FaqCount: o.FaqCount,
		
		DateDocumentLastModified: DateDocumentLastModified,
		
		ArticleCount: o.ArticleCount,
		
		Published: o.Published,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgebase) UnmarshalJSON(b []byte) error {
	var KnowledgebaseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgebaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgebaseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KnowledgebaseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := KnowledgebaseMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if CoreLanguage, ok := KnowledgebaseMap["coreLanguage"].(string); ok {
		o.CoreLanguage = &CoreLanguage
	}
	
	if dateCreatedString, ok := KnowledgebaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgebaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if FaqCount, ok := KnowledgebaseMap["faqCount"].(float64); ok {
		FaqCountInt := int(FaqCount)
		o.FaqCount = &FaqCountInt
	}
	
	if dateDocumentLastModifiedString, ok := KnowledgebaseMap["dateDocumentLastModified"].(string); ok {
		DateDocumentLastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDocumentLastModifiedString)
		o.DateDocumentLastModified = &DateDocumentLastModified
	}
	
	if ArticleCount, ok := KnowledgebaseMap["articleCount"].(float64); ok {
		ArticleCountInt := int(ArticleCount)
		o.ArticleCount = &ArticleCountInt
	}
	
	if Published, ok := KnowledgebaseMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if SelfUri, ok := KnowledgebaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgebase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
