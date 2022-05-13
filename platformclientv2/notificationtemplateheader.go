package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationtemplateheader - Template header object.
type Notificationtemplateheader struct { 
	// VarType - Template header type.
	VarType *string `json:"type,omitempty"`


	// Text - Header text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`


	// Media - Media template header image.
	Media *Contentattachment `json:"media,omitempty"`


	// Parameters - Template parameters for placeholders in template.
	Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`

}

func (o *Notificationtemplateheader) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Notificationtemplateheader
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Media *Contentattachment `json:"media,omitempty"`
		
		Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Media: o.Media,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Notificationtemplateheader) UnmarshalJSON(b []byte) error {
	var NotificationtemplateheaderMap map[string]interface{}
	err := json.Unmarshal(b, &NotificationtemplateheaderMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := NotificationtemplateheaderMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := NotificationtemplateheaderMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Media, ok := NotificationtemplateheaderMap["media"].(map[string]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Parameters, ok := NotificationtemplateheaderMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Notificationtemplateheader) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
