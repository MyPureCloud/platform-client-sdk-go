package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Updatebusinessunitrequest
type Updatebusinessunitrequest struct { 
	// Name - The name of the business unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The ID of the division to which the business unit should be moved
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Configuration for the business unit
	Settings *Updatebusinessunitsettings `json:"settings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatebusinessunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
