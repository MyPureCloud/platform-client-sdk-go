package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Edgemetricstopicedgemetricdisk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricdisk

	

	return json.Marshal(&struct { 
		PartitionName *string `json:"partitionName,omitempty"`
		
		AvailableBytes *int `json:"availableBytes,omitempty"`
		
		TotalBytes *int `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		PartitionName: u.PartitionName,
		
		AvailableBytes: u.AvailableBytes,
		
		TotalBytes: u.TotalBytes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricdisk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
