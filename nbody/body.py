
import numpy as np


class Body:
    def __init__(self, x, y, z, vx, vy, vz, mass, d, n):
        self.x = x
        self.y = y
        self.z = z
        self.vx = vx
        self.vy = vy
        self.vz = vz
        self.mass = mass
        self.diameter = d

        self.ax = 0.0
        self.ay = 0.0
        self.az = 0.0

        # self.ax_hist = mp.Array('d', N + 1, lock=False)
        # self.ay_hist = mp.Array('d', N + 1, lock=False)
        # self.az_hist = mp.Array('d', N + 1, lock=False)
        #
        # self.x_hist = mp.Array('d', N+1, lock=False)
        # self.y_hist = mp.Array('d', N+1, lock=False)
        # self.z_hist = mp.Array('d', N+1, lock=False)
        #
        # self.vx_hist = mp.Array('d', N+1, lock=False)
        # self.vy_hist = mp.Array('d', N+1, lock=False)
        # self.vz_hist = mp.Array('d', N+1, lock=False)

        self.ax_hist = np.zeros(n + 1)
        self.ay_hist = np.zeros(n + 1)
        self.az_hist = np.zeros(n + 1)

        self.x_hist = np.zeros(n + 1)
        self.y_hist = np.zeros(n + 1)
        self.z_hist = np.zeros(n + 1)

        self.vx_hist = np.zeros(n + 1)
        self.vy_hist = np.zeros(n + 1)
        self.vz_hist = np.zeros(n + 1)
