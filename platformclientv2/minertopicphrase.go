package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minertopicphrase
type Minertopicphrase struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Phrase name.
	Name *string `json:"name,omitempty"`


	// Topic - Topic associated with a phrase.
	Topic *Minertopic `json:"topic,omitempty"`


	// Utterances - List of utterances related to a phrase.
	Utterances *[]Utterance `json:"utterances,omitempty"`


	// UtteranceCount - Number of utterances belonging to a phrase
	UtteranceCount *int `json:"utteranceCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Minertopicphrase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minertopicphrase
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Topic *Minertopic `json:"topic,omitempty"`
		
		Utterances *[]Utterance `json:"utterances,omitempty"`
		
		UtteranceCount *int `json:"utteranceCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Topic: o.Topic,
		
		Utterances: o.Utterances,
		
		UtteranceCount: o.UtteranceCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Minertopicphrase) UnmarshalJSON(b []byte) error {
	var MinertopicphraseMap map[string]interface{}
	err := json.Unmarshal(b, &MinertopicphraseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MinertopicphraseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MinertopicphraseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Topic, ok := MinertopicphraseMap["topic"].(map[string]interface{}); ok {
		TopicString, _ := json.Marshal(Topic)
		json.Unmarshal(TopicString, &o.Topic)
	}
	
	if Utterances, ok := MinertopicphraseMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	
	if UtteranceCount, ok := MinertopicphraseMap["utteranceCount"].(float64); ok {
		UtteranceCountInt := int(UtteranceCount)
		o.UtteranceCount = &UtteranceCountInt
	}
	
	if SelfUri, ok := MinertopicphraseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Minertopicphrase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
