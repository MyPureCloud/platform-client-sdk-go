package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimgrouplistresponse - Defines a response for a list of SCIM groups.
type Scimgrouplistresponse struct { 
	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// TotalResults - The total number of results.
	TotalResults *int `json:"totalResults,omitempty"`


	// StartIndex - The 1-based index of the first result returned by this request. Add this to \"itemsPerPage\" when requesting the next page of results.
	StartIndex *int `json:"startIndex,omitempty"`


	// ItemsPerPage - The number of resources returned per page.
	ItemsPerPage *int `json:"itemsPerPage,omitempty"`


	// Resources - The list of requested resources. If \"count\" is 0, then the list will be empty.
	Resources *[]Scimv2group `json:"Resources,omitempty"`

}

func (o *Scimgrouplistresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimgrouplistresponse
	
	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		TotalResults *int `json:"totalResults,omitempty"`
		
		StartIndex *int `json:"startIndex,omitempty"`
		
		ItemsPerPage *int `json:"itemsPerPage,omitempty"`
		
		Resources *[]Scimv2group `json:"Resources,omitempty"`
		*Alias
	}{ 
		Schemas: o.Schemas,
		
		TotalResults: o.TotalResults,
		
		StartIndex: o.StartIndex,
		
		ItemsPerPage: o.ItemsPerPage,
		
		Resources: o.Resources,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimgrouplistresponse) UnmarshalJSON(b []byte) error {
	var ScimgrouplistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ScimgrouplistresponseMap)
	if err != nil {
		return err
	}
	
	if Schemas, ok := ScimgrouplistresponseMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if TotalResults, ok := ScimgrouplistresponseMap["totalResults"].(float64); ok {
		TotalResultsInt := int(TotalResults)
		o.TotalResults = &TotalResultsInt
	}
	
	if StartIndex, ok := ScimgrouplistresponseMap["startIndex"].(float64); ok {
		StartIndexInt := int(StartIndex)
		o.StartIndex = &StartIndexInt
	}
	
	if ItemsPerPage, ok := ScimgrouplistresponseMap["itemsPerPage"].(float64); ok {
		ItemsPerPageInt := int(ItemsPerPage)
		o.ItemsPerPage = &ItemsPerPageInt
	}
	
	if Resources, ok := ScimgrouplistresponseMap["Resources"].([]interface{}); ok {
		ResourcesString, _ := json.Marshal(Resources)
		json.Unmarshal(ResourcesString, &o.Resources)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimgrouplistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
