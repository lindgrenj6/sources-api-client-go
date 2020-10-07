/*
 * Sources
 *
 * Sources
 *
 * API version: 3.0.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapi

import (
	"encoding/json"
)

// SourceTypesCollection struct for SourceTypesCollection
type SourceTypesCollection struct {
	Meta *CollectionMetadata `json:"meta,omitempty"`
	Links *CollectionLinks `json:"links,omitempty"`
	Data *[]SourceType `json:"data,omitempty"`
}

// NewSourceTypesCollection instantiates a new SourceTypesCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceTypesCollection() *SourceTypesCollection {
	this := SourceTypesCollection{}
	return &this
}

// NewSourceTypesCollectionWithDefaults instantiates a new SourceTypesCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceTypesCollectionWithDefaults() *SourceTypesCollection {
	this := SourceTypesCollection{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SourceTypesCollection) GetMeta() CollectionMetadata {
	if o == nil || o.Meta == nil {
		var ret CollectionMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesCollection) GetMetaOk() (*CollectionMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SourceTypesCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMetadata and assigns it to the Meta field.
func (o *SourceTypesCollection) SetMeta(v CollectionMetadata) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SourceTypesCollection) GetLinks() CollectionLinks {
	if o == nil || o.Links == nil {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SourceTypesCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *SourceTypesCollection) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SourceTypesCollection) GetData() []SourceType {
	if o == nil || o.Data == nil {
		var ret []SourceType
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTypesCollection) GetDataOk() (*[]SourceType, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SourceTypesCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SourceType and assigns it to the Data field.
func (o *SourceTypesCollection) SetData(v []SourceType) {
	o.Data = &v
}

func (o SourceTypesCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSourceTypesCollection struct {
	value *SourceTypesCollection
	isSet bool
}

func (v NullableSourceTypesCollection) Get() *SourceTypesCollection {
	return v.value
}

func (v *NullableSourceTypesCollection) Set(val *SourceTypesCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceTypesCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceTypesCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceTypesCollection(val *SourceTypesCollection) *NullableSourceTypesCollection {
	return &NullableSourceTypesCollection{value: val, isSet: true}
}

func (v NullableSourceTypesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceTypesCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


