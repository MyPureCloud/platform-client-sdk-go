package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyparagraph
type Documentbodyparagraph struct { 
	// Blocks - The list of blocks for the paragraph.
	Blocks *[]Documentcontentblock `json:"blocks,omitempty"`

}

func (o *Documentbodyparagraph) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodyparagraph
	
	return json.Marshal(&struct { 
		Blocks *[]Documentcontentblock `json:"blocks,omitempty"`
		*Alias
	}{ 
		Blocks: o.Blocks,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodyparagraph) UnmarshalJSON(b []byte) error {
	var DocumentbodyparagraphMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyparagraphMap)
	if err != nil {
		return err
	}
	
	if Blocks, ok := DocumentbodyparagraphMap["blocks"].([]interface{}); ok {
		BlocksString, _ := json.Marshal(Blocks)
		json.Unmarshal(BlocksString, &o.Blocks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyparagraph) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
