package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Markcontactaddressuncontactableactionsettings
type Markcontactaddressuncontactableactionsettings struct { }

func (o *Markcontactaddressuncontactableactionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Markcontactaddressuncontactableactionsettings
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Markcontactaddressuncontactableactionsettings) UnmarshalJSON(b []byte) error {
	var MarkcontactaddressuncontactableactionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &MarkcontactaddressuncontactableactionsettingsMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Markcontactaddressuncontactableactionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
