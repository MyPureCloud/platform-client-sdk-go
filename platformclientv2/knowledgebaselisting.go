package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Knowledgebaselisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgebaselisting

	

	return json.Marshal(&struct { 
		Entities *[]Knowledgebase `json:"entities,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		NextUri: u.NextUri,
		
		SelfUri: u.SelfUri,
		
		PreviousUri: u.PreviousUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgebaselisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
