package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagingchannelmetadata - Information about the channel.
type Conversationmessagingchannelmetadata struct { }

func (o *Conversationmessagingchannelmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagingchannelmetadata
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagingchannelmetadata) UnmarshalJSON(b []byte) error {
	var ConversationmessagingchannelmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagingchannelmetadataMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannelmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
