package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitlength
type Digitlength struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`

}

func (o *Digitlength) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Digitlength
	
	return json.Marshal(&struct { 
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		*Alias
	}{ 
		Start: o.Start,
		
		End: o.End,
		Alias:    (*Alias)(o),
	})
}

func (o *Digitlength) UnmarshalJSON(b []byte) error {
	var DigitlengthMap map[string]interface{}
	err := json.Unmarshal(b, &DigitlengthMap)
	if err != nil {
		return err
	}
	
	if Start, ok := DigitlengthMap["start"].(string); ok {
		o.Start = &Start
	}
	
	if End, ok := DigitlengthMap["end"].(string); ok {
		o.End = &End
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Digitlength) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
