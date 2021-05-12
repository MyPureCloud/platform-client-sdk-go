package platformclientv2
import (
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


	// Id - Identifier of document for update. Omit for create new Document.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
