package massconv

import (
	"fmt"
)

type KiloGram float64
type Pound float64

func (kg KiloGram) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func (lb Pound) String() string {
	return fmt.Sprintf("%glb", lb)
}
