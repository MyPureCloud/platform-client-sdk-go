package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Edgemetricstopicedgemetricmemory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricmemory
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		AvailableBytes *int `json:"availableBytes,omitempty"`
		
		TotalBytes *int `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		AvailableBytes: o.AvailableBytes,
		
		TotalBytes: o.TotalBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricmemory) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricmemoryMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricmemoryMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EdgemetricstopicedgemetricmemoryMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if AvailableBytes, ok := EdgemetricstopicedgemetricmemoryMap["availableBytes"].(float64); ok {
		AvailableBytesInt := int(AvailableBytes)
		o.AvailableBytes = &AvailableBytesInt
	}
	
	if TotalBytes, ok := EdgemetricstopicedgemetricmemoryMap["totalBytes"].(float64); ok {
		TotalBytesInt := int(TotalBytes)
		o.TotalBytes = &TotalBytesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricmemory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
