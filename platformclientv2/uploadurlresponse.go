package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Uploadurlresponse
type Uploadurlresponse struct { 
	// Url - Presigned URL to PUT the file to
	Url *string `json:"url,omitempty"`


	// UploadKey - Key that identifies the file in the storage including the file name
	UploadKey *string `json:"uploadKey,omitempty"`


	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`

}

func (o *Uploadurlresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Uploadurlresponse
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		UploadKey: o.UploadKey,
		
		Headers: o.Headers,
		Alias:    (*Alias)(o),
	})
}

func (o *Uploadurlresponse) UnmarshalJSON(b []byte) error {
	var UploadurlresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UploadurlresponseMap)
	if err != nil {
		return err
	}
	
	if Url, ok := UploadurlresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if UploadKey, ok := UploadurlresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if Headers, ok := UploadurlresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Uploadurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
