from body import Body
from constants import *
import multiprocessing as mp
import numpy as np
import matplotlib.pyplot as plt
from math import sqrt
import time
import random

mercury = Body(xMercury, yMercury, zMercury, vxMercury, vyMercury, vzMercury, mMercury, dMercury)
venus = Body(xVenus, yVenus, zVenus, vxVenus, vyVenus, vzVenus, mVenus, dVenus)
earth = Body(xEarth, yEarth, zEarth, vxEarth, vyEarth, vzEarth, mEarth, dEarth)
mars = Body(xMars, yMars, zMars, vxMars, vyMars, vzMars, mMars, dMars)
jupiter = Body(xJup, yJup, zJup, vxJup, vyJup, vzJup, mJup, dJup)
saturn = Body(xSaturn, ySaturn, zSaturn, vxSaturn, vySaturn, vzSaturn, mSaturn,dSaturn)
uranus = Body(xUranus, yUranus, zUranus, vxUranus, vyUranus, vzUranus, mUranus, dUranus)
neptune = Body(xNeptune, yNeptune, zNeptune, vxNeptune, vyNeptune, vzNeptune, mNeptune, dNeptune)
pluto = Body(xPluto, yPluto, zPluto, vxPluto, vyPluto, vzPluto, mPluto, dPluto)
sun = Body(xSun, ySun, zSun, vxSun, vySun, vzSun, mSun, dSun)

# solar system
s = [sun, mercury, venus, earth, mars, jupiter, saturn, uranus, neptune, pluto]
masses = [mSun, mMercury, mVenus, mEarth, mMars, mJup, mSaturn, mUranus, mNeptune, mPluto]

random.seed(1)
for i in range(planets-10):
    s.append(Body(random.uniform(x1, x2), random.uniform(y1, y2), random.uniform(z1, z2),
                  vx_average, vy_average, vz_average, mass_average, d_average))
    masses.append(mass_average)

# lists used for passing as argument to parallel tasks
current_pos_x = np.zeros(planets)
current_pos_y = np.zeros(planets)
current_pos_z = np.zeros(planets)

for i in range(planets):
    current_pos_x[i] = s[i].x
    current_pos_y[i] = s[i].y
    current_pos_z[i] = s[i].z
    s[i].x_hist[0] = s[i].x
    s[i].y_hist[0] = s[i].y
    s[i].z_hist[0] = s[i].z
    s[i].vx_hist[0] = s[i].vx * dayToYear
    s[i].vy_hist[0] = s[i].vy * dayToYear
    s[i].vz_hist[0] = s[i].vz * dayToYear


def magnitude(x, y, z):
    return sqrt(x ** 2 + y ** 2 + z ** 2)


# update position of planet with Leap frog method,
# 2*dt is the length of the time step for leap frog
def position(x_old, y_old, z_old, vx, vy, vz, dt):
    x_new = x_old + 2 * dt * vx
    y_new = y_old + 2 * dt * vy
    z_new = z_old + 2 * dt * vz
    return x_new, y_new, z_new


def position_euler(x, y, z, vx, vy, vz, dt):  # Get position using the Euler method
    x_new = x + vx * dt
    y_new = y + vy * dt
    z_new = z + vz * dt
    return x_new, y_new, z_new


def acceleration(pos_x, pos_y, pos_z, j1, j2):  # get acceleration of celestial body
    sum_ax, sum_ay, sum_az = [], [], []
    for j in range(j1, j2):
        a1, a2, a3 = 0, 0, 0
        for k in range(planets):
            if j != k:
                mag = magnitude(pos_x[j]-pos_x[k], pos_y[j]-pos_y[k], pos_z[j]-pos_z[k])
                ax = G * masses[k] * (pos_x[k] - pos_x[j]) / mag ** 3
                ay = G * masses[k] * (pos_y[k] - pos_y[j]) / mag ** 3
                az = G * masses[k] * (pos_z[k] - pos_z[j]) / mag ** 3
                a1 += ax
                a2 += ay
                a3 += az
        sum_ax.append(a1)
        sum_ay.append(a2)
        sum_az.append(a3)

    return sum_ax, sum_ay, sum_az


def velocity(vx_old, vy_old, vz_old, ax, ay, az, dt):  # get velocity of planet with leap frog method
    vx_new = vx_old + 2 * dt * ax
    vy_new = vy_old + 2 * dt * ay
    vz_new = vz_old + 2 * dt * az

    return vx_new, vy_new, vz_new


