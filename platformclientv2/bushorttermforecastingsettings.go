package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bushorttermforecastingsettings
type Bushorttermforecastingsettings struct { 
	// DefaultHistoryWeeks - The number of historical weeks to consider when creating a forecast. This setting is only used for legacy weighted average forecasts
	DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`

}

func (o *Bushorttermforecastingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bushorttermforecastingsettings
	
	return json.Marshal(&struct { 
		DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`
		*Alias
	}{ 
		DefaultHistoryWeeks: o.DefaultHistoryWeeks,
		Alias:    (*Alias)(o),
	})
}

func (o *Bushorttermforecastingsettings) UnmarshalJSON(b []byte) error {
	var BushorttermforecastingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &BushorttermforecastingsettingsMap)
	if err != nil {
		return err
	}
	
	if DefaultHistoryWeeks, ok := BushorttermforecastingsettingsMap["defaultHistoryWeeks"].(float64); ok {
		DefaultHistoryWeeksInt := int(DefaultHistoryWeeks)
		o.DefaultHistoryWeeks = &DefaultHistoryWeeksInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bushorttermforecastingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
