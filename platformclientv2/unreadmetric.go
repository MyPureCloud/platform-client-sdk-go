package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unreadmetric
type Unreadmetric struct { 
	// Count - The count of unread alerts for a specific rule type.
	Count *int `json:"count,omitempty"`

}

func (o *Unreadmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unreadmetric
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Unreadmetric) UnmarshalJSON(b []byte) error {
	var UnreadmetricMap map[string]interface{}
	err := json.Unmarshal(b, &UnreadmetricMap)
	if err != nil {
		return err
	}
	
	if Count, ok := UnreadmetricMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unreadmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
