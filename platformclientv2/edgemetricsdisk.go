package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricsdisk
type Edgemetricsdisk struct { 
	// AvailableBytes - Available memory in bytes.
	AvailableBytes *float64 `json:"availableBytes,omitempty"`


	// PartitionName - Disk partition name.
	PartitionName *string `json:"partitionName,omitempty"`


	// TotalBytes - Total memory in bytes.
	TotalBytes *float64 `json:"totalBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsdisk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
