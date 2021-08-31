package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer
type Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer struct { 
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

func (o *Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer
	
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

func (o *Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap)
	if err != nil {
		return err
	}
	
	if DayOfWeek, ok := WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap["dayOfWeek"].(string); ok {
		o.DayOfWeek = &DayOfWeek
	}
	
	if Weight, ok := WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap["weight"].(float64); ok {
		WeightInt := int(Weight)
		o.Weight = &WeightInt
	}
	
	if Date, ok := WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap["date"].(string); ok {
		o.Date = &Date
	}
	
	if FileName, ok := WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
	
	if DataKey, ok := WfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointerMap["dataKey"].(string); ok {
		o.DataKey = &DataKey
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicforecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
