package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagemetadatacontent - Metadata information about a message content.
type Conversationmessagemetadatacontent struct { 
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`


	// SubType - Content subtype
	SubType *string `json:"subType,omitempty"`

}

func (o *Conversationmessagemetadatacontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagemetadatacontent
	
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

func (o *Conversationmessagemetadatacontent) UnmarshalJSON(b []byte) error {
	var ConversationmessagemetadatacontentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagemetadatacontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := ConversationmessagemetadatacontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if SubType, ok := ConversationmessagemetadatacontentMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadatacontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
