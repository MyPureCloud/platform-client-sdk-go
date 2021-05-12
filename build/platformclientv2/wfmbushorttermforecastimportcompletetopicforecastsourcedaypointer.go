package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
