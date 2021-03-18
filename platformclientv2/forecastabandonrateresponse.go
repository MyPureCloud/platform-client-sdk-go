package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastabandonrateresponse
type Forecastabandonrateresponse struct { 
	// Percent - The target percent abandon rate goal
	Percent *int `json:"percent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastabandonrateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
