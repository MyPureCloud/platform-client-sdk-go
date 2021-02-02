package platformclientv2
import (
	"encoding/json"
)

// Shorttermforecastingsettings - Short Term Forecasting Settings
type Shorttermforecastingsettings struct { 
	// DefaultHistoryWeeks - The number of weeks to consider by default when generating a volume forecast
	DefaultHistoryWeeks *int `json:"defaultHistoryWeeks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastingsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
