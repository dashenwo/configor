package memory

import (
	"github.com/dashenwo/configor/loader"
	"github.com/dashenwo/configor/reader"
	"github.com/dashenwo/configor/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
