package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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


	// Article - Article
	Article *Documentarticle `json:"article,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Knowledgedocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocument

	
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
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Faq *Documentfaq `json:"faq,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Categories *[]Knowledgecategory `json:"categories,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		Article *Documentarticle `json:"article,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		LanguageCode: u.LanguageCode,
		
		VarType: u.VarType,
		
		Faq: u.Faq,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Categories: u.Categories,
		
		KnowledgeBase: u.KnowledgeBase,
		
		ExternalUrl: u.ExternalUrl,
		
		Article: u.Article,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgedocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
