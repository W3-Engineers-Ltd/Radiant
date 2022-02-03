// Copyright 2020
//

package bean

// BeanMetadata, in other words, bean's config.
// it could be read from config file
type BeanMetadata struct {
	// Fields: field name => field metadata
	Fields map[string]*FieldMetadata
}

// FieldMetadata contains metadata
type FieldMetadata struct {
	// default value in string format
	DftValue string
}
