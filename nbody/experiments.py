import matplotlib.pyplot as plt
import constants as c

cpu_count = [1, 2, 3, 4]
body_count = [200, 150, 100, 50]
cpu_count_reversed = [4, 3, 2, 1]
speedup = []

amdahl_speedup = []
gustafson_speedup = []
s_pct = 0.11
p_pct = 0.89


def plot_results(x, y, ymax, type):
    plt.figure(0)
    plt.plot(x, y, color='blue', marker='o')
    plt.plot(x, ymax, color='red', marker='o', linestyle='dashed')

    plt.title(type+" scaling", {'size': '14'})
    plt.grid('on')
    plt.axis('equal')
    plt.xlabel("CPU cores", {'size': '14'})
    plt.ylabel("Speedup", {'size': '14'})
    plt.savefig('resources/strong_scaling.png', bbox_inches='tight')
    plt.show()


def strong_scaling():
    c.N = 1000
    c.planets = 50

    # imports need to be after changing parameters
    import parallel as p
    import sequential as s

    results = []
    for i in range(15):
        t1, _, _ = s.run_simulation(c.N, c.planets)
        res_i = []
        for cc in cpu_count:
            tn = p.run_simulation_parallel(cc, c.N, c.planets)
            print(t1, tn)
            print(t1 / tn)  # speedup
            res_i.append(t1 / tn)
            if i == 0:
                amdahl_speedup.append(1 / (s_pct + p_pct / cc))
                gustafson_speedup.append(s_pct + p_pct * cc)
        results.append(res_i)
        print(res_i)

    sum1, sum2, sum3, sum4 = 0, 0, 0, 0
    for r in results:
        sum1 += r[0]
        sum2 += r[1]
        sum3 += r[2]
        sum4 += r[3]
    speedup.append(sum1 / 15)
    speedup.append(sum2 / 15)
    speedup.append(sum3 / 15)
    speedup.append(sum4 / 15)
    print(speedup)  # 1000 it 50 n
    plot_results(cpu_count, speedup, amdahl_speedup, "Strong")


def weak_scaling():
    results = []

    for i in range(15):
        res_i = []
        for cc in range(len(cpu_count)):
            c.planets = body_count[cc]
            import sequential as s
            import parallel as p
            t1, p1, p2 = s.run_simulation(c.N, c.planets)
            tn = p.run_simulation_parallel(cpu_count_reversed[cc], c.N, c.planets)
            print(t1, tn)
            print(t1 / tn)  # speedup
            res_i.append(t1 / tn)
            if i == 0:
                gustafson_speedup.append(p2/t1 + p1/t1 * cpu_count_reversed[cc])
        results.append(res_i)
        print(res_i)

    sum1, sum2, sum3, sum4 = 0, 0, 0, 0
    for r in results:
        sum1 += r[0]
        sum2 += r[1]
        sum3 += r[2]
        sum4 += r[3]
    speedup.append(sum1 / 15)
    speedup.append(sum2 / 15)
    speedup.append(sum3 / 15)
    speedup.append(sum4 / 15)
    print(speedup)
    plot_results(cpu_count, speedup[::-1], gustafson_speedup[::-1], "Weak")


if __name__ == '__main__':
    # strong_scaling()
    weak_scaling()

