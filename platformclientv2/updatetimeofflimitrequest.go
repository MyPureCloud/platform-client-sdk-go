package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatetimeofflimitrequest - Contains time off limit object property values to be updated.
type Updatetimeofflimitrequest struct { 
	// DefaultLimitMinutes - The default time off limit value in minutes per granularity
	DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`


	// Metadata - Version metadata for the time off limit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updatetimeofflimitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatetimeofflimitrequest
	
	return json.Marshal(&struct { 
		DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		DefaultLimitMinutes: o.DefaultLimitMinutes,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatetimeofflimitrequest) UnmarshalJSON(b []byte) error {
	var UpdatetimeofflimitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatetimeofflimitrequestMap)
	if err != nil {
		return err
	}
	
	if DefaultLimitMinutes, ok := UpdatetimeofflimitrequestMap["defaultLimitMinutes"].(float64); ok {
		DefaultLimitMinutesInt := int(DefaultLimitMinutes)
		o.DefaultLimitMinutes = &DefaultLimitMinutesInt
	}
	
	if Metadata, ok := UpdatetimeofflimitrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatetimeofflimitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
