package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Supportedcontentreference - Reference to supported content profile associated with the integration
type Supportedcontentreference struct { 
	// Id - The SupportedContent unique identifier associated with this integration
	Id *string `json:"id,omitempty"`


	// Name - The SupportedContent profile name
	Name *string `json:"name,omitempty"`


	// SelfUri - The SupportedContent profile URI
	SelfUri *string `json:"selfUri,omitempty"`


	// MediaTypes - Media types definition for the supported content
	MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Supportedcontentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
