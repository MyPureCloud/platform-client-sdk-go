package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingsettings
type Recordingsettings struct { 
	// MaxSimultaneousStreams - Maximum number of simultaneous screen recording streams
	MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`


	// MaxConfigurableScreenRecordingStreams - Upper limit that maxSimultaneousStreams can be configured
	MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`


	// RegionalRecordingStorageEnabled - Store call recordings in the region where they are intended to be recorded, otherwise in the organization's home region
	RegionalRecordingStorageEnabled *bool `json:"regionalRecordingStorageEnabled,omitempty"`

}

func (o *Recordingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingsettings
	
	return json.Marshal(&struct { 
		MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`
		
		MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`
		
		RegionalRecordingStorageEnabled *bool `json:"regionalRecordingStorageEnabled,omitempty"`
		*Alias
	}{ 
		MaxSimultaneousStreams: o.MaxSimultaneousStreams,
		
		MaxConfigurableScreenRecordingStreams: o.MaxConfigurableScreenRecordingStreams,
		
		RegionalRecordingStorageEnabled: o.RegionalRecordingStorageEnabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingsettings) UnmarshalJSON(b []byte) error {
	var RecordingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingsettingsMap)
	if err != nil {
		return err
	}
	
	if MaxSimultaneousStreams, ok := RecordingsettingsMap["maxSimultaneousStreams"].(float64); ok {
		MaxSimultaneousStreamsInt := int(MaxSimultaneousStreams)
		o.MaxSimultaneousStreams = &MaxSimultaneousStreamsInt
	}
	
	if MaxConfigurableScreenRecordingStreams, ok := RecordingsettingsMap["maxConfigurableScreenRecordingStreams"].(float64); ok {
		MaxConfigurableScreenRecordingStreamsInt := int(MaxConfigurableScreenRecordingStreams)
		o.MaxConfigurableScreenRecordingStreams = &MaxConfigurableScreenRecordingStreamsInt
	}
	
	if RegionalRecordingStorageEnabled, ok := RecordingsettingsMap["regionalRecordingStorageEnabled"].(bool); ok {
		o.RegionalRecordingStorageEnabled = &RegionalRecordingStorageEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
