package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
