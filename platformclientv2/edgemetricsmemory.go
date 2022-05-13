package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Edgemetricsmemory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricsmemory
	
	return json.Marshal(&struct { 
		AvailableBytes *float64 `json:"availableBytes,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		TotalBytes *float64 `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		AvailableBytes: o.AvailableBytes,
		
		VarType: o.VarType,
		
		TotalBytes: o.TotalBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricsmemory) UnmarshalJSON(b []byte) error {
	var EdgemetricsmemoryMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsmemoryMap)
	if err != nil {
		return err
	}
	
	if AvailableBytes, ok := EdgemetricsmemoryMap["availableBytes"].(float64); ok {
		o.AvailableBytes = &AvailableBytes
	}
    
	if VarType, ok := EdgemetricsmemoryMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if TotalBytes, ok := EdgemetricsmemoryMap["totalBytes"].(float64); ok {
		o.TotalBytes = &TotalBytes
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricsmemory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
