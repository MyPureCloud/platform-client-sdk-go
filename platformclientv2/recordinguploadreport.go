package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordinguploadreport
type Recordinguploadreport struct { 
	// Id - The report id.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// State - The current status of the upload report.
	State *string `json:"state,omitempty"`


	// SignedUrl - For COMPLETED tasks, the signed url to download the report.
	SignedUrl *string `json:"signedUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Recordinguploadreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordinguploadreport
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		SignedUrl *string `json:"signedUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		State: o.State,
		
		SignedUrl: o.SignedUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordinguploadreport) UnmarshalJSON(b []byte) error {
	var RecordinguploadreportMap map[string]interface{}
	err := json.Unmarshal(b, &RecordinguploadreportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordinguploadreportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := RecordinguploadreportMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if State, ok := RecordinguploadreportMap["state"].(string); ok {
		o.State = &State
	}
    
	if SignedUrl, ok := RecordinguploadreportMap["signedUrl"].(string); ok {
		o.SignedUrl = &SignedUrl
	}
    
	if SelfUri, ok := RecordinguploadreportMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordinguploadreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
