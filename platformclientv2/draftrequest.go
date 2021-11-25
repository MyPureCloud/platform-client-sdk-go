package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Draftrequest
type Draftrequest struct { 
	// Intents - Draft intent object.
	Intents *[]Draftintents `json:"intents,omitempty"`


	// Topic
	Topic *[]Drafttopics `json:"topic,omitempty"`

}

func (o *Draftrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Draftrequest
	
	return json.Marshal(&struct { 
		Intents *[]Draftintents `json:"intents,omitempty"`
		
		Topic *[]Drafttopics `json:"topic,omitempty"`
		*Alias
	}{ 
		Intents: o.Intents,
		
		Topic: o.Topic,
		Alias:    (*Alias)(o),
	})
}

func (o *Draftrequest) UnmarshalJSON(b []byte) error {
	var DraftrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DraftrequestMap)
	if err != nil {
		return err
	}
	
	if Intents, ok := DraftrequestMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if Topic, ok := DraftrequestMap["topic"].([]interface{}); ok {
		TopicString, _ := json.Marshal(Topic)
		json.Unmarshal(TopicString, &o.Topic)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Draftrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
