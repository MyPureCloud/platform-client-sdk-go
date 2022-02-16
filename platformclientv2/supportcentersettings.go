package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
	// Enabled - Whether or not support center is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// KnowledgeBase - The knowledge base for support center
	KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`

}

func (o *Supportcentersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcentersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		KnowledgeBase: o.KnowledgeBase,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcentersettings) UnmarshalJSON(b []byte) error {
	var SupportcentersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SupportcentersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if KnowledgeBase, ok := SupportcentersettingsMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
