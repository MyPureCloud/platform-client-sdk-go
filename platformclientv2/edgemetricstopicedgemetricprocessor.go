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

func (o *Edgemetricstopicedgemetricprocessor) MarshalJSON() ([]byte, error) {
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
		CpuId: o.CpuId,
		
		IdleTimePct: o.IdleTimePct,
		
		ActiveTimePct: o.ActiveTimePct,
		
		PrivilegedTimePct: o.PrivilegedTimePct,
		
		UserTimePct: o.UserTimePct,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricprocessor) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricprocessorMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricprocessorMap)
	if err != nil {
		return err
	}
	
	if CpuId, ok := EdgemetricstopicedgemetricprocessorMap["cpuId"].(string); ok {
		o.CpuId = &CpuId
	}
	
	if IdleTimePct, ok := EdgemetricstopicedgemetricprocessorMap["idleTimePct"].(float64); ok {
		IdleTimePctInt := int(IdleTimePct)
		o.IdleTimePct = &IdleTimePctInt
	}
	
	if ActiveTimePct, ok := EdgemetricstopicedgemetricprocessorMap["activeTimePct"].(float64); ok {
		ActiveTimePctInt := int(ActiveTimePct)
		o.ActiveTimePct = &ActiveTimePctInt
	}
	
	if PrivilegedTimePct, ok := EdgemetricstopicedgemetricprocessorMap["privilegedTimePct"].(float64); ok {
		PrivilegedTimePctInt := int(PrivilegedTimePct)
		o.PrivilegedTimePct = &PrivilegedTimePctInt
	}
	
	if UserTimePct, ok := EdgemetricstopicedgemetricprocessorMap["userTimePct"].(float64); ok {
		UserTimePctInt := int(UserTimePct)
		o.UserTimePct = &UserTimePctInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
