package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata
type V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata struct { 
	// CustomAttributes
	CustomAttributes *map[string]string `json:"customAttributes,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata
	
	return json.Marshal(&struct { 
		CustomAttributes *map[string]string `json:"customAttributes,omitempty"`
		*Alias
	}{ 
		CustomAttributes: o.CustomAttributes,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadataMap)
	if err != nil {
		return err
	}
	
	if CustomAttributes, ok := V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadataMap["customAttributes"].(map[string]interface{}); ok {
		CustomAttributesString, _ := json.Marshal(CustomAttributes)
		json.Unmarshal(CustomAttributesString, &o.CustomAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessagingchannelmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
