package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer
type Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer struct { 
	// DayOfWeek
	DayOfWeek *string `json:"dayOfWeek,omitempty"`


	// Weight
	Weight *int `json:"weight,omitempty"`


	// Date
	Date *string `json:"date,omitempty"`


	// FileName
	FileName *string `json:"fileName,omitempty"`


	// DataKey
	DataKey *string `json:"dataKey,omitempty"`

}

func (u *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer

	

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
func (o *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
