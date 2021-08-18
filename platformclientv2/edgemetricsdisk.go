package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Edgemetricsdisk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricsdisk

	

	return json.Marshal(&struct { 
		AvailableBytes *float64 `json:"availableBytes,omitempty"`
		
		PartitionName *string `json:"partitionName,omitempty"`
		
		TotalBytes *float64 `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		AvailableBytes: u.AvailableBytes,
		
		PartitionName: u.PartitionName,
		
		TotalBytes: u.TotalBytes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricsdisk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
