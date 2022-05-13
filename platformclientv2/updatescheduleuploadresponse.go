package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatescheduleuploadresponse
type Updatescheduleuploadresponse struct { 
	// UploadKey - The key to pass to the secondary request to start processing of the upload
	UploadKey *string `json:"uploadKey,omitempty"`


	// Url - The url to which to PUT the upload body
	Url *string `json:"url,omitempty"`


	// Headers - Required headers for the PUT request to the url
	Headers *map[string]string `json:"headers,omitempty"`


	// UploadBodySchema - Always null. Defines the schema of the json body to be PUT to the url. The json body should be gzip encoded before uploading
	UploadBodySchema *Updatescheduleuploadschema `json:"uploadBodySchema,omitempty"`

}

func (o *Updatescheduleuploadresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatescheduleuploadresponse
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		
		UploadBodySchema *Updatescheduleuploadschema `json:"uploadBodySchema,omitempty"`
		*Alias
	}{ 
		UploadKey: o.UploadKey,
		
		Url: o.Url,
		
		Headers: o.Headers,
		
		UploadBodySchema: o.UploadBodySchema,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatescheduleuploadresponse) UnmarshalJSON(b []byte) error {
	var UpdatescheduleuploadresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatescheduleuploadresponseMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := UpdatescheduleuploadresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if Url, ok := UpdatescheduleuploadresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Headers, ok := UpdatescheduleuploadresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	
	if UploadBodySchema, ok := UpdatescheduleuploadresponseMap["uploadBodySchema"].(map[string]interface{}); ok {
		UploadBodySchemaString, _ := json.Marshal(UploadBodySchema)
		json.Unmarshal(UploadBodySchemaString, &o.UploadBodySchema)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatescheduleuploadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
