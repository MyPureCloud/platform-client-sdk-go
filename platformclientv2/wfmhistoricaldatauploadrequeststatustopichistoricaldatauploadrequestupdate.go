package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate

	

	return json.Marshal(&struct { 
		RequestId *string `json:"requestId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *string `json:"error,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		RequestId: u.RequestId,
		
		Status: u.Status,
		
		VarError: u.VarError,
		
		Active: u.Active,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
