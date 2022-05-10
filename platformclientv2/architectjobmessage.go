package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectjobmessage
type Architectjobmessage struct { 
	// DateTime - The DateTime when the message was generated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTime *time.Time `json:"dateTime,omitempty"`


	// VarType - The message type.
	VarType *string `json:"type,omitempty"`


	// Text - The text of the message.
	Text *string `json:"text,omitempty"`

}

func (o *Architectjobmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectjobmessage
	
	DateTime := new(string)
	if o.DateTime != nil {
		
		*DateTime = timeutil.Strftime(o.DateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTime = nil
	}
	
	return json.Marshal(&struct { 
		DateTime *string `json:"dateTime,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		DateTime: DateTime,
		
		VarType: o.VarType,
		
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectjobmessage) UnmarshalJSON(b []byte) error {
	var ArchitectjobmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectjobmessageMap)
	if err != nil {
		return err
	}
	
	if dateTimeString, ok := ArchitectjobmessageMap["dateTime"].(string); ok {
		DateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateTimeString)
		o.DateTime = &DateTime
	}
	
	if VarType, ok := ArchitectjobmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := ArchitectjobmessageMap["text"].(string); ok {
		o.Text = &Text
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectjobmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
