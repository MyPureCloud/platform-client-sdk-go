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

func (u *Reportingturnknowledge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledge

	

	return json.Marshal(&struct { 
		KnowledgeBaseId *string `json:"knowledgeBaseId,omitempty"`
		
		Feedback *Reportingturnknowledgefeedback `json:"feedback,omitempty"`
		
		Search *Reportingturnknowledgesearch `json:"search,omitempty"`
		*Alias
	}{ 
		KnowledgeBaseId: u.KnowledgeBaseId,
		
		Feedback: u.Feedback,
		
		Search: u.Search,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
