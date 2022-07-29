package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningcoverartuploadurlrequest
type Learningcoverartuploadurlrequest struct { 
	// FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	FileName *string `json:"fileName,omitempty"`


	// ContentMd5 - Content MD5 of the file to upload
	ContentMd5 *string `json:"contentMd5,omitempty"`


	// SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`


	// ContentType - The content type of the file to upload.
	ContentType *string `json:"contentType,omitempty"`


	// ServerSideEncryption
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty"`

}

func (o *Learningcoverartuploadurlrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningcoverartuploadurlrequest
	
	return json.Marshal(&struct { 
		FileName *string `json:"fileName,omitempty"`
		
		ContentMd5 *string `json:"contentMd5,omitempty"`
		
		SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ServerSideEncryption *string `json:"serverSideEncryption,omitempty"`
		*Alias
	}{ 
		FileName: o.FileName,
		
		ContentMd5: o.ContentMd5,
		
		SignedUrlTimeoutSeconds: o.SignedUrlTimeoutSeconds,
		
		ContentType: o.ContentType,
		
		ServerSideEncryption: o.ServerSideEncryption,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningcoverartuploadurlrequest) UnmarshalJSON(b []byte) error {
	var LearningcoverartuploadurlrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningcoverartuploadurlrequestMap)
	if err != nil {
		return err
	}
	
	if FileName, ok := LearningcoverartuploadurlrequestMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if ContentMd5, ok := LearningcoverartuploadurlrequestMap["contentMd5"].(string); ok {
		o.ContentMd5 = &ContentMd5
	}
    
	if SignedUrlTimeoutSeconds, ok := LearningcoverartuploadurlrequestMap["signedUrlTimeoutSeconds"].(float64); ok {
		SignedUrlTimeoutSecondsInt := int(SignedUrlTimeoutSeconds)
		o.SignedUrlTimeoutSeconds = &SignedUrlTimeoutSecondsInt
	}
	
	if ContentType, ok := LearningcoverartuploadurlrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ServerSideEncryption, ok := LearningcoverartuploadurlrequestMap["serverSideEncryption"].(string); ok {
		o.ServerSideEncryption = &ServerSideEncryption
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningcoverartuploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
