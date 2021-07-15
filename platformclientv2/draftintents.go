package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Draftintents
type Draftintents struct { 
	// Id - Id for an intent.
	Id *string `json:"id,omitempty"`


	// Name - Name/Label for an intent.
	Name *string `json:"name,omitempty"`


	// Utterances - The utterances that are extracted for an Intent.
	Utterances *[]string `json:"utterances,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Draftintents) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
