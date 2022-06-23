package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicmessagemetadatacontent
type Conversationeventtopicmessagemetadatacontent struct { 
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`


	// SubType - Content subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Conversationeventtopicmessagemetadatacontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicmessagemetadatacontent
	
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

func (o *Conversationeventtopicmessagemetadatacontent) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicmessagemetadatacontentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicmessagemetadatacontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := ConversationeventtopicmessagemetadatacontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if SubType, ok := ConversationeventtopicmessagemetadatacontentMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagemetadatacontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
