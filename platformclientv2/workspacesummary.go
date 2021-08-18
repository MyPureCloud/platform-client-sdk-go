package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Workspacesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspacesummary

	

	return json.Marshal(&struct { 
		TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`
		
		TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`
		*Alias
	}{ 
		TotalDocumentCount: u.TotalDocumentCount,
		
		TotalDocumentByteCount: u.TotalDocumentByteCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workspacesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
