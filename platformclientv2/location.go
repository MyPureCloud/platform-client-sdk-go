package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Location
type Location struct { 
	// Id - Unique identifier for the location
	Id *string `json:"id,omitempty"`


	// FloorplanId - Unique identifier for the location floorplan image
	FloorplanId *string `json:"floorplanId,omitempty"`


	// Coordinates - Users coordinates on the floorplan. Only used when floorplanImage is set
	Coordinates *map[string]float64 `json:"coordinates,omitempty"`


	// Notes - Optional description on the users location
	Notes *string `json:"notes,omitempty"`


	// LocationDefinition
	LocationDefinition *Locationdefinition `json:"locationDefinition,omitempty"`

}

func (u *Location) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Location

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		FloorplanId *string `json:"floorplanId,omitempty"`
		
		Coordinates *map[string]float64 `json:"coordinates,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		LocationDefinition *Locationdefinition `json:"locationDefinition,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		FloorplanId: u.FloorplanId,
		
		Coordinates: u.Coordinates,
		
		Notes: u.Notes,
		
		LocationDefinition: u.LocationDefinition,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Location) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
