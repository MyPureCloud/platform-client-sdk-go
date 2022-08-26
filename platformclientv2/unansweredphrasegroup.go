package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredphrasegroup
type Unansweredphrasegroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Label - Knowledge base phrase group label
	Label *string `json:"label,omitempty"`


	// Phrases - List of unanswered phrases in a phrase group
	Phrases *[]Unansweredphrase `json:"phrases,omitempty"`


	// UnlinkedPhraseHitCount - Hit count of the unlinked phrase group
	UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`


	// UnlinkedPhraseCount - Unique phrase count of the unlinked phrase group
	UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Unansweredphrasegroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredphrasegroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		Phrases *[]Unansweredphrase `json:"phrases,omitempty"`
		
		UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`
		
		UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Label: o.Label,
		
		Phrases: o.Phrases,
		
		UnlinkedPhraseHitCount: o.UnlinkedPhraseHitCount,
		
		UnlinkedPhraseCount: o.UnlinkedPhraseCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredphrasegroup) UnmarshalJSON(b []byte) error {
	var UnansweredphrasegroupMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredphrasegroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UnansweredphrasegroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Label, ok := UnansweredphrasegroupMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if Phrases, ok := UnansweredphrasegroupMap["phrases"].([]interface{}); ok {
		PhrasesString, _ := json.Marshal(Phrases)
		json.Unmarshal(PhrasesString, &o.Phrases)
	}
	
	if UnlinkedPhraseHitCount, ok := UnansweredphrasegroupMap["unlinkedPhraseHitCount"].(float64); ok {
		UnlinkedPhraseHitCountInt := int(UnlinkedPhraseHitCount)
		o.UnlinkedPhraseHitCount = &UnlinkedPhraseHitCountInt
	}
	
	if UnlinkedPhraseCount, ok := UnansweredphrasegroupMap["unlinkedPhraseCount"].(float64); ok {
		UnlinkedPhraseCountInt := int(UnlinkedPhraseCount)
		o.UnlinkedPhraseCount = &UnlinkedPhraseCountInt
	}
	
	if SelfUri, ok := UnansweredphrasegroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredphrasegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
