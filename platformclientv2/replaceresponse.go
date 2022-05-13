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

func (o *Replaceresponse) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		ChangeNumber: o.ChangeNumber,
		
		UploadStatus: o.UploadStatus,
		
		UploadDestinationUri: o.UploadDestinationUri,
		
		UploadMethod: o.UploadMethod,
		Alias:    (*Alias)(o),
	})
}

func (o *Replaceresponse) UnmarshalJSON(b []byte) error {
	var ReplaceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ReplaceresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReplaceresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReplaceresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ChangeNumber, ok := ReplaceresponseMap["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if UploadStatus, ok := ReplaceresponseMap["uploadStatus"].(map[string]interface{}); ok {
		UploadStatusString, _ := json.Marshal(UploadStatus)
		json.Unmarshal(UploadStatusString, &o.UploadStatus)
	}
	
	if UploadDestinationUri, ok := ReplaceresponseMap["uploadDestinationUri"].(string); ok {
		o.UploadDestinationUri = &UploadDestinationUri
	}
    
	if UploadMethod, ok := ReplaceresponseMap["uploadMethod"].(string); ok {
		o.UploadMethod = &UploadMethod
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Replaceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
