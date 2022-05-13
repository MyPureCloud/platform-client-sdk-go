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

func (o *Conversationproperties) MarshalJSON() ([]byte, error) {
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
		IsWaiting: o.IsWaiting,
		
		IsActive: o.IsActive,
		
		IsAcd: o.IsAcd,
		
		IsPreferred: o.IsPreferred,
		
		IsScreenshare: o.IsScreenshare,
		
		IsCobrowse: o.IsCobrowse,
		
		IsVoicemail: o.IsVoicemail,
		
		IsFlagged: o.IsFlagged,
		
		IsMonitored: o.IsMonitored,
		
		FilterWrapUpNotes: o.FilterWrapUpNotes,
		
		MatchAll: o.MatchAll,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationproperties) UnmarshalJSON(b []byte) error {
	var ConversationpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationpropertiesMap)
	if err != nil {
		return err
	}
	
	if IsWaiting, ok := ConversationpropertiesMap["isWaiting"].(bool); ok {
		o.IsWaiting = &IsWaiting
	}
    
	if IsActive, ok := ConversationpropertiesMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if IsAcd, ok := ConversationpropertiesMap["isAcd"].(bool); ok {
		o.IsAcd = &IsAcd
	}
    
	if IsPreferred, ok := ConversationpropertiesMap["isPreferred"].(bool); ok {
		o.IsPreferred = &IsPreferred
	}
    
	if IsScreenshare, ok := ConversationpropertiesMap["isScreenshare"].(bool); ok {
		o.IsScreenshare = &IsScreenshare
	}
    
	if IsCobrowse, ok := ConversationpropertiesMap["isCobrowse"].(bool); ok {
		o.IsCobrowse = &IsCobrowse
	}
    
	if IsVoicemail, ok := ConversationpropertiesMap["isVoicemail"].(bool); ok {
		o.IsVoicemail = &IsVoicemail
	}
    
	if IsFlagged, ok := ConversationpropertiesMap["isFlagged"].(bool); ok {
		o.IsFlagged = &IsFlagged
	}
    
	if IsMonitored, ok := ConversationpropertiesMap["isMonitored"].(bool); ok {
		o.IsMonitored = &IsMonitored
	}
    
	if FilterWrapUpNotes, ok := ConversationpropertiesMap["filterWrapUpNotes"].(bool); ok {
		o.FilterWrapUpNotes = &FilterWrapUpNotes
	}
    
	if MatchAll, ok := ConversationpropertiesMap["matchAll"].(bool); ok {
		o.MatchAll = &MatchAll
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
