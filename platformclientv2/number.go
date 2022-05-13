package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Number
type Number struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`

}

func (o *Number) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Number
	
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

func (o *Number) UnmarshalJSON(b []byte) error {
	var NumberMap map[string]interface{}
	err := json.Unmarshal(b, &NumberMap)
	if err != nil {
		return err
	}
	
	if Start, ok := NumberMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := NumberMap["end"].(string); ok {
		o.End = &End
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Number) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
