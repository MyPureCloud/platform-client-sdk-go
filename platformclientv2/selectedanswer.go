package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Selectedanswer
type Selectedanswer struct { 
	// Document - The search result document chosen as the answer.
	Document *Addressableentityref `json:"document,omitempty"`

}

func (o *Selectedanswer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Selectedanswer
	
	return json.Marshal(&struct { 
		Document *Addressableentityref `json:"document,omitempty"`
		*Alias
	}{ 
		Document: o.Document,
		Alias:    (*Alias)(o),
	})
}

func (o *Selectedanswer) UnmarshalJSON(b []byte) error {
	var SelectedanswerMap map[string]interface{}
	err := json.Unmarshal(b, &SelectedanswerMap)
	if err != nil {
		return err
	}
	
	if Document, ok := SelectedanswerMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Selectedanswer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
