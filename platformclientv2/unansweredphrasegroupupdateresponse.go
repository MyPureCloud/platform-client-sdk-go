package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredphrasegroupupdateresponse
type Unansweredphrasegroupupdateresponse struct { 
	// PhraseAssociations - List of phrases and documents linked in the patch request
	PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`


	// Group - Knowledge base unanswered group response
	Group *Unansweredgroup `json:"group,omitempty"`

}

func (o *Unansweredphrasegroupupdateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredphrasegroupupdateresponse
	
	return json.Marshal(&struct { 
		PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`
		
		Group *Unansweredgroup `json:"group,omitempty"`
		*Alias
	}{ 
		PhraseAssociations: o.PhraseAssociations,
		
		Group: o.Group,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredphrasegroupupdateresponse) UnmarshalJSON(b []byte) error {
	var UnansweredphrasegroupupdateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredphrasegroupupdateresponseMap)
	if err != nil {
		return err
	}
	
	if PhraseAssociations, ok := UnansweredphrasegroupupdateresponseMap["phraseAssociations"].([]interface{}); ok {
		PhraseAssociationsString, _ := json.Marshal(PhraseAssociations)
		json.Unmarshal(PhraseAssociationsString, &o.PhraseAssociations)
	}
	
	if Group, ok := UnansweredphrasegroupupdateresponseMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredphrasegroupupdateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
