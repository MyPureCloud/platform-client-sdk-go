package platformclientv2
import (
	"encoding/json"
)

// Conversationproperties
type Conversationproperties struct { 
	// IsWaiting - Indicates filtering for waiting
	IsWaiting *bool `json:"isWaiting,omitempty"`


	// IsActive - Indicates filtering for active
	IsActive *bool `json:"isActive,omitempty"`


	// IsAcd - Indicates filtering for Acd
	IsAcd *bool `json:"isAcd,omitempty"`


	// IsPreferred - Indicates filtering for Preferred Agent Routing
	IsPreferred *bool `json:"isPreferred,omitempty"`


	// IsScreenshare - Indicates filtering for screenshare
	IsScreenshare *bool `json:"isScreenshare,omitempty"`


	// IsCobrowse - Indicates filtering for Cobrowse
	IsCobrowse *bool `json:"isCobrowse,omitempty"`


	// IsVoicemail - Indicates filtering for Voice mail
	IsVoicemail *bool `json:"isVoicemail,omitempty"`


	// IsFlagged - Indicates filtering for flagged
	IsFlagged *bool `json:"isFlagged,omitempty"`


	// IsMonitored - Indicates filtering for monitored
	IsMonitored *bool `json:"isMonitored,omitempty"`


	// FilterWrapUpNotes - Indicates filtering for WrapUpNotes
	FilterWrapUpNotes *bool `json:"filterWrapUpNotes,omitempty"`


	// MatchAll - Indicates comparison operation, TRUE indicates filters will use AND logic, FALSE indicates OR logic
	MatchAll *bool `json:"matchAll,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationproperties) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
