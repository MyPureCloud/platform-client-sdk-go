package platformclientv2
import (
	"encoding/json"
)

// Bushorttermforecastingsettings
type Bushorttermforecastingsettings struct { 
	// DefaultHistoryWeeks - The number of historical weeks to consider when creating a forecast. This setting is only used for legacy weighted average forecasts
	DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastingsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
