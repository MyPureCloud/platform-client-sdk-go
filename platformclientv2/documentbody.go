package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbody
type Documentbody struct { 
	// Blocks - The list of building blocks for the document body.
	Blocks *[]Documentbodyblock `json:"blocks,omitempty"`

}

func (o *Documentbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbody
	
	return json.Marshal(&struct { 
		Blocks *[]Documentbodyblock `json:"blocks,omitempty"`
		*Alias
	}{ 
		Blocks: o.Blocks,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbody) UnmarshalJSON(b []byte) error {
	var DocumentbodyMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyMap)
	if err != nil {
		return err
	}
	
	if Blocks, ok := DocumentbodyMap["blocks"].([]interface{}); ok {
		BlocksString, _ := json.Marshal(Blocks)
		json.Unmarshal(BlocksString, &o.Blocks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
