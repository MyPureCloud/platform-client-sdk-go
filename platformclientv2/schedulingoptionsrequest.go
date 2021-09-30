package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingoptionsrequest
type Schedulingoptionsrequest struct { 
	// NoForecastOptions - Schedule generation options to apply if no forecast is supplied
	NoForecastOptions *Schedulingnoforecastoptionsrequest `json:"noForecastOptions,omitempty"`

}

func (o *Schedulingoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingoptionsrequest
	
	return json.Marshal(&struct { 
		NoForecastOptions *Schedulingnoforecastoptionsrequest `json:"noForecastOptions,omitempty"`
		*Alias
	}{ 
		NoForecastOptions: o.NoForecastOptions,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulingoptionsrequest) UnmarshalJSON(b []byte) error {
	var SchedulingoptionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingoptionsrequestMap)
	if err != nil {
		return err
	}
	
	if NoForecastOptions, ok := SchedulingoptionsrequestMap["noForecastOptions"].(map[string]interface{}); ok {
		NoForecastOptionsString, _ := json.Marshal(NoForecastOptions)
		json.Unmarshal(NoForecastOptionsString, &o.NoForecastOptions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
