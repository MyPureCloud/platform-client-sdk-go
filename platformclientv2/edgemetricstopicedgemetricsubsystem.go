package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricsubsystem
type Edgemetricstopicedgemetricsubsystem struct { 
	// ProcessName
	ProcessName *string `json:"processName,omitempty"`


	// DelayMs
	DelayMs *int `json:"delayMs,omitempty"`


	// MediaSubsystem
	MediaSubsystem *Edgemetricstopicedgemetricsubsystemmedia `json:"mediaSubsystem,omitempty"`

}

func (o *Edgemetricstopicedgemetricsubsystem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricsubsystem
	
	return json.Marshal(&struct { 
		ProcessName *string `json:"processName,omitempty"`
		
		DelayMs *int `json:"delayMs,omitempty"`
		
		MediaSubsystem *Edgemetricstopicedgemetricsubsystemmedia `json:"mediaSubsystem,omitempty"`
		*Alias
	}{ 
		ProcessName: o.ProcessName,
		
		DelayMs: o.DelayMs,
		
		MediaSubsystem: o.MediaSubsystem,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricsubsystem) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricsubsystemMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricsubsystemMap)
	if err != nil {
		return err
	}
	
	if ProcessName, ok := EdgemetricstopicedgemetricsubsystemMap["processName"].(string); ok {
		o.ProcessName = &ProcessName
	}
    
	if DelayMs, ok := EdgemetricstopicedgemetricsubsystemMap["delayMs"].(float64); ok {
		DelayMsInt := int(DelayMs)
		o.DelayMs = &DelayMsInt
	}
	
	if MediaSubsystem, ok := EdgemetricstopicedgemetricsubsystemMap["mediaSubsystem"].(map[string]interface{}); ok {
		MediaSubsystemString, _ := json.Marshal(MediaSubsystem)
		json.Unmarshal(MediaSubsystemString, &o.MediaSubsystem)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricsubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
