package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Additionallanguagessynonyms
type Additionallanguagessynonyms struct { 
	// Synonyms - Synonyms for additional language
	Synonyms *[]string `json:"synonyms,omitempty"`

}

func (o *Additionallanguagessynonyms) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Additionallanguagessynonyms
	
	return json.Marshal(&struct { 
		Synonyms *[]string `json:"synonyms,omitempty"`
		*Alias
	}{ 
		Synonyms: o.Synonyms,
		Alias:    (*Alias)(o),
	})
}

func (o *Additionallanguagessynonyms) UnmarshalJSON(b []byte) error {
	var AdditionallanguagessynonymsMap map[string]interface{}
	err := json.Unmarshal(b, &AdditionallanguagessynonymsMap)
	if err != nil {
		return err
	}
	
	if Synonyms, ok := AdditionallanguagessynonymsMap["synonyms"].([]interface{}); ok {
		SynonymsString, _ := json.Marshal(Synonyms)
		json.Unmarshal(SynonymsString, &o.Synonyms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Additionallanguagessynonyms) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
