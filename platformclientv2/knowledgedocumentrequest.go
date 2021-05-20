package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Knowledgedocumentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
