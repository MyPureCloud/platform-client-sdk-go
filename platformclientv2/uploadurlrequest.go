package platformclientv2
import (
	"encoding/json"
)

// Uploadurlrequest
type Uploadurlrequest struct { 
	// FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	FileName *string `json:"fileName,omitempty"`


	// ContentMd5 - Content MD-5 of the file to upload
	ContentMd5 *string `json:"contentMd5,omitempty"`


	// SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`


	// ServerSideEncryption
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty"`

}

// String returns a JSON representation of the model
func (o *Uploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
