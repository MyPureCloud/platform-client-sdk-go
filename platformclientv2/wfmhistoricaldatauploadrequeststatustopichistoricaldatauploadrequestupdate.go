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

func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) MarshalJSON() ([]byte, error) {
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
		RequestId: o.RequestId,
		
		Status: o.Status,
		
		VarError: o.VarError,
		
		Active: o.Active,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) UnmarshalJSON(b []byte) error {
	var WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap)
	if err != nil {
		return err
	}
	
	if RequestId, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["requestId"].(string); ok {
		o.RequestId = &RequestId
	}
	
	if Status, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if VarError, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["error"].(string); ok {
		o.VarError = &VarError
	}
	
	if Active, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if VarType, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
