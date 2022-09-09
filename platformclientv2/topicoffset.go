package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topicoffset
type Topicoffset struct { 
	// WordCount - Count of words before the topic 
	WordCount *int `json:"wordCount,omitempty"`


	// CharacterCount - Count of characters before the topic 
	CharacterCount *int `json:"characterCount,omitempty"`

}

func (o *Topicoffset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicoffset
	
	return json.Marshal(&struct { 
		WordCount *int `json:"wordCount,omitempty"`
		
		CharacterCount *int `json:"characterCount,omitempty"`
		*Alias
	}{ 
		WordCount: o.WordCount,
		
		CharacterCount: o.CharacterCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Topicoffset) UnmarshalJSON(b []byte) error {
	var TopicoffsetMap map[string]interface{}
	err := json.Unmarshal(b, &TopicoffsetMap)
	if err != nil {
		return err
	}
	
	if WordCount, ok := TopicoffsetMap["wordCount"].(float64); ok {
		WordCountInt := int(WordCount)
		o.WordCount = &WordCountInt
	}
	
	if CharacterCount, ok := TopicoffsetMap["characterCount"].(float64); ok {
		CharacterCountInt := int(CharacterCount)
		o.CharacterCount = &CharacterCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topicoffset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
