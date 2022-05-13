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

func (o *Location) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		FloorplanId: o.FloorplanId,
		
		Coordinates: o.Coordinates,
		
		Notes: o.Notes,
		
		LocationDefinition: o.LocationDefinition,
		Alias:    (*Alias)(o),
	})
}

func (o *Location) UnmarshalJSON(b []byte) error {
	var LocationMap map[string]interface{}
	err := json.Unmarshal(b, &LocationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LocationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if FloorplanId, ok := LocationMap["floorplanId"].(string); ok {
		o.FloorplanId = &FloorplanId
	}
    
	if Coordinates, ok := LocationMap["coordinates"].(map[string]interface{}); ok {
		CoordinatesString, _ := json.Marshal(Coordinates)
		json.Unmarshal(CoordinatesString, &o.Coordinates)
	}
	
	if Notes, ok := LocationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if LocationDefinition, ok := LocationMap["locationDefinition"].(map[string]interface{}); ok {
		LocationDefinitionString, _ := json.Marshal(LocationDefinition)
		json.Unmarshal(LocationDefinitionString, &o.LocationDefinition)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Location) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
