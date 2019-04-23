package registry

// Blob describes a
type Blob struct {
	// MediaType describe the type of the content. All text based formats are
	// encoded as utf-8.
	MediaType string

	// Size in bytes of content.
	Size int64
}
