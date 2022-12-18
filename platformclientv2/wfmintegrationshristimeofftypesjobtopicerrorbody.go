package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintegrationshristimeofftypesjobtopicerrorbody
type Wfmintegrationshristimeofftypesjobtopicerrorbody struct { 
	// Status
	Status *int `json:"status,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`

}

func (o *Wfmintegrationshristimeofftypesjobtopicerrorbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintegrationshristimeofftypesjobtopicerrorbody
	
	return json.Marshal(&struct { 
		Status *int `json:"status,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		MessageParams *map[string]string `json:"messageParams,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Code: o.Code,
		
		Message: o.Message,
		
		MessageParams: o.MessageParams,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintegrationshristimeofftypesjobtopicerrorbody) UnmarshalJSON(b []byte) error {
	var WfmintegrationshristimeofftypesjobtopicerrorbodyMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintegrationshristimeofftypesjobtopicerrorbodyMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmintegrationshristimeofftypesjobtopicerrorbodyMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Code, ok := WfmintegrationshristimeofftypesjobtopicerrorbodyMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Message, ok := WfmintegrationshristimeofftypesjobtopicerrorbodyMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if MessageParams, ok := WfmintegrationshristimeofftypesjobtopicerrorbodyMap["messageParams"].(map[string]interface{}); ok {
		MessageParamsString, _ := json.Marshal(MessageParams)
		json.Unmarshal(MessageParamsString, &o.MessageParams)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintegrationshristimeofftypesjobtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
