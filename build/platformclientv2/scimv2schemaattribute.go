package platformclientv2
import (
	"encoding/json"
)

// Scimv2schemaattribute - A complex type that defines service provider attributes, or sub-attributes and their qualities.
type Scimv2schemaattribute struct { 
	// Name - The attribute's name
	Name *string `json:"name,omitempty"`


	// VarType - The attribute's data type.  Valid values are \"string\", \"boolean\", \"decimal\", \"integer\", \"dateTime\", \"reference\", and \"complex\".
	VarType *string `json:"type,omitempty"`


	// SubAttributes - When an attribute is of type \"complex\", \"subAttributes\" defines a set of sub-attributes. \"subAttributes\" has the same schema sub-attributes as \"attributes\"
	SubAttributes *[]Scimv2schemaattribute `json:"subAttributes,omitempty"`


	// MultiValued - A Boolean value indicating the attribute's plurality.
	MultiValued *bool `json:"multiValued,omitempty"`


	// Description - The attribute's human-readable description.
	Description *string `json:"description,omitempty"`


	// Required - A Boolean value that specifies whether or not the attribute is required.
	Required *bool `json:"required,omitempty"`


	// CanonicalValues - A collection of suggested canonical values that MAY be used (e.g., \"work\" and \"home\").  In some cases, service providers MAY choose to ignore unsupported values.  OPTIONAL.
	CanonicalValues *[]string `json:"canonicalValues,omitempty"`


	// CaseExact - A Boolean value that specifies whether or not a string attribute is case sensitive.  The server SHALL use case sensitivity when evaluating filters.  For attributes that are case exact, the server SHALL preserve case for any value submitted.  If the attribute is case insensitive, the server MAY alter case for a submitted value.  Case sensitivity also impacts how attribute values MAY be compared against filter values (see Section 3.4.2.2 of [RFC7644])
	CaseExact *bool `json:"caseExact,omitempty"`


	// Mutability - A single keyword indicating the circumstances under which the value of the attribute can be (re)defined. Value are readOnly, readWrite, immutable, writeOnly
	Mutability *string `json:"mutability,omitempty"`


	// Returned - A single keyword that indicates when an attribute and associated values are returned in response to a GET request, or in response to a PUT, POST, or PATCH request.  Valid keywords are as follows: always, never, default, request
	Returned *string `json:"returned,omitempty"`


	// Uniqueness - A single keyword value that specifies how the service provider enforces uniqueness of attribute values.  A server MAY reject an invalid value based on uniqueness by returning HTTP response code 400 (Bad Request).  A client MAY enforce uniqueness on the client side to a greater degree than the service provider enforces.  For example, a client could make a value unique while the server has uniqueness of \"none\".  Valid keywords are as follows: none, server, global
	Uniqueness *string `json:"uniqueness,omitempty"`


	// ReferenceTypes - A multi-valued array of JSON strings that indicate the SCIM resource types that may be referenced. Values include User, Group, external and uri.
	ReferenceTypes *[]string `json:"referenceTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2schemaattribute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
