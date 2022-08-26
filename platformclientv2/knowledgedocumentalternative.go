package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentalternative
type Knowledgedocumentalternative struct { 
	// Phrase - Alternate phrasing to the document title.
	Phrase *string `json:"phrase,omitempty"`


	// Autocomplete - Autocomplete enabled for the alternate phrase.
	Autocomplete *bool `json:"autocomplete,omitempty"`

}

func (o *Knowledgedocumentalternative) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentalternative
	
	return json.Marshal(&struct { 
		Phrase *string `json:"phrase,omitempty"`
		
		Autocomplete *bool `json:"autocomplete,omitempty"`
		*Alias
	}{ 
		Phrase: o.Phrase,
		
		Autocomplete: o.Autocomplete,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentalternative) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentalternativeMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentalternativeMap)
	if err != nil {
		return err
	}
	
	if Phrase, ok := KnowledgedocumentalternativeMap["phrase"].(string); ok {
		o.Phrase = &Phrase
	}
    
	if Autocomplete, ok := KnowledgedocumentalternativeMap["autocomplete"].(bool); ok {
		o.Autocomplete = &Autocomplete
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