def velocity_euler(vx, vy, vz, ax, ay, az, dt):  # Euler method # v = v0 + at
    vx_new = vx + ax * dt
    vy_new = vy + ay * dt
    vz_new = vz + az * dt

    return vx_new, vy_new, vz_new


def generate_file():
    file = open("resources/parallel.txt", "w")
    for i in range(N):
        file.write("ITERATION" + str(i) + "\n")
        for j in range(planets):
            file.write(str(s[j].x_hist[i]) + "," + str(s[j].y_hist[i]) + "," + str(s[j].z_hist[i])
                       + "," + str(s[j].vx_hist[i]) + "," + str(s[j].vy_hist[i]) + "," + str(s[j].vz_hist[i]) + "\n")
    file.close()


def plot_results():
    plt.figure(0)
    for i in range(planets):
        plt.plot(s[i].x_hist, s[i].y_hist, color=[(i % 2) / 2, (i % 3) / 3, (((planets - 1) + i) % 4) / 4])

    plt.title("Solar System over " + str(N * dt) + " years", {'size': '14'})
    plt.grid('on')
    plt.axis('equal')
    plt.legend(['Sun', 'Mercury', 'Venus', 'Earth', 'Mars', 'Jupiter', 'Saturn', 'Uranus', 'Neptune', 'Pluto'])
    plt.xlabel("x (AU)", {'size': '14'})
    plt.ylabel("y (AU)", {'size': '14'})
    # plt.savefig('resources/NBodyOrbit10_parallel.png', bbox_inches='tight')
    plt.show()


if __name__ == '__main__':
    print("Calculating...")
    start = time.time()
    pool = mp.Pool(mp.cpu_count())  # cpu count is 4
    for i in range(N):
        # acceleration
        array = [[], [], []]
        # parallel
        results = pool.starmap(acceleration, [(current_pos_x, current_pos_y, current_pos_z, 0, round(planets/4)),
                                          (current_pos_x, current_pos_y, current_pos_z, round(planets/4), round(planets/2)),
                                          (current_pos_x, current_pos_y, current_pos_z, round(planets/2), round(3 * planets / 4)),
                                          (current_pos_x, current_pos_y, current_pos_z, round(3 * planets / 4), planets)])

        # create preferred structure from results
        for t in range(4):  # 4 tasks
            for l in range(3):  # 3 dimensions
                array[l % 3] += results[t][l]

        # velocity and position update
        # at the beginning of the array, you can't use leap frog, so it suffices to use the 1st order Euler method
        if i == 0:
            for j in range(planets):
                # velocity
                vx_new, vy_new, vz_new = velocity_euler(s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i],
                                                        array[0][j], array[1][j], array[2][j], dt)
                s[j].vx_hist[i+1], s[j].vy_hist[i+1], s[j].vz_hist[i+1] = vx_new, vy_new, vz_new
                # position
                x_new, y_new, z_new = position_euler(s[j].x_hist[i], s[j].y_hist[i], s[j].z_hist[i],
                                                     s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i], dt)

                s[j].x_hist[i + 1], s[j].y_hist[i + 1], s[j].z_hist[i + 1] = x_new, y_new, z_new
                current_pos_x[j], current_pos_y[j], current_pos_z[j] = x_new, y_new, z_new
        else:  # use leap frog method to update velocity
            for j in range(planets):
                # velocity
                vx_new, vy_new, vz_new = velocity(s[j].vx_hist[i-1], s[j].vy_hist[i-1], s[j].vz_hist[i-1],
                                                  array[0][j], array[1][j], array[2][j], dt)
                s[j].vx_hist[i + 1], s[j].vy_hist[i + 1], s[j].vz_hist[i + 1] = vx_new, vy_new, vz_new
                # position
                x_new, y_new, z_new = position(s[j].x_hist[i-1], s[j].y_hist[i-1], s[j].z_hist[i-1],
                                               s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i], dt)

                s[j].x_hist[i + 1], s[j].y_hist[i + 1], s[j].z_hist[i + 1] = x_new, y_new, z_new
                current_pos_x[j], current_pos_y[j], current_pos_z[j] = x_new, y_new, z_new

    end = time.time()
    print(end-start)

    # generate_file()
    plot_results()
