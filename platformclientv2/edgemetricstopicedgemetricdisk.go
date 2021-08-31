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

func (o *Edgemetricstopicedgemetricdisk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricdisk
	
	return json.Marshal(&struct { 
		PartitionName *string `json:"partitionName,omitempty"`
		
		AvailableBytes *int `json:"availableBytes,omitempty"`
		
		TotalBytes *int `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		PartitionName: o.PartitionName,
		
		AvailableBytes: o.AvailableBytes,
		
		TotalBytes: o.TotalBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricdisk) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricdiskMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricdiskMap)
	if err != nil {
		return err
	}
	
	if PartitionName, ok := EdgemetricstopicedgemetricdiskMap["partitionName"].(string); ok {
		o.PartitionName = &PartitionName
	}
	
	if AvailableBytes, ok := EdgemetricstopicedgemetricdiskMap["availableBytes"].(float64); ok {
		AvailableBytesInt := int(AvailableBytes)
		o.AvailableBytes = &AvailableBytesInt
	}
	
	if TotalBytes, ok := EdgemetricstopicedgemetricdiskMap["totalBytes"].(float64); ok {
		TotalBytesInt := int(TotalBytes)
		o.TotalBytes = &TotalBytesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricdisk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
