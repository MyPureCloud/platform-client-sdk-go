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

func (o *Edgemetricsprocessor) MarshalJSON() ([]byte, error) {
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
		ActiveTimePct: o.ActiveTimePct,
		
		CpuId: o.CpuId,
		
		IdleTimePct: o.IdleTimePct,
		
		PrivilegedTimePct: o.PrivilegedTimePct,
		
		UserTimePct: o.UserTimePct,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricsprocessor) UnmarshalJSON(b []byte) error {
	var EdgemetricsprocessorMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsprocessorMap)
	if err != nil {
		return err
	}
	
	if ActiveTimePct, ok := EdgemetricsprocessorMap["activeTimePct"].(float64); ok {
		o.ActiveTimePct = &ActiveTimePct
	}
	
	if CpuId, ok := EdgemetricsprocessorMap["cpuId"].(string); ok {
		o.CpuId = &CpuId
	}
	
	if IdleTimePct, ok := EdgemetricsprocessorMap["idleTimePct"].(float64); ok {
		o.IdleTimePct = &IdleTimePct
	}
	
	if PrivilegedTimePct, ok := EdgemetricsprocessorMap["privilegedTimePct"].(float64); ok {
		o.PrivilegedTimePct = &PrivilegedTimePct
	}
	
	if UserTimePct, ok := EdgemetricsprocessorMap["userTimePct"].(float64); ok {
		o.UserTimePct = &UserTimePct
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricsprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
