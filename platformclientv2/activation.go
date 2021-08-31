package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activation
type Activation struct { 
	// VarType - Type of activation.
	VarType *string `json:"type,omitempty"`


	// DelayInSeconds - Activation delay time amount.
	DelayInSeconds *int `json:"delayInSeconds,omitempty"`

}

func (o *Activation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activation
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		DelayInSeconds *int `json:"delayInSeconds,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		DelayInSeconds: o.DelayInSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Activation) UnmarshalJSON(b []byte) error {
	var ActivationMap map[string]interface{}
	err := json.Unmarshal(b, &ActivationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ActivationMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if DelayInSeconds, ok := ActivationMap["delayInSeconds"].(float64); ok {
		DelayInSecondsInt := int(DelayInSeconds)
		o.DelayInSeconds = &DelayInSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Activation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
