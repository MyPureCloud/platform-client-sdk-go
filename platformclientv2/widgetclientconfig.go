package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfig
type Widgetclientconfig struct { 
	// V1
	V1 *Widgetclientconfigv1 `json:"v1,omitempty"`


	// V2
	V2 *Widgetclientconfigv2 `json:"v2,omitempty"`


	// V1Http
	V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`


	// ThirdParty
	ThirdParty *Widgetclientconfigthirdparty `json:"third-party,omitempty"`

}

func (o *Widgetclientconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfig
	
	return json.Marshal(&struct { 
		V1 *Widgetclientconfigv1 `json:"v1,omitempty"`
		
		V2 *Widgetclientconfigv2 `json:"v2,omitempty"`
		
		V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`
		
		ThirdParty *Widgetclientconfigthirdparty `json:"third-party,omitempty"`
		*Alias
	}{ 
		V1: o.V1,
		
		V2: o.V2,
		
		V1Http: o.V1Http,
		
		ThirdParty: o.ThirdParty,
		Alias:    (*Alias)(o),
	})
}

func (o *Widgetclientconfig) UnmarshalJSON(b []byte) error {
	var WidgetclientconfigMap map[string]interface{}
	err := json.Unmarshal(b, &WidgetclientconfigMap)
	if err != nil {
		return err
	}
	
	if V1, ok := WidgetclientconfigMap["v1"].(map[string]interface{}); ok {
		V1String, _ := json.Marshal(V1)
		json.Unmarshal(V1String, &o.V1)
	}
	
	if V2, ok := WidgetclientconfigMap["v2"].(map[string]interface{}); ok {
		V2String, _ := json.Marshal(V2)
		json.Unmarshal(V2String, &o.V2)
	}
	
	if V1Http, ok := WidgetclientconfigMap["v1-http"].(map[string]interface{}); ok {
		V1HttpString, _ := json.Marshal(V1Http)
		json.Unmarshal(V1HttpString, &o.V1Http)
	}
	
	if ThirdParty, ok := WidgetclientconfigMap["third-party"].(map[string]interface{}); ok {
		ThirdPartyString, _ := json.Marshal(ThirdParty)
		json.Unmarshal(ThirdPartyString, &o.ThirdParty)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetclientconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
