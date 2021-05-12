package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumntimezone
type Contactcolumntimezone struct { 
	// TimeZone - Time zone that the column matched to. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`


	// ColumnType - Column Type will be either PHONE or ZIP
	ColumnType *string `json:"columnType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactcolumntimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
