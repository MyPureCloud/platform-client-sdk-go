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

func (o *Edgemetricsdisk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricsdisk
	
	return json.Marshal(&struct { 
		AvailableBytes *float64 `json:"availableBytes,omitempty"`
		
		PartitionName *string `json:"partitionName,omitempty"`
		
		TotalBytes *float64 `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		AvailableBytes: o.AvailableBytes,
		
		PartitionName: o.PartitionName,
		
		TotalBytes: o.TotalBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricsdisk) UnmarshalJSON(b []byte) error {
	var EdgemetricsdiskMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsdiskMap)
	if err != nil {
		return err
	}
	
	if AvailableBytes, ok := EdgemetricsdiskMap["availableBytes"].(float64); ok {
		o.AvailableBytes = &AvailableBytes
	}
    
	if PartitionName, ok := EdgemetricsdiskMap["partitionName"].(string); ok {
		o.PartitionName = &PartitionName
	}
    
	if TotalBytes, ok := EdgemetricsdiskMap["totalBytes"].(float64); ok {
		o.TotalBytes = &TotalBytes
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricsdisk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
