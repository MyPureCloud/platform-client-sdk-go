package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replaceresponse
type Replaceresponse struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// UploadStatus
	UploadStatus *Domainentityref `json:"uploadStatus,omitempty"`


	// UploadDestinationUri
	UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`


	// UploadMethod
	UploadMethod *string `json:"uploadMethod,omitempty"`

}

func (u *Replaceresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replaceresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		UploadStatus *Domainentityref `json:"uploadStatus,omitempty"`
		
		UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`
		
		UploadMethod *string `json:"uploadMethod,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ChangeNumber: u.ChangeNumber,
		
		UploadStatus: u.UploadStatus,
		
		UploadDestinationUri: u.UploadDestinationUri,
		
		UploadMethod: u.UploadMethod,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Replaceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
