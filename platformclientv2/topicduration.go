package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topicduration
type Topicduration struct { 
	// TotalMilliseconds - The total duration of the topic phrase within the conversation
	TotalMilliseconds *int `json:"totalMilliseconds,omitempty"`

}

func (o *Topicduration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicduration
	
	return json.Marshal(&struct { 
		TotalMilliseconds *int `json:"totalMilliseconds,omitempty"`
		*Alias
	}{ 
		TotalMilliseconds: o.TotalMilliseconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Topicduration) UnmarshalJSON(b []byte) error {
	var TopicdurationMap map[string]interface{}
	err := json.Unmarshal(b, &TopicdurationMap)
	if err != nil {
		return err
	}
	
	if TotalMilliseconds, ok := TopicdurationMap["totalMilliseconds"].(float64); ok {
		TotalMillisecondsInt := int(TotalMilliseconds)
		o.TotalMilliseconds = &TotalMillisecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topicduration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
