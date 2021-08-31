package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotturnreference - A reference to a bot flow turn.
type Textbotturnreference struct { 
	// Id - The id of the turn.
	Id *string `json:"id,omitempty"`

}

func (o *Textbotturnreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotturnreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotturnreference) UnmarshalJSON(b []byte) error {
	var TextbotturnreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotturnreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotturnreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotturnreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
