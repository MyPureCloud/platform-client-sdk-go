package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Empty
type Empty struct { }

func (o *Empty) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Empty
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Empty) UnmarshalJSON(b []byte) error {
	var EmptyMap map[string]interface{}
	err := json.Unmarshal(b, &EmptyMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Empty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
