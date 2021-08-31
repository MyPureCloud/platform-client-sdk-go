package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentrequest
type Knowledgedocumentrequest struct { 
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

}

func (o *Knowledgedocumentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentrequest
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		Faq *Documentfaq `json:"faq,omitempty"`
		
		Categories *[]Documentcategoryinput `json:"categories,omitempty"`
		
		Article *Documentarticle `json:"article,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ExternalUrl: o.ExternalUrl,
		
		Faq: o.Faq,
		
		Categories: o.Categories,
		
		Article: o.Article,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentrequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := KnowledgedocumentrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ExternalUrl, ok := KnowledgedocumentrequestMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
	
	if Faq, ok := KnowledgedocumentrequestMap["faq"].(map[string]interface{}); ok {
		FaqString, _ := json.Marshal(Faq)
		json.Unmarshal(FaqString, &o.Faq)
	}
	
	if Categories, ok := KnowledgedocumentrequestMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if Article, ok := KnowledgedocumentrequestMap["article"].(map[string]interface{}); ok {
		ArticleString, _ := json.Marshal(Article)
		json.Unmarshal(ArticleString, &o.Article)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
