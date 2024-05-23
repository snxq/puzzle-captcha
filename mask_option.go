package puzzlecaptcha

type Option func(m *PuzzleMask)

func WithHoles(holes int) Option {
	return func(m *PuzzleMask) {
		m.Holes = holes
	}
}

func WithNotches(notches int) Option {
	return func(m *PuzzleMask) {
		m.notches = notches
	}
}

func WithMaskType(masktype PuzzleMaskType) Option {
	return func(m *PuzzleMask) {
		m.masktype = masktype
	}
}
