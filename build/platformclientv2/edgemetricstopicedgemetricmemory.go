package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopicedgemetricmemory
type Edgemetricstopicedgemetricmemory struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// AvailableBytes
	AvailableBytes *int `json:"availableBytes,omitempty"`


	// TotalBytes
	TotalBytes *int `json:"totalBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricmemory) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
