package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsuggestionresult
type Knowledgedocumentsuggestionresult struct { 
	// MatchedPhrase - Matched phrase to the autocomplete suggestions query.
	MatchedPhrase *string `json:"matchedPhrase,omitempty"`


	// Document
	Document *Knowledgedocumentsuggestionresultdocument `json:"document,omitempty"`

}

func (o *Knowledgedocumentsuggestionresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsuggestionresult
	
	return json.Marshal(&struct { 
		MatchedPhrase *string `json:"matchedPhrase,omitempty"`
		
		Document *Knowledgedocumentsuggestionresultdocument `json:"document,omitempty"`
		*Alias
	}{ 
		MatchedPhrase: o.MatchedPhrase,
		
		Document: o.Document,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsuggestionresult) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsuggestionresultMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsuggestionresultMap)
	if err != nil {
		return err
	}
	
	if MatchedPhrase, ok := KnowledgedocumentsuggestionresultMap["matchedPhrase"].(string); ok {
		o.MatchedPhrase = &MatchedPhrase
	}
    
	if Document, ok := KnowledgedocumentsuggestionresultMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
