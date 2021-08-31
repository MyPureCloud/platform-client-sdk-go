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

func (o *Workspacesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspacesummary
	
	return json.Marshal(&struct { 
		TotalDocumentCount *int `json:"totalDocumentCount,omitempty"`
		
		TotalDocumentByteCount *int `json:"totalDocumentByteCount,omitempty"`
		*Alias
	}{ 
		TotalDocumentCount: o.TotalDocumentCount,
		
		TotalDocumentByteCount: o.TotalDocumentByteCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Workspacesummary) UnmarshalJSON(b []byte) error {
	var WorkspacesummaryMap map[string]interface{}
	err := json.Unmarshal(b, &WorkspacesummaryMap)
	if err != nil {
		return err
	}
	
	if TotalDocumentCount, ok := WorkspacesummaryMap["totalDocumentCount"].(float64); ok {
		TotalDocumentCountInt := int(TotalDocumentCount)
		o.TotalDocumentCount = &TotalDocumentCountInt
	}
	
	if TotalDocumentByteCount, ok := WorkspacesummaryMap["totalDocumentByteCount"].(float64); ok {
		TotalDocumentByteCountInt := int(TotalDocumentByteCount)
		o.TotalDocumentByteCount = &TotalDocumentByteCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workspacesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
