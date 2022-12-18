package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topicphrase
type Topicphrase struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// UtteranceCount
	UtteranceCount *int `json:"utteranceCount,omitempty"`

}

func (o *Topicphrase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicphrase
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UtteranceCount *int `json:"utteranceCount,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		UtteranceCount: o.UtteranceCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Topicphrase) UnmarshalJSON(b []byte) error {
	var TopicphraseMap map[string]interface{}
	err := json.Unmarshal(b, &TopicphraseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TopicphraseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := TopicphraseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if UtteranceCount, ok := TopicphraseMap["utteranceCount"].(float64); ok {
		UtteranceCountInt := int(UtteranceCount)
		o.UtteranceCount = &UtteranceCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topicphrase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
