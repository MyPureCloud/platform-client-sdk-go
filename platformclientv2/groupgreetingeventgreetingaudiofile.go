package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupgreetingeventgreetingaudiofile
type Groupgreetingeventgreetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`

}

func (o *Groupgreetingeventgreetingaudiofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupgreetingeventgreetingaudiofile
	
	return json.Marshal(&struct { 
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		SizeBytes *int `json:"sizeBytes,omitempty"`
		*Alias
	}{ 
		DurationMilliseconds: o.DurationMilliseconds,
		
		SizeBytes: o.SizeBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Groupgreetingeventgreetingaudiofile) UnmarshalJSON(b []byte) error {
	var GroupgreetingeventgreetingaudiofileMap map[string]interface{}
	err := json.Unmarshal(b, &GroupgreetingeventgreetingaudiofileMap)
	if err != nil {
		return err
	}
	
	if DurationMilliseconds, ok := GroupgreetingeventgreetingaudiofileMap["durationMilliseconds"].(float64); ok {
		DurationMillisecondsInt := int(DurationMilliseconds)
		o.DurationMilliseconds = &DurationMillisecondsInt
	}
	
	if SizeBytes, ok := GroupgreetingeventgreetingaudiofileMap["sizeBytes"].(float64); ok {
		SizeBytesInt := int(SizeBytes)
		o.SizeBytes = &SizeBytesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
