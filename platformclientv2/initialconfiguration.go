package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Initialconfiguration
type Initialconfiguration struct { 
	// AudioState - Indicates the initial audio state for the communication.
	AudioState *Audiostate `json:"audioState,omitempty"`


	// Alerting - Indicates that this communication's initial state is alerting. If false, the communication started in a connected state.
	Alerting *bool `json:"alerting,omitempty"`


	// Inbound - Indicates the direction of this communication with respect to the contact center. `true` means the communication is INBOUND. `false` means the communication is OUTBOUND.
	Inbound *bool `json:"inbound,omitempty"`


	// InvitedBy - The id of the communication (the \"peer\") that \"invited\" this communication, if this occurred.
	InvitedBy *string `json:"invitedBy,omitempty"`


	// RecordingActive - Indicates whether recording is active for this communication at creation.
	RecordingActive *bool `json:"recordingActive,omitempty"`


	// AdditionalInfo - Additional metadata about this session which should be recorded by the platform but which will not be indexed or searchable. Primarily for diagnostic value. Any information that needs to be accessible through other components like Analytics should be moved to dedicated fields.
	AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`

}

func (o *Initialconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Initialconfiguration
	
	return json.Marshal(&struct { 
		AudioState *Audiostate `json:"audioState,omitempty"`
		
		Alerting *bool `json:"alerting,omitempty"`
		
		Inbound *bool `json:"inbound,omitempty"`
		
		InvitedBy *string `json:"invitedBy,omitempty"`
		
		RecordingActive *bool `json:"recordingActive,omitempty"`
		
		AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
		*Alias
	}{ 
		AudioState: o.AudioState,
		
		Alerting: o.Alerting,
		
		Inbound: o.Inbound,
		
		InvitedBy: o.InvitedBy,
		
		RecordingActive: o.RecordingActive,
		
		AdditionalInfo: o.AdditionalInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Initialconfiguration) UnmarshalJSON(b []byte) error {
	var InitialconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &InitialconfigurationMap)
	if err != nil {
		return err
	}
	
	if AudioState, ok := InitialconfigurationMap["audioState"].(map[string]interface{}); ok {
		AudioStateString, _ := json.Marshal(AudioState)
		json.Unmarshal(AudioStateString, &o.AudioState)
	}
	
	if Alerting, ok := InitialconfigurationMap["alerting"].(bool); ok {
		o.Alerting = &Alerting
	}
    
	if Inbound, ok := InitialconfigurationMap["inbound"].(bool); ok {
		o.Inbound = &Inbound
	}
    
	if InvitedBy, ok := InitialconfigurationMap["invitedBy"].(string); ok {
		o.InvitedBy = &InvitedBy
	}
    
	if RecordingActive, ok := InitialconfigurationMap["recordingActive"].(bool); ok {
		o.RecordingActive = &RecordingActive
	}
    
	if AdditionalInfo, ok := InitialconfigurationMap["additionalInfo"].(map[string]interface{}); ok {
		AdditionalInfoString, _ := json.Marshal(AdditionalInfo)
		json.Unmarshal(AdditionalInfoString, &o.AdditionalInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Initialconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
