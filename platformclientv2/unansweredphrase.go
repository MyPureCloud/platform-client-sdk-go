package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredphrase
type Unansweredphrase struct { 
	// Id - Id of an unanswered phrase
	Id *string `json:"id,omitempty"`


	// Text - Phrase text of an unanswered phrase
	Text *string `json:"text,omitempty"`


	// UnlinkedPhraseHitCount - Hit count of an unlinked phrase
	UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`

}

func (o *Unansweredphrase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredphrase
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		UnlinkedPhraseHitCount: o.UnlinkedPhraseHitCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredphrase) UnmarshalJSON(b []byte) error {
	var UnansweredphraseMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredphraseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UnansweredphraseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := UnansweredphraseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if UnlinkedPhraseHitCount, ok := UnansweredphraseMap["unlinkedPhraseHitCount"].(float64); ok {
		UnlinkedPhraseHitCountInt := int(UnlinkedPhraseHitCount)
		o.UnlinkedPhraseHitCount = &UnlinkedPhraseHitCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredphrase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
