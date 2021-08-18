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

func (u *Knowledgedocumentbulkrequest) MarshalJSON() ([]byte, error) {
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
		VarType: u.VarType,
		
		ExternalUrl: u.ExternalUrl,
		
		Faq: u.Faq,
		
		Categories: u.Categories,
		
		Article: u.Article,
		
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
