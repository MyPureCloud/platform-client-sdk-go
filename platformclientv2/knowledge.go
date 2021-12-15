package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledge
type Knowledge struct { 
	// Enabled - whether or not knowledge base is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// KnowledgeBase - The knowledge base for messenger
	KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`

}

func (o *Knowledge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledge
	
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

func (o *Knowledge) UnmarshalJSON(b []byte) error {
	var KnowledgeMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := KnowledgeMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if KnowledgeBase, ok := KnowledgeMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
