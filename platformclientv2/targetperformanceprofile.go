package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Targetperformanceprofile
type Targetperformanceprofile struct { 
	// TargetPerformanceProfileId - The target destination performanceProfileId for the linked metric.
	TargetPerformanceProfileId *string `json:"targetPerformanceProfileId,omitempty"`

}

func (o *Targetperformanceprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Targetperformanceprofile
	
	return json.Marshal(&struct { 
		TargetPerformanceProfileId *string `json:"targetPerformanceProfileId,omitempty"`
		*Alias
	}{ 
		TargetPerformanceProfileId: o.TargetPerformanceProfileId,
		Alias:    (*Alias)(o),
	})
}

func (o *Targetperformanceprofile) UnmarshalJSON(b []byte) error {
	var TargetperformanceprofileMap map[string]interface{}
	err := json.Unmarshal(b, &TargetperformanceprofileMap)
	if err != nil {
		return err
	}
	
	if TargetPerformanceProfileId, ok := TargetperformanceprofileMap["targetPerformanceProfileId"].(string); ok {
		o.TargetPerformanceProfileId = &TargetPerformanceProfileId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Targetperformanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
