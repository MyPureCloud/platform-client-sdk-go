package platformclientv2
import (
	"encoding/json"
)

// Importshorttermforecastrequest - Request body for importing a short term forecast
type Importshorttermforecastrequest struct { 
	// FileName - The file name, if applicable, this data was extracted from (display purposes only)
	FileName *string `json:"fileName,omitempty"`


	// Description - Description for the imported forecast.  Pass an empty string for no description
	Description *string `json:"description,omitempty"`


	// RouteGroupList - The raw data to import
	RouteGroupList *Routegrouplist `json:"routeGroupList,omitempty"`


	// PartialUploadIds - IDs of partial uploads to include in this imported forecast.  Only relevant for large forecasts
	PartialUploadIds *[]string `json:"partialUploadIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importshorttermforecastrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
