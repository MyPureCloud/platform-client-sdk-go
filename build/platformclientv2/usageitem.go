package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Usageitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
