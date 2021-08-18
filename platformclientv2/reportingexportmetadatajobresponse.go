package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Reportingexportmetadatajobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingexportmetadatajobresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ViewType *string `json:"viewType,omitempty"`
		
		DateLimitations *string `json:"dateLimitations,omitempty"`
		
		RequiredFilters *[]string `json:"requiredFilters,omitempty"`
		
		SupportedFilters *[]string `json:"supportedFilters,omitempty"`
		
		RequiredColumnIds *[]string `json:"requiredColumnIds,omitempty"`
		
		DependentColumnIds *map[string][]string `json:"dependentColumnIds,omitempty"`
		
		AvailableColumnIds *[]string `json:"availableColumnIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ViewType: u.ViewType,
		
		DateLimitations: u.DateLimitations,
		
		RequiredFilters: u.RequiredFilters,
		
		SupportedFilters: u.SupportedFilters,
		
		RequiredColumnIds: u.RequiredColumnIds,
		
		DependentColumnIds: u.DependentColumnIds,
		
		AvailableColumnIds: u.AvailableColumnIds,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingexportmetadatajobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
