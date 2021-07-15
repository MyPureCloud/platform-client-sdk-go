package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nuancemixdlgsettings
type Nuancemixdlgsettings struct { 
	// ChannelId - The Nuance channel ID to use when launching the Nuance bot, which must one of the code names of the bot's registered input channels.
	ChannelId *string `json:"channelId,omitempty"`


	// InputParameters - Name/value pairs of input variables to be sent to the Nuance bot. The values must be in the appropriate format for the variable's type (see https://docs.mix.nuance.com/dialog-grpc/v1/#simple-variable-types for help)
	InputParameters *map[string]interface{} `json:"inputParameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nuancemixdlgsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
