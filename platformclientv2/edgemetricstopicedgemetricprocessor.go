package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricprocessor
type Edgemetricstopicedgemetricprocessor struct { 
	// CpuId
	CpuId *string `json:"cpuId,omitempty"`


	// IdleTimePct
	IdleTimePct *int `json:"idleTimePct,omitempty"`


	// ActiveTimePct
	ActiveTimePct *int `json:"activeTimePct,omitempty"`


	// PrivilegedTimePct
	PrivilegedTimePct *int `json:"privilegedTimePct,omitempty"`


	// UserTimePct
	UserTimePct *int `json:"userTimePct,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
