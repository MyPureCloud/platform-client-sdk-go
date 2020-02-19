package platformclientv2
import (
	"encoding/json"
)

// Forecastsourcedaypointer - Pointer to look up source data for a short term forecast
type Forecastsourcedaypointer struct { 
	// DayOfWeek - The forecast day of week for this source data
	DayOfWeek *string `json:"dayOfWeek,omitempty"`


	// Weight - The relative weight to apply to this source data item for weighted averages
	Weight *int32 `json:"weight,omitempty"`


	// Date - The date this source data represents, in yyyy-MM-dd format
	Date *string `json:"date,omitempty"`


	// FileName - The name of the source file this data came from if it originated from a data import
	FileName *string `json:"fileName,omitempty"`


	// DataKey - The key to look up the forecast source data for this source day
	DataKey *string `json:"dataKey,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
