from body import Body
from constants import *
import matplotlib.pyplot as plt
from math import sqrt
import time
import random

mercury = Body(xMercury, yMercury, zMercury, vxMercury, vyMercury, vzMercury, mMercury, dMercury, N)
venus = Body(xVenus, yVenus, zVenus, vxVenus, vyVenus, vzVenus, mVenus, dVenus, N)
earth = Body(xEarth, yEarth, zEarth, vxEarth, vyEarth, vzEarth, mEarth, dEarth, N)
mars = Body(xMars, yMars, zMars, vxMars, vyMars, vzMars, mMars, dMars, N)
jupiter = Body(xJup, yJup, zJup, vxJup, vyJup, vzJup, mJup, dJup, N)
saturn = Body(xSaturn, ySaturn, zSaturn, vxSaturn, vySaturn, vzSaturn, mSaturn, dSaturn, N)
uranus = Body(xUranus, yUranus, zUranus, vxUranus, vyUranus, vzUranus, mUranus, dUranus, N)
neptune = Body(xNeptune, yNeptune, zNeptune, vxNeptune, vyNeptune, vzNeptune, mNeptune, dNeptune, N)
pluto = Body(xPluto, yPluto, zPluto, vxPluto, vyPluto, vzPluto, mPluto, dPluto, N)
sun = Body(xSun, ySun, zSun, vxSun, vySun, vzSun, mSun, dSun, N)

# solar system
s = [sun, mercury, venus, earth, mars, jupiter, saturn, uranus, neptune, pluto]
masses = [mSun, mMercury, mVenus, mEarth, mMars, mJup, mSaturn, mUranus, mNeptune, mPluto]

random.seed(1)
for i in range(planets - 10):
    s.append(Body(random.uniform(x1, x2), random.uniform(y1, y2), random.uniform(z1, z2),
                  vx_average, vy_average, vz_average, mass_average, d_average, N))
    masses.append(mass_average)

for i in range(planets):
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


def position_euler(x, y, z, vx, vy, vz, dt):  # get position using the Euler method
    x_new = x + vx * dt
    y_new = y + vy * dt
    z_new = z + vz * dt
    return x_new, y_new, z_new


def acceleration(a, b, c, G, j, i):  # get acceleration of celestial body
    # reset acceleration
    s[j].ax = 0
    s[j].ay = 0
    s[j].az = 0
    for k in range(planets):
        if j != k:
            mag = magnitude(a-s[k].x_hist[i], b-s[k].y_hist[i], c-s[k].z_hist[i])
            ax = G * masses[k] * (s[k].x_hist[i] - a) / mag ** 3
            ay = G * masses[k] * (s[k].y_hist[i] - b) / mag ** 3
            az = G * masses[k] * (s[k].z_hist[i] - c) / mag ** 3
            s[j].ax += ax
            s[j].ay += ay
            s[j].az += az


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
    file = open("resources/sequential.txt", "w")
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
    # plt.savefig('resources/NBodyOrbit10.png', bbox_inches='tight')
    plt.show()


def run_simulation(N, planets):
    start = time.time()
    p1 = 0
    p2 = 0

    print(N, planets)
    for i in range(N):
        # find acceleration due to gravity
        s1 = time.time()
        for j in range(planets):
            acceleration(s[j].x_hist[i], s[j].y_hist[i], s[j].z_hist[i], G, j, i)

        e1 = time.time()

        # velocity and position update
        # at the beginning of the array, you can't use leap frog, so it suffices to use the 1st order Euler method
        s2 = time.time()
        if i == 0:
            for j in range(planets):
                # velocity
                vx_new, vy_new, vz_new = velocity_euler(s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i],
                                                        s[j].ax, s[j].ay, s[j].az, dt)
                s[j].vx_hist[i + 1], s[j].vy_hist[i + 1], s[j].vz_hist[i + 1] = vx_new, vy_new, vz_new
                # position
                x_new, y_new, z_new = position_euler(s[j].x_hist[i], s[j].y_hist[i], s[j].z_hist[i],
                                                     s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i], dt)

                s[j].x_hist[i + 1], s[j].y_hist[i + 1], s[j].z_hist[i + 1] = x_new, y_new, z_new
        else:  # use leap frog method
            for j in range(planets):
                # velocity
                vx_new, vy_new, vz_new = velocity(s[j].vx_hist[i - 1], s[j].vy_hist[i - 1], s[j].vz_hist[i - 1],
                                                  s[j].ax, s[j].ay, s[j].az, dt)
                s[j].vx_hist[i + 1], s[j].vy_hist[i + 1], s[j].vz_hist[i + 1] = vx_new, vy_new, vz_new
                # position
                x_new, y_new, z_new = position(s[j].x_hist[i - 1], s[j].y_hist[i - 1], s[j].z_hist[i - 1],
                                               s[j].vx_hist[i], s[j].vy_hist[i], s[j].vz_hist[i], dt)

                s[j].x_hist[i + 1], s[j].y_hist[i + 1], s[j].z_hist[i + 1] = x_new, y_new, z_new

        e2 = time.time()
        p1 += e1 - s1
        p2 += e2 - s2

    end = time.time()

    return end-start, p1, p2


if __name__ == '__main__':
    print("Calculating...")
    p, p1, p2 = run_simulation(N, planets)

    # print(p1)
    # print(p2)
    # print((p1/p)*100, "% can be parallelized")
    # print((p2/p)*100, "% cannot be parallelized")

    # generate_file()
    plot_results()
