package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usageitem
type Usageitem struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// TotalDocumentByteCount
	TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`


	// TotalDocumentCount
	TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`

}

func (u *Usageitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usageitem

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`
		
		TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		TotalDocumentByteCount: u.TotalDocumentByteCount,
		
		TotalDocumentCount: u.TotalDocumentCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usageitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
