package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestdocumentsuggestionresult
type Knowledgeguestdocumentsuggestionresult struct { 
	// MatchedPhrase - Matched phrase to the autocomplete suggestions query.
	MatchedPhrase *string `json:"matchedPhrase,omitempty"`

}

func (o *Knowledgeguestdocumentsuggestionresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestdocumentsuggestionresult
	
	return json.Marshal(&struct { 
		MatchedPhrase *string `json:"matchedPhrase,omitempty"`
		*Alias
	}{ 
		MatchedPhrase: o.MatchedPhrase,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestdocumentsuggestionresult) UnmarshalJSON(b []byte) error {
	var KnowledgeguestdocumentsuggestionresultMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestdocumentsuggestionresultMap)
	if err != nil {
		return err
	}
	
	if MatchedPhrase, ok := KnowledgeguestdocumentsuggestionresultMap["matchedPhrase"].(string); ok {
		o.MatchedPhrase = &MatchedPhrase
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
