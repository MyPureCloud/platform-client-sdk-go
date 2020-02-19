package platformclientv2
import (
	"encoding/json"
)

// Shorttermforecastlistitemresponse - Abbreviated information for a short term forecast to be returned in a list
type Shorttermforecastlistitemresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - The description of the short term forecast
	Description *string `json:"description,omitempty"`


	// CreationMethod - The method used to create this forecast
	CreationMethod *string `json:"creationMethod,omitempty"`


	// Metadata - Metadata for this forecast
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastlistitemresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
