package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Edgemetricstopicedgemetricprocessor) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricprocessor

	

	return json.Marshal(&struct { 
		CpuId *string `json:"cpuId,omitempty"`
		
		IdleTimePct *int `json:"idleTimePct,omitempty"`
		
		ActiveTimePct *int `json:"activeTimePct,omitempty"`
		
		PrivilegedTimePct *int `json:"privilegedTimePct,omitempty"`
		
		UserTimePct *int `json:"userTimePct,omitempty"`
		*Alias
	}{ 
		CpuId: u.CpuId,
		
		IdleTimePct: u.IdleTimePct,
		
		ActiveTimePct: u.ActiveTimePct,
		
		PrivilegedTimePct: u.PrivilegedTimePct,
		
		UserTimePct: u.UserTimePct,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
