package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagemetadatacontent
type Conversationmessageeventtopicmessagemetadatacontent struct { 
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`


	// SubType - Content subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Conversationmessageeventtopicmessagemetadatacontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagemetadatacontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		SubType *string `json:"subType,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		SubType: o.SubType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicmessagemetadatacontent) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagemetadatacontentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagemetadatacontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := ConversationmessageeventtopicmessagemetadatacontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if SubType, ok := ConversationmessageeventtopicmessagemetadatacontentMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagemetadatacontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
