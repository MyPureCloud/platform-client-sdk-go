package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsearchresult
type Knowledgedocumentsearchresult struct { 
	// Confidence - The confidence associated with a document with respect to a search query.
	Confidence *float64 `json:"confidence,omitempty"`


	// Document - Document that matched the query.
	Document *Knowledgedocumentresponse `json:"document,omitempty"`

}

func (o *Knowledgedocumentsearchresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsearchresult
	
	return json.Marshal(&struct { 
		Confidence *float64 `json:"confidence,omitempty"`
		
		Document *Knowledgedocumentresponse `json:"document,omitempty"`
		*Alias
	}{ 
		Confidence: o.Confidence,
		
		Document: o.Document,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsearchresult) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsearchresultMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsearchresultMap)
	if err != nil {
		return err
	}
	
	if Confidence, ok := KnowledgedocumentsearchresultMap["confidence"].(float64); ok {
		o.Confidence = &Confidence
	}
    
	if Document, ok := KnowledgedocumentsearchresultMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
