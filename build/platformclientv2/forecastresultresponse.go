package platformclientv2
import (
	"encoding/json"
)

// Forecastresultresponse
type Forecastresultresponse struct { 
	// Result - The forecast result.  If null, fetch the result from the url in downloadUrl
	Result *Routegrouplist `json:"result,omitempty"`


	// DownloadUrl - The downloadUrl to fetch the result if it is too large to be sent directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastresultresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
