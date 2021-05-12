package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workspacesummary
type Workspacesummary struct { 
	// TotalDocumentCount
	TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`


	// TotalDocumentByteCount
	TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workspacesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
