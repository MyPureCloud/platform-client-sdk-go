package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastsourcedaypointer - Pointer to look up source data for a short term forecast
type Forecastsourcedaypointer struct { 
	// DayOfWeek - The forecast day of week for this source data
	DayOfWeek *string `json:"dayOfWeek,omitempty"`


	// Weight - The relative weight to apply to this source data item for weighted averages
	Weight *int `json:"weight,omitempty"`


	// Date - The date this source data represents, in yyyy-MM-dd format
	Date *string `json:"date,omitempty"`


	// FileName - The name of the source file this data came from if it originated from a data import
	FileName *string `json:"fileName,omitempty"`


	// DataKey - The key to look up the forecast source data for this source day
	DataKey *string `json:"dataKey,omitempty"`

}

func (u *Forecastsourcedaypointer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastsourcedaypointer

	

	return json.Marshal(&struct { 
		DayOfWeek *string `json:"dayOfWeek,omitempty"`
		
		Weight *int `json:"weight,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		DataKey *string `json:"dataKey,omitempty"`
		*Alias
	}{ 
		DayOfWeek: u.DayOfWeek,
		
		Weight: u.Weight,
		
		Date: u.Date,
		
		FileName: u.FileName,
		
		DataKey: u.DataKey,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
