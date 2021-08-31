package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigthirdparty
type Widgetclientconfigthirdparty struct { }

func (o *Widgetclientconfigthirdparty) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfigthirdparty
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Widgetclientconfigthirdparty) UnmarshalJSON(b []byte) error {
	var WidgetclientconfigthirdpartyMap map[string]interface{}
	err := json.Unmarshal(b, &WidgetclientconfigthirdpartyMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetclientconfigthirdparty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
