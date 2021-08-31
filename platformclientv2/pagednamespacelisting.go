package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Pagednamespacelisting
type Pagednamespacelisting struct { }

func (o *Pagednamespacelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Pagednamespacelisting
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Pagednamespacelisting) UnmarshalJSON(b []byte) error {
	var PagednamespacelistingMap map[string]interface{}
	err := json.Unmarshal(b, &PagednamespacelistingMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Pagednamespacelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
