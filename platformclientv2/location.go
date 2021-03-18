package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Location) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
