package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Void
type Void struct { }

func (o *Void) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Void
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Void) UnmarshalJSON(b []byte) error {
	var VoidMap map[string]interface{}
	err := json.Unmarshal(b, &VoidMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Void) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
