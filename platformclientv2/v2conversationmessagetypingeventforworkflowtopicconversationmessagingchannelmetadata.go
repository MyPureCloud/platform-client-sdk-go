package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata
type V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata struct { 
	// CustomAttributes
	CustomAttributes *map[string]string `json:"customAttributes,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata
	
	return json.Marshal(&struct { 
		CustomAttributes *map[string]string `json:"customAttributes,omitempty"`
		*Alias
	}{ 
		CustomAttributes: o.CustomAttributes,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadataMap)
	if err != nil {
		return err
	}
	
	if CustomAttributes, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadataMap["customAttributes"].(map[string]interface{}); ok {
		CustomAttributesString, _ := json.Marshal(CustomAttributes)
		json.Unmarshal(CustomAttributesString, &o.CustomAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannelmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
