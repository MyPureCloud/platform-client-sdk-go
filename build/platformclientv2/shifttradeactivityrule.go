package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradeactivityrule
type Shifttradeactivityrule struct { 
	// ActivityCategory - The activity category to which to apply this rule
	ActivityCategory *string `json:"activityCategory,omitempty"`


	// Action - The action this rule invokes
	Action *string `json:"action,omitempty"`


	// ActivityCodeIdReplacement - The activity code ID with which to replace activities belonging to the original category if applicable (required if action == Replace, must be a default activity code ID)
	ActivityCodeIdReplacement *string `json:"activityCodeIdReplacement,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradeactivityrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
