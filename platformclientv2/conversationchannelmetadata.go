package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchannelmetadata - Information about the channel.
type Conversationchannelmetadata struct { }

func (o *Conversationchannelmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchannelmetadata
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Conversationchannelmetadata) UnmarshalJSON(b []byte) error {
	var ConversationchannelmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchannelmetadataMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchannelmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
