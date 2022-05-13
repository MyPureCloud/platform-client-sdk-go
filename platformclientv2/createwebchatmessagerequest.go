package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatmessagerequest
type Createwebchatmessagerequest struct { 
	// Body - The message body. Note that message bodies are limited to 4,000 characters.
	Body *string `json:"body,omitempty"`


	// BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
	BodyType *string `json:"bodyType,omitempty"`

}

func (o *Createwebchatmessagerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createwebchatmessagerequest
	
	return json.Marshal(&struct { 
		Body *string `json:"body,omitempty"`
		
		BodyType *string `json:"bodyType,omitempty"`
		*Alias
	}{ 
		Body: o.Body,
		
		BodyType: o.BodyType,
		Alias:    (*Alias)(o),
	})
}

func (o *Createwebchatmessagerequest) UnmarshalJSON(b []byte) error {
	var CreatewebchatmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatmessagerequestMap)
	if err != nil {
		return err
	}
	
	if Body, ok := CreatewebchatmessagerequestMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if BodyType, ok := CreatewebchatmessagerequestMap["bodyType"].(string); ok {
		o.BodyType = &BodyType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
