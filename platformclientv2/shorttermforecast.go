package platformclientv2
import (
	"time"
	"encoding/json"
)

// Shorttermforecast - Short Term Forecast
type Shorttermforecast struct { 
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


	// SourceData - The source data references and metadata for this forecast
	SourceData *Listwrapperforecastsourcedaypointer `json:"sourceData,omitempty"`


	// ReferenceStartDate - ISO-8601 date that serves as the reference date for interval-based modifications
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// Modifications - The modifications that have been applied to this forecast
	Modifications *Listwrapperwfmforecastmodification `json:"modifications,omitempty"`


	// GenerationResults - Forecast generation results, if applicable
	GenerationResults *Forecastgenerationresult `json:"generationResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecast) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
