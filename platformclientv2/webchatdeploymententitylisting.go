package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatdeploymententitylisting
type Webchatdeploymententitylisting struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Webchatdeployment `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatdeploymententitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
