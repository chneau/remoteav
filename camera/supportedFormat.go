package camera

type SupportedFormat struct {
	name       string
	frameSizes []string
}

func (s SupportedFormat) Name() string {
	return s.name
}

func (s SupportedFormat) FrameSizes() []string {
	return s.frameSizes
}
