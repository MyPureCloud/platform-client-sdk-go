package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigv2
type Widgetclientconfigv2 struct { }

func (u *Widgetclientconfigv2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfigv2

	

	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Widgetclientconfigv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
