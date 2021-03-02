package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopicedgemetricdisk
type Edgemetricstopicedgemetricdisk struct { 
	// PartitionName
	PartitionName *string `json:"partitionName,omitempty"`


	// AvailableBytes
	AvailableBytes *int `json:"availableBytes,omitempty"`


	// TotalBytes
	TotalBytes *int `json:"totalBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricdisk) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
