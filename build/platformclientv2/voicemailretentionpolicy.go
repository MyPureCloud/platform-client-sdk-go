package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailretentionpolicy - Governs how the voicemail is retained
type Voicemailretentionpolicy struct { 
	// VoicemailRetentionPolicyType - The retention policy type
	VoicemailRetentionPolicyType *string `json:"voicemailRetentionPolicyType,omitempty"`


	// NumberOfDays - If retentionPolicyType == RETAIN_WITH_TTL, then this value represents the number of days for the TTL
	NumberOfDays *int `json:"numberOfDays,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailretentionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
