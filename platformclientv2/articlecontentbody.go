package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Articlecontentbody
type Articlecontentbody struct { 
	// LocationUrl - Presigned URL to retrieve the document content.
	LocationUrl *string `json:"locationUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Articlecontentbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
