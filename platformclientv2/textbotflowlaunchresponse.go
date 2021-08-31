package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlaunchresponse - Information related to a successful launch of a bot flow. The ID will be used in subsequent turn requests of the bot flow.
type Textbotflowlaunchresponse struct { 
	// Id - The session ID of the bot flow, used to send to subsequent turn requests
	Id *string `json:"id,omitempty"`

}

func (o *Textbotflowlaunchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowlaunchresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowlaunchresponse) UnmarshalJSON(b []byte) error {
	var TextbotflowlaunchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowlaunchresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotflowlaunchresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
