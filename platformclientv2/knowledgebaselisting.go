package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgebaselisting
type Knowledgebaselisting struct { 
	// Entities
	Entities *[]Knowledgebase `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgebaselisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
