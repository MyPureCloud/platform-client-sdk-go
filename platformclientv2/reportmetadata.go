package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportmetadata
type Reportmetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Keywords
	Keywords *[]string `json:"keywords,omitempty"`


	// AvailableLocales
	AvailableLocales *[]string `json:"availableLocales,omitempty"`


	// Parameters
	Parameters *[]Parameter `json:"parameters,omitempty"`


	// ExampleUrl
	ExampleUrl *string `json:"exampleUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Reportmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportmetadata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Keywords *[]string `json:"keywords,omitempty"`
		
		AvailableLocales *[]string `json:"availableLocales,omitempty"`
		
		Parameters *[]Parameter `json:"parameters,omitempty"`
		
		ExampleUrl *string `json:"exampleUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Title: o.Title,
		
		Description: o.Description,
		
		Keywords: o.Keywords,
		
		AvailableLocales: o.AvailableLocales,
		
		Parameters: o.Parameters,
		
		ExampleUrl: o.ExampleUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportmetadata) UnmarshalJSON(b []byte) error {
	var ReportmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ReportmetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportmetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportmetadataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Title, ok := ReportmetadataMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ReportmetadataMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Keywords, ok := ReportmetadataMap["keywords"].([]interface{}); ok {
		KeywordsString, _ := json.Marshal(Keywords)
		json.Unmarshal(KeywordsString, &o.Keywords)
	}
	
	if AvailableLocales, ok := ReportmetadataMap["availableLocales"].([]interface{}); ok {
		AvailableLocalesString, _ := json.Marshal(AvailableLocales)
		json.Unmarshal(AvailableLocalesString, &o.AvailableLocales)
	}
	
	if Parameters, ok := ReportmetadataMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ExampleUrl, ok := ReportmetadataMap["exampleUrl"].(string); ok {
		o.ExampleUrl = &ExampleUrl
	}
    
	if SelfUri, ok := ReportmetadataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
