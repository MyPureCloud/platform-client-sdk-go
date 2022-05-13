package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetentry
type Facetentry struct { 
	// Attribute
	Attribute *Termattribute `json:"attribute,omitempty"`


	// Statistics
	Statistics *Facetstatistics `json:"statistics,omitempty"`


	// Other
	Other *int `json:"other,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// Missing
	Missing *int `json:"missing,omitempty"`


	// TermCount
	TermCount *int `json:"termCount,omitempty"`


	// TermType
	TermType *string `json:"termType,omitempty"`


	// Terms
	Terms *[]Facetterm `json:"terms,omitempty"`

}

func (o *Facetentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetentry
	
	return json.Marshal(&struct { 
		Attribute *Termattribute `json:"attribute,omitempty"`
		
		Statistics *Facetstatistics `json:"statistics,omitempty"`
		
		Other *int `json:"other,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		Missing *int `json:"missing,omitempty"`
		
		TermCount *int `json:"termCount,omitempty"`
		
		TermType *string `json:"termType,omitempty"`
		
		Terms *[]Facetterm `json:"terms,omitempty"`
		*Alias
	}{ 
		Attribute: o.Attribute,
		
		Statistics: o.Statistics,
		
		Other: o.Other,
		
		Total: o.Total,
		
		Missing: o.Missing,
		
		TermCount: o.TermCount,
		
		TermType: o.TermType,
		
		Terms: o.Terms,
		Alias:    (*Alias)(o),
	})
}

func (o *Facetentry) UnmarshalJSON(b []byte) error {
	var FacetentryMap map[string]interface{}
	err := json.Unmarshal(b, &FacetentryMap)
	if err != nil {
		return err
	}
	
	if Attribute, ok := FacetentryMap["attribute"].(map[string]interface{}); ok {
		AttributeString, _ := json.Marshal(Attribute)
		json.Unmarshal(AttributeString, &o.Attribute)
	}
	
	if Statistics, ok := FacetentryMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	
	if Other, ok := FacetentryMap["other"].(float64); ok {
		OtherInt := int(Other)
		o.Other = &OtherInt
	}
	
	if Total, ok := FacetentryMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if Missing, ok := FacetentryMap["missing"].(float64); ok {
		MissingInt := int(Missing)
		o.Missing = &MissingInt
	}
	
	if TermCount, ok := FacetentryMap["termCount"].(float64); ok {
		TermCountInt := int(TermCount)
		o.TermCount = &TermCountInt
	}
	
	if TermType, ok := FacetentryMap["termType"].(string); ok {
		o.TermType = &TermType
	}
    
	if Terms, ok := FacetentryMap["terms"].([]interface{}); ok {
		TermsString, _ := json.Marshal(Terms)
		json.Unmarshal(TermsString, &o.Terms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
