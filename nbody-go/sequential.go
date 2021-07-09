package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/teodoranedic/ntp2021/types"
)

var mer = types.Body{X: types.XMercury, Y: types.YMercury, Z: types.ZMercury, VX: types.VXMercury, VY: types.VYMercury,
	VZ: types.VZMercury, Mass: types.MMercury, Diameter: types.DMercury, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var ven = types.Body{X: types.XVenus, Y: types.YVenus, Z: types.ZVenus, VX: types.VXVenus, VY: types.VYVenus,
	VZ: types.VZVenus, Mass: types.MVenus, Diameter: types.DVenus, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var mar = types.Body{X: types.XMars, Y: types.YMars, Z: types.ZMars, VX: types.VXMars, VY: types.VYMars,
	VZ: types.VZMars, Mass: types.MMars, Diameter: types.DMars, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var ear = types.Body{X: types.XEarth, Y: types.YEarth, Z: types.ZEarth, VX: types.VXEarth, VY: types.VYEarth,
	VZ: types.VZEarth, Mass: types.MEarth, Diameter: types.DEarth, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var jupit = types.Body{X: types.XJup, Y: types.YJup, Z: types.ZJup, VX: types.VXJup, VY: types.VYJup,
	VZ: types.VZJup, Mass: types.MJup, Diameter: types.DJup, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var satu = types.Body{X: types.XSaturn, Y: types.YSaturn, Z: types.ZSaturn, VX: types.VXSaturn, VY: types.VYSaturn,
	VZ: types.VZSaturn, Mass: types.MSaturn, Diameter: types.DSaturn, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var uran = types.Body{X: types.XUranus, Y: types.YUranus, Z: types.ZUranus, VX: types.VXUranus, VY: types.VYUranus,
	VZ: types.VZUranus, Mass: types.MUranus, Diameter: types.DUranus, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var nept = types.Body{X: types.XNeptune, Y: types.YNeptune, Z: types.ZNeptune, VX: types.VXNeptune, VY: types.VYNeptune,
	VZ: types.VZNeptune, Mass: types.MNeptune, Diameter: types.DNeptune, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var plut = types.Body{X: types.XPluto, Y: types.YPluto, Z: types.ZPluto, VX: types.VXPluto, VY: types.VYPluto,
	VZ: types.VZPluto, Mass: types.MPluto, Diameter: types.DPluto, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

var su = types.Body{X: types.XSun, Y: types.YSun, Z: types.ZSun, VX: types.VXSun, VY: types.VYSun,
	VZ: types.VZSun, Mass: types.MSun, Diameter: types.DSun, X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
	VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
	AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}

// solar system
var s []types.Body = []types.Body{su, mer, ven, ear, mar, jupit, satu, uran, nept, plut}
var masses []float64 = []float64{types.MSun, types.MMercury, types.MVenus, types.MEarth, types.MMars, types.MJup, types.MSaturn, types.MUranus, types.MNeptune, types.MPluto}

func magnitude(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}

// # update position of planet with Leap frog method,
// # 2*dt is the length of the time step for leap frog
func position(x_old, y_old, z_old, vx, vy, vz, dt float64) (float64, float64, float64) {
	x_new := x_old + 2*dt*vx
	y_new := y_old + 2*dt*vy
	z_new := z_old + 2*dt*vz
	return x_new, y_new, z_new
}

// # get position using the Euler method
func position_euler(x, y, z, vx, vy, vz, dt float64) (float64, float64, float64) {
	x_new := x + vx*dt
	y_new := y + vy*dt
	z_new := z + vz*dt
	return x_new, y_new, z_new
}

// # get acceleration of celestial body
func acceleration(a, b, c, G float64, j, i int) {
	// reset acceleration
	s[j].AX = 0.0
	s[j].AY = 0.0
	s[j].AZ = 0.0
	for k := 0; k < types.Planets; k++ {
		if j != k {
			mag := magnitude(a-s[k].X_hist[i], b-s[k].Y_hist[i], c-s[k].Z_hist[i])
			ax := G * masses[k] * (s[k].X_hist[i] - a) / math.Pow(mag, 3)
			ay := G * masses[k] * (s[k].Y_hist[i] - b) / math.Pow(mag, 3)
			az := G * masses[k] * (s[k].Z_hist[i] - c) / math.Pow(mag, 3)
			s[j].AX += ax
			s[j].AY += ay
			s[j].AZ += az
		}
	}
}

// # get velocity of planet with leap frog method
func velocity(vx_old, vy_old, vz_old, ax, ay, az, dt float64) (float64, float64, float64) {
	vx_new := vx_old + 2*dt*ax
	vy_new := vy_old + 2*dt*ay
	vz_new := vz_old + 2*dt*az

	return vx_new, vy_new, vz_new
}

func velocity_euler(vx, vy, vz, ax, ay, az, dt float64) (float64, float64, float64) {
	vx_new := vx + ax*dt
	vy_new := vy + ay*dt
	vz_new := vz + az*dt

	return vx_new, vy_new, vz_new
}

func run_simulation(N, planets int) { //(float64, float64, float64) {
	// start := time.Now()
	// p1 = 0
	// p2 = 0

	// print(N, planets)
	for i := 0; i < N; i++ {
		// # find acceleration due to gravity
		// s1 = time.time()
		for j := 0; j < planets; j++ {
			acceleration(s[j].X_hist[i], s[j].Y_hist[i], s[j].Z_hist[i], types.G, j, i)
		}

		// e1 = time.time()

		//# velocity and position update
		//# at the beginning of the array, you can't use leap frog, so it suffices to use the 1st order Euler method
		// s2 = time.time()
		if i == 0 {
			for j := 0; j < planets; j++ {
				// # velocity
				vx_new, vy_new, vz_new := velocity_euler(s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i],
					s[j].AX, s[j].AY, s[j].AZ, types.Dt)
				s[j].VX_hist[i+1], s[j].VY_hist[i+1], s[j].VZ_hist[i+1] = vx_new, vy_new, vz_new
				// # position
				x_new, y_new, z_new := position_euler(s[j].X_hist[i], s[j].Y_hist[i], s[j].Z_hist[i],
					s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i], types.Dt)

				s[j].X_hist[i+1], s[j].Y_hist[i+1], s[j].Z_hist[i+1] = x_new, y_new, z_new
			}
		} else { // # use leap frog method
			for j := 0; j < planets; j++ {
				// # velocity
				vx_new, vy_new, vz_new := velocity(s[j].VX_hist[i-1], s[j].VY_hist[i-1], s[j].VZ_hist[i-1],
					s[j].AX, s[j].AY, s[j].AZ, types.Dt)
				s[j].VX_hist[i+1], s[j].VY_hist[i+1], s[j].VZ_hist[i+1] = vx_new, vy_new, vz_new
				// # position
				x_new, y_new, z_new := position(s[j].X_hist[i-1], s[j].Y_hist[i-1], s[j].Z_hist[i-1],
					s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i], types.Dt)

				s[j].X_hist[i+1], s[j].Y_hist[i+1], s[j].Z_hist[i+1] = x_new, y_new, z_new
			}
		}
		//     e2 = time.time()
		//     p1 += e1 - s1
		//     p2 += e2 - s2
		//

		// end = time.time()
	}
	// return end - start, p1, p2
}

func plot_results() {
	all_points := make([]plotter.XYs, types.Planets)
	for i, planet := range s {
		points := make(plotter.XYs, types.N+1)
		for i := range planet.X_hist {
			points[i].X = planet.X_hist[i]
			points[i].Y = planet.Y_hist[i]
		}
		all_points[i] = points
	}

	p := plot.New()

	p.Title.Text = "Solar system"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Sun", all_points[0],
		"Mercury", all_points[1],
		"Venus", all_points[2],
		"Earth", all_points[3],
		"Mars", all_points[4],
		"Jupiter", all_points[5],
		"Saturn", all_points[6],
		"Uranus", all_points[7],
		"Neptune", all_points[8],
		"Pluto", all_points[9],
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "SolarSystem.png"); err != nil {
		panic(err)
	}
}

func main() {

	for i := 0; i < types.Planets; i++ {
		s[i].X_hist[0] = s[i].X
		s[i].Y_hist[0] = s[i].Y
		s[i].Z_hist[0] = s[i].Z
		s[i].VX_hist[0] = s[i].VX * types.DaYToYear
		s[i].VY_hist[0] = s[i].VY * types.DaYToYear
		s[i].VZ_hist[0] = s[i].VZ * types.DaYToYear
	}

	run_simulation(types.N, types.Planets)

	plot_results()

}
