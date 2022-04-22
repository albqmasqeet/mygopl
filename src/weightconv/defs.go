package weightconv

import "fmt"

type KiloGram float64
type Pound float64

func (k KiloGram) String() string { return fmt.Sprintf("%gKg", k) }
func (lb Pound) String() string   { return fmt.Sprintf("%glb", lb) }
