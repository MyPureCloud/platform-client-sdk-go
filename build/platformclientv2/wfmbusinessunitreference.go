package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbusinessunitreference
type Wfmbusinessunitreference struct { 
	// Id - The ID of the business unit
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbusinessunitreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
