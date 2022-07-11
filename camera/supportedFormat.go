package camera

type SupportedFormat struct {
	format     string
	frameSizes []string
}

func (s SupportedFormat) Format() string {
	return s.format
}

func (s SupportedFormat) FrameSizes() []string {
	return s.frameSizes
}
