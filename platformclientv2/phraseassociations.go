package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phraseassociations
type Phraseassociations struct { 
	// PhraseId - Id of the phrase to be linked
	PhraseId *string `json:"phraseId,omitempty"`


	// DocumentId - Id of the document to be linked
	DocumentId *string `json:"documentId,omitempty"`

}

func (o *Phraseassociations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phraseassociations
	
	return json.Marshal(&struct { 
		PhraseId *string `json:"phraseId,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		*Alias
	}{ 
		PhraseId: o.PhraseId,
		
		DocumentId: o.DocumentId,
		Alias:    (*Alias)(o),
	})
}

func (o *Phraseassociations) UnmarshalJSON(b []byte) error {
	var PhraseassociationsMap map[string]interface{}
	err := json.Unmarshal(b, &PhraseassociationsMap)
	if err != nil {
		return err
	}
	
	if PhraseId, ok := PhraseassociationsMap["phraseId"].(string); ok {
		o.PhraseId = &PhraseId
	}
    
	if DocumentId, ok := PhraseassociationsMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phraseassociations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
