package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recallentry
type Recallentry struct { 
	// NbrAttempts
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`

}

func (o *Recallentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recallentry
	
	return json.Marshal(&struct { 
		NbrAttempts *int `json:"nbrAttempts,omitempty"`
		
		MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`
		*Alias
	}{ 
		NbrAttempts: o.NbrAttempts,
		
		MinutesBetweenAttempts: o.MinutesBetweenAttempts,
		Alias:    (*Alias)(o),
	})
}

func (o *Recallentry) UnmarshalJSON(b []byte) error {
	var RecallentryMap map[string]interface{}
	err := json.Unmarshal(b, &RecallentryMap)
	if err != nil {
		return err
	}
	
	if NbrAttempts, ok := RecallentryMap["nbrAttempts"].(float64); ok {
		NbrAttemptsInt := int(NbrAttempts)
		o.NbrAttempts = &NbrAttemptsInt
	}
	
	if MinutesBetweenAttempts, ok := RecallentryMap["minutesBetweenAttempts"].(float64); ok {
		MinutesBetweenAttemptsInt := int(MinutesBetweenAttempts)
		o.MinutesBetweenAttempts = &MinutesBetweenAttemptsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recallentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
