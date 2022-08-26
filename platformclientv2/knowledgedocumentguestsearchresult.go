package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentguestsearchresult
type Knowledgedocumentguestsearchresult struct { 
	// Confidence - The confidence associated with a document with respect to a search query.
	Confidence *float64 `json:"confidence,omitempty"`


	// Document - Document that matched the query.
	Document *Knowledgeguestdocument `json:"document,omitempty"`

}

func (o *Knowledgedocumentguestsearchresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentguestsearchresult
	
	return json.Marshal(&struct { 
		Confidence *float64 `json:"confidence,omitempty"`
		
		Document *Knowledgeguestdocument `json:"document,omitempty"`
		*Alias
	}{ 
		Confidence: o.Confidence,
		
		Document: o.Document,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentguestsearchresult) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentguestsearchresultMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentguestsearchresultMap)
	if err != nil {
		return err
	}
	
	if Confidence, ok := KnowledgedocumentguestsearchresultMap["confidence"].(float64); ok {
		o.Confidence = &Confidence
	}
    
	if Document, ok := KnowledgedocumentguestsearchresultMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
