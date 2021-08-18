package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Failedobject
type Failedobject struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

}

func (u *Failedobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Failedobject

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Version: u.Version,
		
		Name: u.Name,
		
		ErrorCode: u.ErrorCode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Failedobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
