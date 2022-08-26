package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodylistblock
type Documentbodylistblock struct { 
	// VarType - The type of the list block.
	VarType *string `json:"type,omitempty"`


	// Blocks - The list of items for an OrderedList or an UnorderedList.
	Blocks *[]Documentcontentblock `json:"blocks,omitempty"`

}

func (o *Documentbodylistblock) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodylistblock
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Blocks *[]Documentcontentblock `json:"blocks,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Blocks: o.Blocks,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodylistblock) UnmarshalJSON(b []byte) error {
	var DocumentbodylistblockMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodylistblockMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DocumentbodylistblockMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Blocks, ok := DocumentbodylistblockMap["blocks"].([]interface{}); ok {
		BlocksString, _ := json.Marshal(Blocks)
		json.Unmarshal(BlocksString, &o.Blocks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodylistblock) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
