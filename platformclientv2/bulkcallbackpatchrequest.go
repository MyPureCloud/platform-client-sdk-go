package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkcallbackpatchrequest
type Bulkcallbackpatchrequest struct { 
	// PatchCallbackRequests - The list of requests to update callbacks in bulk
	PatchCallbackRequests *[]Patchcallbackrequest `json:"patchCallbackRequests,omitempty"`

}

func (o *Bulkcallbackpatchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkcallbackpatchrequest
	
	return json.Marshal(&struct { 
		PatchCallbackRequests *[]Patchcallbackrequest `json:"patchCallbackRequests,omitempty"`
		*Alias
	}{ 
		PatchCallbackRequests: o.PatchCallbackRequests,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkcallbackpatchrequest) UnmarshalJSON(b []byte) error {
	var BulkcallbackpatchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkcallbackpatchrequestMap)
	if err != nil {
		return err
	}
	
	if PatchCallbackRequests, ok := BulkcallbackpatchrequestMap["patchCallbackRequests"].([]interface{}); ok {
		PatchCallbackRequestsString, _ := json.Marshal(PatchCallbackRequests)
		json.Unmarshal(PatchCallbackRequestsString, &o.PatchCallbackRequests)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkcallbackpatchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
