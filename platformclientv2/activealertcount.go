package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activealertcount
type Activealertcount struct { 
	// Count - The count of active alerts for a user.
	Count *int `json:"count,omitempty"`

}

func (o *Activealertcount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activealertcount
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Activealertcount) UnmarshalJSON(b []byte) error {
	var ActivealertcountMap map[string]interface{}
	err := json.Unmarshal(b, &ActivealertcountMap)
	if err != nil {
		return err
	}
	
	if Count, ok := ActivealertcountMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Activealertcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
