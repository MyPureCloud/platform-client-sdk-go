package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchdocument
type Knowledgesearchdocument struct { 
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


	// Confidence - The confidence associated with a document with respect to a search query
	Confidence *float64 `json:"confidence,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgesearchdocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgesearchdocument
	
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
		
		Confidence *float64 `json:"confidence,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		LanguageCode: o.LanguageCode,
		
		VarType: o.VarType,
		
		Faq: o.Faq,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Categories: o.Categories,
		
		KnowledgeBase: o.KnowledgeBase,
		
		ExternalUrl: o.ExternalUrl,
		
		Article: o.Article,
		
		Confidence: o.Confidence,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgesearchdocument) UnmarshalJSON(b []byte) error {
	var KnowledgesearchdocumentMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchdocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgesearchdocumentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KnowledgesearchdocumentMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if LanguageCode, ok := KnowledgesearchdocumentMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
	
	if VarType, ok := KnowledgesearchdocumentMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Faq, ok := KnowledgesearchdocumentMap["faq"].(map[string]interface{}); ok {
		FaqString, _ := json.Marshal(Faq)
		json.Unmarshal(FaqString, &o.Faq)
	}
	
	if dateCreatedString, ok := KnowledgesearchdocumentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgesearchdocumentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Categories, ok := KnowledgesearchdocumentMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if KnowledgeBase, ok := KnowledgesearchdocumentMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if ExternalUrl, ok := KnowledgesearchdocumentMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
	
	if Article, ok := KnowledgesearchdocumentMap["article"].(map[string]interface{}); ok {
		ArticleString, _ := json.Marshal(Article)
		json.Unmarshal(ArticleString, &o.Article)
	}
	
	if Confidence, ok := KnowledgesearchdocumentMap["confidence"].(float64); ok {
		o.Confidence = &Confidence
	}
	
	if SelfUri, ok := KnowledgesearchdocumentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchdocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
