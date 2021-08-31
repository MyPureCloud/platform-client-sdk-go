package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicprovisioninfo
type Phonechangetopicprovisioninfo struct { 
	// Time
	Time *time.Time `json:"time,omitempty"`


	// Source
	Source *string `json:"source,omitempty"`


	// ErrorInfo
	ErrorInfo *string `json:"errorInfo,omitempty"`

}

func (o *Phonechangetopicprovisioninfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicprovisioninfo
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		Time *string `json:"time,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		ErrorInfo *string `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Time: Time,
		
		Source: o.Source,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonechangetopicprovisioninfo) UnmarshalJSON(b []byte) error {
	var PhonechangetopicprovisioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopicprovisioninfoMap)
	if err != nil {
		return err
	}
	
	if timeString, ok := PhonechangetopicprovisioninfoMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if Source, ok := PhonechangetopicprovisioninfoMap["source"].(string); ok {
		o.Source = &Source
	}
	
	if ErrorInfo, ok := PhonechangetopicprovisioninfoMap["errorInfo"].(string); ok {
		o.ErrorInfo = &ErrorInfo
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopicprovisioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
