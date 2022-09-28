package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Viprouting
type Viprouting struct { 
	// VipCallMediaSettings - VIP Settings specific to Call media.
	VipCallMediaSettings *Vipcallmediasettings `json:"vipCallMediaSettings,omitempty"`


	// VipEmailMediaSettings - VIP Settings specific to Email media.
	VipEmailMediaSettings *Vipmediasettings `json:"vipEmailMediaSettings,omitempty"`


	// VipMessageMediaSettings - VIP Settings specific to Message media.
	VipMessageMediaSettings *Vipmediasettings `json:"vipMessageMediaSettings,omitempty"`


	// VipMaxOwnershipTimeSeconds - The max amount of time a VIP interaction will wait for the VIP user before going to the selected backup option (in seconds). Allowable range 10 seconds - 864000 seconds (ten days).
	VipMaxOwnershipTimeSeconds *int `json:"vipMaxOwnershipTimeSeconds,omitempty"`


	// VipBackup - VIP queue default VIP Backup settings for unanswered VIP interactions to fallback to. VIP Backup set for a specific VIP user has preference before queue default.
	VipBackup *Vipbackup `json:"vipBackup,omitempty"`

}

func (o *Viprouting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Viprouting
	
	return json.Marshal(&struct { 
		VipCallMediaSettings *Vipcallmediasettings `json:"vipCallMediaSettings,omitempty"`
		
		VipEmailMediaSettings *Vipmediasettings `json:"vipEmailMediaSettings,omitempty"`
		
		VipMessageMediaSettings *Vipmediasettings `json:"vipMessageMediaSettings,omitempty"`
		
		VipMaxOwnershipTimeSeconds *int `json:"vipMaxOwnershipTimeSeconds,omitempty"`
		
		VipBackup *Vipbackup `json:"vipBackup,omitempty"`
		*Alias
	}{ 
		VipCallMediaSettings: o.VipCallMediaSettings,
		
		VipEmailMediaSettings: o.VipEmailMediaSettings,
		
		VipMessageMediaSettings: o.VipMessageMediaSettings,
		
		VipMaxOwnershipTimeSeconds: o.VipMaxOwnershipTimeSeconds,
		
		VipBackup: o.VipBackup,
		Alias:    (*Alias)(o),
	})
}

func (o *Viprouting) UnmarshalJSON(b []byte) error {
	var ViproutingMap map[string]interface{}
	err := json.Unmarshal(b, &ViproutingMap)
	if err != nil {
		return err
	}
	
	if VipCallMediaSettings, ok := ViproutingMap["vipCallMediaSettings"].(map[string]interface{}); ok {
		VipCallMediaSettingsString, _ := json.Marshal(VipCallMediaSettings)
		json.Unmarshal(VipCallMediaSettingsString, &o.VipCallMediaSettings)
	}
	
	if VipEmailMediaSettings, ok := ViproutingMap["vipEmailMediaSettings"].(map[string]interface{}); ok {
		VipEmailMediaSettingsString, _ := json.Marshal(VipEmailMediaSettings)
		json.Unmarshal(VipEmailMediaSettingsString, &o.VipEmailMediaSettings)
	}
	
	if VipMessageMediaSettings, ok := ViproutingMap["vipMessageMediaSettings"].(map[string]interface{}); ok {
		VipMessageMediaSettingsString, _ := json.Marshal(VipMessageMediaSettings)
		json.Unmarshal(VipMessageMediaSettingsString, &o.VipMessageMediaSettings)
	}
	
	if VipMaxOwnershipTimeSeconds, ok := ViproutingMap["vipMaxOwnershipTimeSeconds"].(float64); ok {
		VipMaxOwnershipTimeSecondsInt := int(VipMaxOwnershipTimeSeconds)
		o.VipMaxOwnershipTimeSeconds = &VipMaxOwnershipTimeSecondsInt
	}
	
	if VipBackup, ok := ViproutingMap["vipBackup"].(map[string]interface{}); ok {
		VipBackupString, _ := json.Marshal(VipBackup)
		json.Unmarshal(VipBackupString, &o.VipBackup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Viprouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
