package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitsettings
type Businessunitsettings struct { 
	// StartDayOfWeek - The start day of week for this business unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for this business unit, using the Olsen tz database format
	TimeZone *string `json:"timeZone,omitempty"`


	// ShortTermForecasting - Short term forecasting settings
	ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// Metadata - Version metadata for this business unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Businessunitsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
