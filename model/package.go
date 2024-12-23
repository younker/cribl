package model

const MAX_VOLUME int = 1000000
const MAX_MASS float32 = 20

// Currently all weights are reported in kilograms
type Mass float32

// Currently all dimensions are assumed to be centimeters
type Dimensions struct {
	Height int
	Width  int
	Depth  int
}

// Represents a physical package on our line. These are managed by robotic arms
// which are bound by what they can maniulate based on dimensions and mass.
// Other attributes (eg contents) are not important at this time for these
// purposes but may be in the future.
type Package struct {
	Mass
	Dimensions
}

func (m Package) IsValid() bool {
	return float32(m.Volume())*float32(m.Mass) > 0
}

func (m Package) Volume() int {
	return m.Height * m.Width * m.Depth
}

func (m Package) IsBulky() bool {
	return m.Volume() >= MAX_VOLUME
}

func (m Package) IsHeavy() bool {
	return float32(m.Mass) >= MAX_MASS
}
