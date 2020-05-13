package platformclientv2
import (
	"encoding/json"
)

// Buforecastresultresponse
type Buforecastresultresponse struct { 
	// Result - The result of the operation.  Populated whenever the result is small enough to pass through the api directly
	Result *Buforecastresult `json:"result,omitempty"`


	// DownloadUrl - The download url to fetch the result.  Only populated if the result is too large to pass through the api directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buforecastresultresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
