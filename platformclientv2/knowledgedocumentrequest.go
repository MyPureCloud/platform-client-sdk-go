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

func (u *Knowledgedocumentrequest) MarshalJSON() ([]byte, error) {
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
		VarType: u.VarType,
		
		ExternalUrl: u.ExternalUrl,
		
		Faq: u.Faq,
		
		Categories: u.Categories,
		
		Article: u.Article,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
