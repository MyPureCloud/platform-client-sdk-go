package platformclientv2
import (
	"encoding/json"
)

// Reportingexportmetadatajobresponse
type Reportingexportmetadatajobresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ViewType - The view type of the export metadata
	ViewType *string `json:"viewType,omitempty"`


	// DateLimitations - The date limitations of the export metadata
	DateLimitations *string `json:"dateLimitations,omitempty"`


	// RequiredFilters - The list of required filters for the export metadata
	RequiredFilters *[]string `json:"requiredFilters,omitempty"`


	// SupportedFilters - The list of supported filters for the export metadata
	SupportedFilters *[]string `json:"supportedFilters,omitempty"`


	// RequiredColumnIds - The list of required column ids for the export metadata
	RequiredColumnIds *[]string `json:"requiredColumnIds,omitempty"`


	// DependentColumnIds - The list of dependent column ids for the export metadata
	DependentColumnIds *map[string][]string `json:"dependentColumnIds,omitempty"`


	// AvailableColumnIds - The list of available column ids for the export metadata
	AvailableColumnIds *[]string `json:"availableColumnIds,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportingexportmetadatajobresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
