package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Directrouting
type Directrouting struct { 
	// CallMediaSettings - Direct Routing Settings specific to Call media.
	CallMediaSettings *Directroutingcallmediasettings `json:"callMediaSettings,omitempty"`


	// EmailMediaSettings - Direct Routing Settings specific to Email media.
	EmailMediaSettings *Directroutingmediasettings `json:"emailMediaSettings,omitempty"`


	// MessageMediaSettings - Direct Routing Settings specific to Message media.
	MessageMediaSettings *Directroutingmediasettings `json:"messageMediaSettings,omitempty"`


	// BackupQueueId - ID of queue to be used as the default backup if an agent does not have their Backup Settings configured.
	BackupQueueId *string `json:"backupQueueId,omitempty"`


	// WaitForAgent - Flag indicating if Direct Routing interactions should wait for Direct Routing agent or go immediately to selected backup.
	WaitForAgent *bool `json:"waitForAgent,omitempty"`


	// AgentWaitSeconds - Time (in seconds) that a Direct Routing interaction will wait for Direct Routing agent before going to selected backup. Valid range [60, 864000].
	AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`

}

func (o *Directrouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Directrouting
	
	return json.Marshal(&struct { 
		CallMediaSettings *Directroutingcallmediasettings `json:"callMediaSettings,omitempty"`
		
		EmailMediaSettings *Directroutingmediasettings `json:"emailMediaSettings,omitempty"`
		
		MessageMediaSettings *Directroutingmediasettings `json:"messageMediaSettings,omitempty"`
		
		BackupQueueId *string `json:"backupQueueId,omitempty"`
		
		WaitForAgent *bool `json:"waitForAgent,omitempty"`
		
		AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`
		*Alias
	}{ 
		CallMediaSettings: o.CallMediaSettings,
		
		EmailMediaSettings: o.EmailMediaSettings,
		
		MessageMediaSettings: o.MessageMediaSettings,
		
		BackupQueueId: o.BackupQueueId,
		
		WaitForAgent: o.WaitForAgent,
		
		AgentWaitSeconds: o.AgentWaitSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Directrouting) UnmarshalJSON(b []byte) error {
	var DirectroutingMap map[string]interface{}
	err := json.Unmarshal(b, &DirectroutingMap)
	if err != nil {
		return err
	}
	
	if CallMediaSettings, ok := DirectroutingMap["callMediaSettings"].(map[string]interface{}); ok {
		CallMediaSettingsString, _ := json.Marshal(CallMediaSettings)
		json.Unmarshal(CallMediaSettingsString, &o.CallMediaSettings)
	}
	
	if EmailMediaSettings, ok := DirectroutingMap["emailMediaSettings"].(map[string]interface{}); ok {
		EmailMediaSettingsString, _ := json.Marshal(EmailMediaSettings)
		json.Unmarshal(EmailMediaSettingsString, &o.EmailMediaSettings)
	}
	
	if MessageMediaSettings, ok := DirectroutingMap["messageMediaSettings"].(map[string]interface{}); ok {
		MessageMediaSettingsString, _ := json.Marshal(MessageMediaSettings)
		json.Unmarshal(MessageMediaSettingsString, &o.MessageMediaSettings)
	}
	
	if BackupQueueId, ok := DirectroutingMap["backupQueueId"].(string); ok {
		o.BackupQueueId = &BackupQueueId
	}
    
	if WaitForAgent, ok := DirectroutingMap["waitForAgent"].(bool); ok {
		o.WaitForAgent = &WaitForAgent
	}
    
	if AgentWaitSeconds, ok := DirectroutingMap["agentWaitSeconds"].(float64); ok {
		AgentWaitSecondsInt := int(AgentWaitSeconds)
		o.AgentWaitSeconds = &AgentWaitSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Directrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
