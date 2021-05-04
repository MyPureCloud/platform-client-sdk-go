package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulerequest - Learning module request
type Learningmodulerequest struct { 
	// Name - The name of learning module
	Name *string `json:"name,omitempty"`


	// Description - The description of learning module
	Description *string `json:"description,omitempty"`


	// CompletionTimeInDays - The completion time of learning module in days
	CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`


	// InformSteps - The list of inform steps in a learning module
	InformSteps *[]Learningmoduleinformsteprequest `json:"informSteps,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningmodulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
