package lenconv

import (
	"fmt"
)

type Metre float64
type Feet float64

func (m Metre) String() string {
	return fmt.Sprintf("%gm", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%gft", ft)
}
