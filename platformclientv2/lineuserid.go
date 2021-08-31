package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineuserid - Channel-specific User ID for Line accounts
type Lineuserid struct { 
	// UserId - The unique channel-specific userId for the user
	UserId *string `json:"userId,omitempty"`

}

func (o *Lineuserid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineuserid
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		Alias:    (*Alias)(o),
	})
}

func (o *Lineuserid) UnmarshalJSON(b []byte) error {
	var LineuseridMap map[string]interface{}
	err := json.Unmarshal(b, &LineuseridMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := LineuseridMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lineuserid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
