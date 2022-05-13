package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licenseupdatestatus
type Licenseupdatestatus struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`


	// LicenseId
	LicenseId *string `json:"licenseId,omitempty"`


	// Result
	Result *string `json:"result,omitempty"`

}

func (o *Licenseupdatestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseupdatestatus
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		Result *string `json:"result,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		LicenseId: o.LicenseId,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Licenseupdatestatus) UnmarshalJSON(b []byte) error {
	var LicenseupdatestatusMap map[string]interface{}
	err := json.Unmarshal(b, &LicenseupdatestatusMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := LicenseupdatestatusMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if LicenseId, ok := LicenseupdatestatusMap["licenseId"].(string); ok {
		o.LicenseId = &LicenseId
	}
    
	if Result, ok := LicenseupdatestatusMap["result"].(string); ok {
		o.Result = &Result
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Licenseupdatestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
