package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Voicemailretentionpolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailretentionpolicy
	
	return json.Marshal(&struct { 
		VoicemailRetentionPolicyType *string `json:"voicemailRetentionPolicyType,omitempty"`
		
		NumberOfDays *int `json:"numberOfDays,omitempty"`
		*Alias
	}{ 
		VoicemailRetentionPolicyType: o.VoicemailRetentionPolicyType,
		
		NumberOfDays: o.NumberOfDays,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailretentionpolicy) UnmarshalJSON(b []byte) error {
	var VoicemailretentionpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailretentionpolicyMap)
	if err != nil {
		return err
	}
	
	if VoicemailRetentionPolicyType, ok := VoicemailretentionpolicyMap["voicemailRetentionPolicyType"].(string); ok {
		o.VoicemailRetentionPolicyType = &VoicemailRetentionPolicyType
	}
    
	if NumberOfDays, ok := VoicemailretentionpolicyMap["numberOfDays"].(float64); ok {
		NumberOfDaysInt := int(NumberOfDays)
		o.NumberOfDays = &NumberOfDaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailretentionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
