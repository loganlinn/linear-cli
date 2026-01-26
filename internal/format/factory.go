package format

// RendererFactory creates Renderer instances based on output type.
// This implements the Strategy Pattern for output format selection.
type RendererFactory struct{}

// NewRendererFactory creates a new RendererFactory.
func NewRendererFactory() *RendererFactory {
	return &RendererFactory{}
}

// GetRenderer returns the appropriate renderer for the given output type.
// Defaults to TextRenderer if the output type is unknown.
func (f *RendererFactory) GetRenderer(outputType OutputType) Renderer {
	switch outputType {
	case OutputJSON:
		return &JSONRenderer{}
	case OutputText:
		return &TextRenderer{}
	default:
		// Fallback to text renderer for safety
		return &TextRenderer{}
	}
}
