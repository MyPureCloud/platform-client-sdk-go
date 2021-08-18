package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatebusinessunitsettings
type Updatebusinessunitsettings struct { 
	// StartDayOfWeek - The start day of week for this business unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for this business unit, using the Olsen tz database format
	TimeZone *string `json:"timeZone,omitempty"`


	// ShortTermForecasting - Short term forecasting settings
	ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// Metadata - Version metadata for this business unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Updatebusinessunitsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatebusinessunitsettings

	

	return json.Marshal(&struct { 
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		StartDayOfWeek: u.StartDayOfWeek,
		
		TimeZone: u.TimeZone,
		
		ShortTermForecasting: u.ShortTermForecasting,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updatebusinessunitsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
