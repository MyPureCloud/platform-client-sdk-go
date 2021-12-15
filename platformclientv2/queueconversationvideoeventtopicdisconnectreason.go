package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicdisconnectreason
type Queueconversationvideoeventtopicdisconnectreason struct { 
	// VarType - Disconnect reason protocol type.
	VarType *string `json:"type,omitempty"`


	// Code - Protocol specific reason code. See the Q.850 and SIP specs.
	Code *int `json:"code,omitempty"`


	// Phrase - Human readable English description of the disconnect reason.
	Phrase *string `json:"phrase,omitempty"`

}

func (o *Queueconversationvideoeventtopicdisconnectreason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicdisconnectreason
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Code *int `json:"code,omitempty"`
		
		Phrase *string `json:"phrase,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Code: o.Code,
		
		Phrase: o.Phrase,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicdisconnectreason) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicdisconnectreasonMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicdisconnectreasonMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := QueueconversationvideoeventtopicdisconnectreasonMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Code, ok := QueueconversationvideoeventtopicdisconnectreasonMap["code"].(float64); ok {
		CodeInt := int(Code)
		o.Code = &CodeInt
	}
	
	if Phrase, ok := QueueconversationvideoeventtopicdisconnectreasonMap["phrase"].(string); ok {
		o.Phrase = &Phrase
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
