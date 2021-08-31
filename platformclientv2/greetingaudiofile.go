package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Greetingaudiofile
type Greetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Greetingaudiofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greetingaudiofile
	
	return json.Marshal(&struct { 
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		SizeBytes *int `json:"sizeBytes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		DurationMilliseconds: o.DurationMilliseconds,
		
		SizeBytes: o.SizeBytes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Greetingaudiofile) UnmarshalJSON(b []byte) error {
	var GreetingaudiofileMap map[string]interface{}
	err := json.Unmarshal(b, &GreetingaudiofileMap)
	if err != nil {
		return err
	}
	
	if DurationMilliseconds, ok := GreetingaudiofileMap["durationMilliseconds"].(float64); ok {
		DurationMillisecondsInt := int(DurationMilliseconds)
		o.DurationMilliseconds = &DurationMillisecondsInt
	}
	
	if SizeBytes, ok := GreetingaudiofileMap["sizeBytes"].(float64); ok {
		SizeBytesInt := int(SizeBytes)
		o.SizeBytes = &SizeBytesInt
	}
	
	if SelfUri, ok := GreetingaudiofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Greetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
