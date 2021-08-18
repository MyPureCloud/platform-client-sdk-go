package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricsprocessor
type Edgemetricsprocessor struct { 
	// ActiveTimePct - Percent time processor was active.
	ActiveTimePct *float64 `json:"activeTimePct,omitempty"`


	// CpuId - Machine CPU identifier. 'total' will always be included in the array and is the total of all CPU resources.
	CpuId *string `json:"cpuId,omitempty"`


	// IdleTimePct - Percent time processor was idle.
	IdleTimePct *float64 `json:"idleTimePct,omitempty"`


	// PrivilegedTimePct - Percent time processor spent in privileged mode.
	PrivilegedTimePct *float64 `json:"privilegedTimePct,omitempty"`


	// UserTimePct - Percent time processor spent in user mode.
	UserTimePct *float64 `json:"userTimePct,omitempty"`

}

func (u *Edgemetricsprocessor) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricsprocessor

	

	return json.Marshal(&struct { 
		ActiveTimePct *float64 `json:"activeTimePct,omitempty"`
		
		CpuId *string `json:"cpuId,omitempty"`
		
		IdleTimePct *float64 `json:"idleTimePct,omitempty"`
		
		PrivilegedTimePct *float64 `json:"privilegedTimePct,omitempty"`
		
		UserTimePct *float64 `json:"userTimePct,omitempty"`
		*Alias
	}{ 
		ActiveTimePct: u.ActiveTimePct,
		
		CpuId: u.CpuId,
		
		IdleTimePct: u.IdleTimePct,
		
		PrivilegedTimePct: u.PrivilegedTimePct,
		
		UserTimePct: u.UserTimePct,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricsprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
