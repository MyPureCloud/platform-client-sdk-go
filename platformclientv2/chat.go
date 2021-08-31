package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chat
type Chat struct { 
	// JabberId
	JabberId *string `json:"jabberId,omitempty"`

}

func (o *Chat) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chat
	
	return json.Marshal(&struct { 
		JabberId *string `json:"jabberId,omitempty"`
		*Alias
	}{ 
		JabberId: o.JabberId,
		Alias:    (*Alias)(o),
	})
}

func (o *Chat) UnmarshalJSON(b []byte) error {
	var ChatMap map[string]interface{}
	err := json.Unmarshal(b, &ChatMap)
	if err != nil {
		return err
	}
	
	if JabberId, ok := ChatMap["jabberId"].(string); ok {
		o.JabberId = &JabberId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chat) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
