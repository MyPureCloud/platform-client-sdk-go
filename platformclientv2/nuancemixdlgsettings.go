package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Nuancemixdlgsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nuancemixdlgsettings
	
	return json.Marshal(&struct { 
		ChannelId *string `json:"channelId,omitempty"`
		
		InputParameters *map[string]interface{} `json:"inputParameters,omitempty"`
		*Alias
	}{ 
		ChannelId: o.ChannelId,
		
		InputParameters: o.InputParameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Nuancemixdlgsettings) UnmarshalJSON(b []byte) error {
	var NuancemixdlgsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &NuancemixdlgsettingsMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := NuancemixdlgsettingsMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
	
	if InputParameters, ok := NuancemixdlgsettingsMap["inputParameters"].(map[string]interface{}); ok {
		InputParametersString, _ := json.Marshal(InputParameters)
		json.Unmarshal(InputParametersString, &o.InputParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nuancemixdlgsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
