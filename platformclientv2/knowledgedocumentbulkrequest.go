package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentbulkrequest
type Knowledgedocumentbulkrequest struct { 
	// VarType - Document type according to assigned template
	VarType *string `json:"type,omitempty"`


	// ExternalUrl - External Url to the document
	ExternalUrl *string `json:"externalUrl,omitempty"`


	// Faq - Faq document details
	Faq *Documentfaq `json:"faq,omitempty"`


	// Categories - Document categories
	Categories *[]Documentcategoryinput `json:"categories,omitempty"`


	// Article - Article details
	Article *Documentarticle `json:"article,omitempty"`


	// Id - Identifier of document for update. Omit for create new Document.
	Id *string `json:"id,omitempty"`

}

func (o *Knowledgedocumentbulkrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentbulkrequest
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		Faq *Documentfaq `json:"faq,omitempty"`
		
		Categories *[]Documentcategoryinput `json:"categories,omitempty"`
		
		Article *Documentarticle `json:"article,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ExternalUrl: o.ExternalUrl,
		
		Faq: o.Faq,
		
		Categories: o.Categories,
		
		Article: o.Article,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentbulkrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentbulkrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentbulkrequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := KnowledgedocumentbulkrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ExternalUrl, ok := KnowledgedocumentbulkrequestMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
	
	if Faq, ok := KnowledgedocumentbulkrequestMap["faq"].(map[string]interface{}); ok {
		FaqString, _ := json.Marshal(Faq)
		json.Unmarshal(FaqString, &o.Faq)
	}
	
	if Categories, ok := KnowledgedocumentbulkrequestMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if Article, ok := KnowledgedocumentbulkrequestMap["article"].(map[string]interface{}); ok {
		ArticleString, _ := json.Marshal(Article)
		json.Unmarshal(ArticleString, &o.Article)
	}
	
	if Id, ok := KnowledgedocumentbulkrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
