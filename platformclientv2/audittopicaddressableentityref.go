package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicaddressableentityref
type Audittopicaddressableentityref struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Audittopicaddressableentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
