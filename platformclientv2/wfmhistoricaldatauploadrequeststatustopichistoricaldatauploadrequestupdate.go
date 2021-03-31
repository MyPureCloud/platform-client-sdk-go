package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate
type Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate struct { 
	// RequestId
	RequestId *string `json:"requestId,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// VarError
	VarError *string `json:"error,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
