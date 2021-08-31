package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionresponse
type Nludetectionresponse struct { 
	// Version - The NLU domain version which performed the detection.
	Version *Nludomainversion `json:"version,omitempty"`


	// Output
	Output *Nludetectionoutput `json:"output,omitempty"`


	// Input
	Input *Nludetectioninput `json:"input,omitempty"`

}

func (o *Nludetectionresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectionresponse
	
	return json.Marshal(&struct { 
		Version *Nludomainversion `json:"version,omitempty"`
		
		Output *Nludetectionoutput `json:"output,omitempty"`
		
		Input *Nludetectioninput `json:"input,omitempty"`
		*Alias
	}{ 
		Version: o.Version,
		
		Output: o.Output,
		
		Input: o.Input,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludetectionresponse) UnmarshalJSON(b []byte) error {
	var NludetectionresponseMap map[string]interface{}
	err := json.Unmarshal(b, &NludetectionresponseMap)
	if err != nil {
		return err
	}
	
	if Version, ok := NludetectionresponseMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if Output, ok := NludetectionresponseMap["output"].(map[string]interface{}); ok {
		OutputString, _ := json.Marshal(Output)
		json.Unmarshal(OutputString, &o.Output)
	}
	
	if Input, ok := NludetectionresponseMap["input"].(map[string]interface{}); ok {
		InputString, _ := json.Marshal(Input)
		json.Unmarshal(InputString, &o.Input)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
