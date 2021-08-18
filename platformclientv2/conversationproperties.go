package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Conversationproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationproperties

	

	return json.Marshal(&struct { 
		IsWaiting *bool `json:"isWaiting,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		IsAcd *bool `json:"isAcd,omitempty"`
		
		IsPreferred *bool `json:"isPreferred,omitempty"`
		
		IsScreenshare *bool `json:"isScreenshare,omitempty"`
		
		IsCobrowse *bool `json:"isCobrowse,omitempty"`
		
		IsVoicemail *bool `json:"isVoicemail,omitempty"`
		
		IsFlagged *bool `json:"isFlagged,omitempty"`
		
		IsMonitored *bool `json:"isMonitored,omitempty"`
		
		FilterWrapUpNotes *bool `json:"filterWrapUpNotes,omitempty"`
		
		MatchAll *bool `json:"matchAll,omitempty"`
		*Alias
	}{ 
		IsWaiting: u.IsWaiting,
		
		IsActive: u.IsActive,
		
		IsAcd: u.IsAcd,
		
		IsPreferred: u.IsPreferred,
		
		IsScreenshare: u.IsScreenshare,
		
		IsCobrowse: u.IsCobrowse,
		
		IsVoicemail: u.IsVoicemail,
		
		IsFlagged: u.IsFlagged,
		
		IsMonitored: u.IsMonitored,
		
		FilterWrapUpNotes: u.FilterWrapUpNotes,
		
		MatchAll: u.MatchAll,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
