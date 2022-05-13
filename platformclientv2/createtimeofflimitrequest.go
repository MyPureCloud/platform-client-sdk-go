package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createtimeofflimitrequest - Contains property values of time off limit to be created
type Createtimeofflimitrequest struct { 
	// Granularity - Granularity choice for time off limit. If not specified, 'Daily' is assumed
	Granularity *string `json:"granularity,omitempty"`


	// DefaultLimitMinutes - The default limit value in minutes per granularity.If not specified, then 0 is assumed, which means there are no time off minutes available
	DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`

}

func (o *Createtimeofflimitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createtimeofflimitrequest
	
	return json.Marshal(&struct { 
		Granularity *string `json:"granularity,omitempty"`
		
		DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`
		*Alias
	}{ 
		Granularity: o.Granularity,
		
		DefaultLimitMinutes: o.DefaultLimitMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Createtimeofflimitrequest) UnmarshalJSON(b []byte) error {
	var CreatetimeofflimitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatetimeofflimitrequestMap)
	if err != nil {
		return err
	}
	
	if Granularity, ok := CreatetimeofflimitrequestMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if DefaultLimitMinutes, ok := CreatetimeofflimitrequestMap["defaultLimitMinutes"].(float64); ok {
		DefaultLimitMinutesInt := int(DefaultLimitMinutes)
		o.DefaultLimitMinutes = &DefaultLimitMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createtimeofflimitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
