package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate
type Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate struct { 
	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
