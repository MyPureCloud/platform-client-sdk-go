package platformclientv2
import (
	"encoding/json"
)

// Workspacesummary
type Workspacesummary struct { 
	// TotalDocumentCount
	TotalDocumentCount *int64 `json:"totalDocumentCount,omitempty"`


	// TotalDocumentByteCount
	TotalDocumentByteCount *int64 `json:"totalDocumentByteCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workspacesummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
