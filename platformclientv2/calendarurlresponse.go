package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Calendarurlresponse
type Calendarurlresponse struct { 
	// CalendarUrl - The calendar url for the user to subscribe with supported clients
	CalendarUrl *string `json:"calendarUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Calendarurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
