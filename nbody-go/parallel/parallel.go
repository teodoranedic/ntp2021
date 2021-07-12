package parallel

import (
	"fmt"
	"math"
	"math/rand"
	"time"

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
var current_pos_x = make([]float64, types.Planets)
var current_pos_y = make([]float64, types.Planets)
var current_pos_z = make([]float64, types.Planets)

func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

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

func acceleration_parallel(pos_x, pos_y, pos_z []float64, j1, j2 int, ch_ax, ch_ay, ch_az chan []float64) {
	sum_ax, sum_ay, sum_az := make([]float64, types.Planets), make([]float64, types.Planets), make([]float64, types.Planets)
	for j := j1; j < j2; j++ {
		a1, a2, a3 := 0.0, 0.0, 0.0
		for k := 0; k < types.Planets; k++ {
			if j != k {
				mag := magnitude(pos_x[j]-pos_x[k], pos_y[j]-pos_y[k], pos_z[j]-pos_z[k])
				ax := types.G * masses[k] * (pos_x[k] - pos_x[j]) / math.Pow(mag, 3)
				ay := types.G * masses[k] * (pos_y[k] - pos_y[j]) / math.Pow(mag, 3)
				az := types.G * masses[k] * (pos_z[k] - pos_z[j]) / math.Pow(mag, 3)
				a1 += ax
				a2 += ay
				a3 += az
			}
		}
		sum_ax[j] = a1
		sum_ay[j] = a2
		sum_az[j] = a3
	}
	ch_ax <- sum_ax
	ch_ay <- sum_ay
	ch_az <- sum_az
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

func add_array(base []float64, adding []float64) {

	for i := range base {
		if adding[i] == 0 {

		} else {
			base[i] = adding[i]
		}
	}
}

func run_simulation(N, planets int) time.Duration { //(float64, float64, float64) {
	start := time.Now()

	for i := 0; i < N; i++ {
		// # find acceleration due to gravity

		ax_array := make([]float64, types.Planets)
		ay_array := make([]float64, types.Planets)
		az_array := make([]float64, types.Planets)

		ch_ax := make(chan []float64, types.Planets)
		ch_ay := make(chan []float64, types.Planets)
		ch_az := make(chan []float64, types.Planets)

		go acceleration_parallel(current_pos_x, current_pos_y, current_pos_z, 0, int(math.Round(float64(types.Planets)/4)), ch_ax, ch_ay, ch_az)
		go acceleration_parallel(current_pos_x, current_pos_y, current_pos_z, int(math.Round(float64(types.Planets)/4)), int(math.Round(float64(types.Planets)/2)), ch_ax, ch_ay, ch_az)
		go acceleration_parallel(current_pos_x, current_pos_y, current_pos_z, int(math.Round(float64(types.Planets)/2)), int(math.Round(3*float64(types.Planets)/4)), ch_ax, ch_ay, ch_az)
		go acceleration_parallel(current_pos_x, current_pos_y, current_pos_z, int(math.Round(3*float64(types.Planets)/4)), types.Planets, ch_ax, ch_ay, ch_az)

		add_array(ax_array, <-ch_ax)
		add_array(ax_array, <-ch_ax)
		add_array(ax_array, <-ch_ax)
		add_array(ax_array, <-ch_ax)
		add_array(ay_array, <-ch_ay)
		add_array(ay_array, <-ch_ay)
		add_array(ay_array, <-ch_ay)
		add_array(ay_array, <-ch_ay)
		add_array(az_array, <-ch_az)
		add_array(az_array, <-ch_az)
		add_array(az_array, <-ch_az)
		add_array(az_array, <-ch_az)

		//# velocity and position update
		//# at the beginning of the array, you can't use leap frog, so it suffices to use the 1st order Euler method
		if i == 0 {
			for j := 0; j < planets; j++ {
				// # velocity
				vx_new, vy_new, vz_new := velocity_euler(s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i],
					ax_array[j], ay_array[j], az_array[j], types.Dt)
				s[j].VX_hist[i+1], s[j].VY_hist[i+1], s[j].VZ_hist[i+1] = vx_new, vy_new, vz_new
				// # position
				x_new, y_new, z_new := position_euler(s[j].X_hist[i], s[j].Y_hist[i], s[j].Z_hist[i],
					s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i], types.Dt)

				s[j].X_hist[i+1], s[j].Y_hist[i+1], s[j].Z_hist[i+1] = x_new, y_new, z_new
				current_pos_x[j], current_pos_y[j], current_pos_z[j] = x_new, y_new, z_new
			}
		} else { // # use leap frog method
			for j := 0; j < planets; j++ {
				// # velocity
				vx_new, vy_new, vz_new := velocity(s[j].VX_hist[i-1], s[j].VY_hist[i-1], s[j].VZ_hist[i-1],
					ax_array[j], ay_array[j], az_array[j], types.Dt)
				s[j].VX_hist[i+1], s[j].VY_hist[i+1], s[j].VZ_hist[i+1] = vx_new, vy_new, vz_new
				// # position
				x_new, y_new, z_new := position(s[j].X_hist[i-1], s[j].Y_hist[i-1], s[j].Z_hist[i-1],
					s[j].VX_hist[i], s[j].VY_hist[i], s[j].VZ_hist[i], types.Dt)

				s[j].X_hist[i+1], s[j].Y_hist[i+1], s[j].Z_hist[i+1] = x_new, y_new, z_new
				current_pos_x[j], current_pos_y[j], current_pos_z[j] = x_new, y_new, z_new
			}
		}
	}
	end := time.Now()
	elapsed_time := end.Sub(start)
	return elapsed_time
}

func plot_results() {
	all_points := make([]plotter.XYs, types.Planets)
	for i, planet := range s {
		points := make(plotter.XYs, types.N+1)
		for j := range planet.X_hist {
			points[j].X = planet.X_hist[j]
			points[j].Y = planet.Y_hist[j]
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
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "SolarSystemParallel.png"); err != nil {
		panic(err)
	}
}

func Parallel() {
	rand.Seed(1)
	for i := 0; i < types.Planets-10; i++ {
		a := types.Body{X: randFloats(types.X1, types.X2), Y: randFloats(types.Y1, types.Y2), Z: randFloats(types.Z1, types.Z2), VX: types.VX_average, VY: types.VY_average,
			VZ: types.VZ_average, Mass: types.Mass_average, Diameter: float64(types.D_average), X_hist: make([]float64, types.N+1), Y_hist: make([]float64, types.N+1), Z_hist: make([]float64, types.N+1),
			VX_hist: make([]float64, types.N+1), VY_hist: make([]float64, types.N+1), VZ_hist: make([]float64, types.N+1), AX: 0.0, AY: 0.0, AZ: 0.0, AX_hist: make([]float64, types.N+1),
			AY_hist: make([]float64, types.N+1), AZ_hist: make([]float64, types.N+1)}
		s = append(s, a)
		masses = append(masses, types.Mass_average)
	}

	for i := 0; i < types.Planets; i++ {
		s[i].X_hist[0] = s[i].X
		s[i].Y_hist[0] = s[i].Y
		s[i].Z_hist[0] = s[i].Z
		s[i].VX_hist[0] = s[i].VX * types.DaYToYear
		s[i].VY_hist[0] = s[i].VY * types.DaYToYear
		s[i].VZ_hist[0] = s[i].VZ * types.DaYToYear
		current_pos_x[i] = s[i].X
		current_pos_y[i] = s[i].Y
		current_pos_z[i] = s[i].Z
	}

	elapsed_time := run_simulation(types.N, types.Planets)
	fmt.Println(elapsed_time)

	plot_results()

}
