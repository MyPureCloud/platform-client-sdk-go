package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationresponse
type Updatenotificationresponse struct { 
	// MutableGroupId - The mutableGroupId of the notification
	MutableGroupId *string `json:"mutableGroupId,omitempty"`


	// Id - The id of the notification for mapping the potentially new mutableGroupId
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
