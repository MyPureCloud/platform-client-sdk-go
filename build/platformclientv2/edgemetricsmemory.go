package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricsmemory
type Edgemetricsmemory struct { 
	// AvailableBytes - Available memory in bytes.
	AvailableBytes *float64 `json:"availableBytes,omitempty"`


	// VarType - Type of memory. Virtual or physical.
	VarType *string `json:"type,omitempty"`


	// TotalBytes - Total memory in bytes.
	TotalBytes *float64 `json:"totalBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsmemory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
