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

func (o *Updatebusinessunitsettings) MarshalJSON() ([]byte, error) {
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
		StartDayOfWeek: o.StartDayOfWeek,
		
		TimeZone: o.TimeZone,
		
		ShortTermForecasting: o.ShortTermForecasting,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatebusinessunitsettings) UnmarshalJSON(b []byte) error {
	var UpdatebusinessunitsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatebusinessunitsettingsMap)
	if err != nil {
		return err
	}
	
	if StartDayOfWeek, ok := UpdatebusinessunitsettingsMap["startDayOfWeek"].(string); ok {
		o.StartDayOfWeek = &StartDayOfWeek
	}
	
	if TimeZone, ok := UpdatebusinessunitsettingsMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if ShortTermForecasting, ok := UpdatebusinessunitsettingsMap["shortTermForecasting"].(map[string]interface{}); ok {
		ShortTermForecastingString, _ := json.Marshal(ShortTermForecasting)
		json.Unmarshal(ShortTermForecastingString, &o.ShortTermForecasting)
	}
	
	if Metadata, ok := UpdatebusinessunitsettingsMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatebusinessunitsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
