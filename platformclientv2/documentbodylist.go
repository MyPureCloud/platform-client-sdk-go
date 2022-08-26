package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodylist
type Documentbodylist struct { 
	// Blocks - The list of items for an OrderedList or an UnorderedList.
	Blocks *[]Documentbodylistblock `json:"blocks,omitempty"`

}

func (o *Documentbodylist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodylist
	
	return json.Marshal(&struct { 
		Blocks *[]Documentbodylistblock `json:"blocks,omitempty"`
		*Alias
	}{ 
		Blocks: o.Blocks,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodylist) UnmarshalJSON(b []byte) error {
	var DocumentbodylistMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodylistMap)
	if err != nil {
		return err
	}
	
	if Blocks, ok := DocumentbodylistMap["blocks"].([]interface{}); ok {
		BlocksString, _ := json.Marshal(Blocks)
		json.Unmarshal(BlocksString, &o.Blocks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
