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

func (u *Edgemetricstopicedgemetricmemory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricmemory

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		AvailableBytes *int `json:"availableBytes,omitempty"`
		
		TotalBytes *int `json:"totalBytes,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		AvailableBytes: u.AvailableBytes,
		
		TotalBytes: u.TotalBytes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricmemory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
