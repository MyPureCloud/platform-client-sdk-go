package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phrase
type Phrase struct { 
	// Text - The phrase text
	Text *string `json:"text,omitempty"`


	// Strictness - The phrase strictness, default value is null
	Strictness *string `json:"strictness,omitempty"`


	// Sentiment - The phrase sentiment, default value is Unspecified
	Sentiment *string `json:"sentiment,omitempty"`

}

func (o *Phrase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phrase
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		Sentiment *string `json:"sentiment,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Strictness: o.Strictness,
		
		Sentiment: o.Sentiment,
		Alias:    (*Alias)(o),
	})
}

func (o *Phrase) UnmarshalJSON(b []byte) error {
	var PhraseMap map[string]interface{}
	err := json.Unmarshal(b, &PhraseMap)
	if err != nil {
		return err
	}
	
	if Text, ok := PhraseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Strictness, ok := PhraseMap["strictness"].(string); ok {
		o.Strictness = &Strictness
	}
    
	if Sentiment, ok := PhraseMap["sentiment"].(string); ok {
		o.Sentiment = &Sentiment
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phrase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
