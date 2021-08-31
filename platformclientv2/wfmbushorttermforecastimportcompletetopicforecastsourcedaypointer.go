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

func (o *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) MarshalJSON() ([]byte, error) {
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
		DayOfWeek: o.DayOfWeek,
		
		Weight: o.Weight,
		
		Date: o.Date,
		
		FileName: o.FileName,
		
		DataKey: o.DataKey,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap)
	if err != nil {
		return err
	}
	
	if DayOfWeek, ok := WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap["dayOfWeek"].(string); ok {
		o.DayOfWeek = &DayOfWeek
	}
	
	if Weight, ok := WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap["weight"].(float64); ok {
		WeightInt := int(Weight)
		o.Weight = &WeightInt
	}
	
	if Date, ok := WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap["date"].(string); ok {
		o.Date = &Date
	}
	
	if FileName, ok := WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
	
	if DataKey, ok := WfmbushorttermforecastimportcompletetopicforecastsourcedaypointerMap["dataKey"].(string); ok {
		o.DataKey = &DataKey
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
