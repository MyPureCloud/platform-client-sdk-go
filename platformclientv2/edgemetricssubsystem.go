package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricssubsystem
type Edgemetricssubsystem struct { 
	// DelayMs - Delay in milliseconds.
	DelayMs *int `json:"delayMs,omitempty"`


	// ProcessName - Name of the Edge process.
	ProcessName *string `json:"processName,omitempty"`


	// MediaSubsystem - Subsystem for an Edge device.
	MediaSubsystem **Edgemetricssubsystem `json:"mediaSubsystem,omitempty"`

}

func (o *Edgemetricssubsystem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricssubsystem
	
	return json.Marshal(&struct { 
		DelayMs *int `json:"delayMs,omitempty"`
		
		ProcessName *string `json:"processName,omitempty"`
		
		MediaSubsystem **Edgemetricssubsystem `json:"mediaSubsystem,omitempty"`
		*Alias
	}{ 
		DelayMs: o.DelayMs,
		
		ProcessName: o.ProcessName,
		
		MediaSubsystem: o.MediaSubsystem,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricssubsystem) UnmarshalJSON(b []byte) error {
	var EdgemetricssubsystemMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricssubsystemMap)
	if err != nil {
		return err
	}
	
	if DelayMs, ok := EdgemetricssubsystemMap["delayMs"].(float64); ok {
		DelayMsInt := int(DelayMs)
		o.DelayMs = &DelayMsInt
	}
	
	if ProcessName, ok := EdgemetricssubsystemMap["processName"].(string); ok {
		o.ProcessName = &ProcessName
	}
	
	if MediaSubsystem, ok := EdgemetricssubsystemMap["mediaSubsystem"].(map[string]interface{}); ok {
		MediaSubsystemString, _ := json.Marshal(MediaSubsystem)
		json.Unmarshal(MediaSubsystemString, &o.MediaSubsystem)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricssubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
