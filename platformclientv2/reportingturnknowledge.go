package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledge
type Reportingturnknowledge struct { 
	// KnowledgeBaseId - The Knowledge Base ID that the captured knowledge data relates to.
	KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty"`


	// Feedback - The knowledge feedback data that was captured during this reporting turn.
	Feedback *Reportingturnknowledgefeedback `json:"feedback,omitempty"`


	// Search - The knowledge search data that was captured during this reporting turn.
	Search *Reportingturnknowledgesearch `json:"search,omitempty"`

}

func (o *Reportingturnknowledge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledge
	
	return json.Marshal(&struct { 
		KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty"`
		
		Feedback *Reportingturnknowledgefeedback `json:"feedback,omitempty"`
		
		Search *Reportingturnknowledgesearch `json:"search,omitempty"`
		*Alias
	}{ 
		KnowledgeBaseId: o.KnowledgeBaseId,
		
		Feedback: o.Feedback,
		
		Search: o.Search,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingturnknowledge) UnmarshalJSON(b []byte) error {
	var ReportingturnknowledgeMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnknowledgeMap)
	if err != nil {
		return err
	}
	
	if KnowledgeBaseId, ok := ReportingturnknowledgeMap["knowledgeBaseId"].(string); ok {
		o.KnowledgeBaseId = &KnowledgeBaseId
	}
    
	if Feedback, ok := ReportingturnknowledgeMap["feedback"].(map[string]interface{}); ok {
		FeedbackString, _ := json.Marshal(Feedback)
		json.Unmarshal(FeedbackString, &o.Feedback)
	}
	
	if Search, ok := ReportingturnknowledgeMap["search"].(map[string]interface{}); ok {
		SearchString, _ := json.Marshal(Search)
		json.Unmarshal(SearchString, &o.Search)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
