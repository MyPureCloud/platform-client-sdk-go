package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastaveragespeedofanswerresponse
type Forecastaveragespeedofanswerresponse struct { 
	// Seconds - the average speed of answer goal in seconds
	Seconds *int `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastaveragespeedofanswerresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
