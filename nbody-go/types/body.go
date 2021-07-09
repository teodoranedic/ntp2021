package types

type Body struct {
	X        float64
	Y        float64
	Z        float64
	VX       float64
	VY       float64
	VZ       float64
	Mass     float64
	Diameter float64

	AX float64
	AY float64
	AZ float64

	AX_hist []float64 // make([]float64, N+1)
	AY_hist []float64
	AZ_hist []float64

	X_hist []float64
	Y_hist []float64
	Z_hist []float64

	VX_hist []float64
	VY_hist []float64
	VZ_hist []float64
}
