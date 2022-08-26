package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredphrasegrouppatchrequestbody
type Unansweredphrasegrouppatchrequestbody struct { 
	// PhraseAssociations - List of phrases and documents to be linked
	PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`

}

func (o *Unansweredphrasegrouppatchrequestbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredphrasegrouppatchrequestbody
	
	return json.Marshal(&struct { 
		PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`
		*Alias
	}{ 
		PhraseAssociations: o.PhraseAssociations,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredphrasegrouppatchrequestbody) UnmarshalJSON(b []byte) error {
	var UnansweredphrasegrouppatchrequestbodyMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredphrasegrouppatchrequestbodyMap)
	if err != nil {
		return err
	}
	
	if PhraseAssociations, ok := UnansweredphrasegrouppatchrequestbodyMap["phraseAssociations"].([]interface{}); ok {
		PhraseAssociationsString, _ := json.Marshal(PhraseAssociations)
		json.Unmarshal(PhraseAssociationsString, &o.PhraseAssociations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredphrasegrouppatchrequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
