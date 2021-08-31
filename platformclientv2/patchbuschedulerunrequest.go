package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchbuschedulerunrequest
type Patchbuschedulerunrequest struct { 
	// ReschedulingOptions - The rescheduling options to update
	ReschedulingOptions *Patchbureschedulingoptionsrequest `json:"reschedulingOptions,omitempty"`

}

func (o *Patchbuschedulerunrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchbuschedulerunrequest
	
	return json.Marshal(&struct { 
		ReschedulingOptions *Patchbureschedulingoptionsrequest `json:"reschedulingOptions,omitempty"`
		*Alias
	}{ 
		ReschedulingOptions: o.ReschedulingOptions,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchbuschedulerunrequest) UnmarshalJSON(b []byte) error {
	var PatchbuschedulerunrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchbuschedulerunrequestMap)
	if err != nil {
		return err
	}
	
	if ReschedulingOptions, ok := PatchbuschedulerunrequestMap["reschedulingOptions"].(map[string]interface{}); ok {
		ReschedulingOptionsString, _ := json.Marshal(ReschedulingOptions)
		json.Unmarshal(ReschedulingOptionsString, &o.ReschedulingOptions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchbuschedulerunrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
