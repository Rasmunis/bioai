import math


def fitness(solution, x, y, m, n, t):
    fitness = 0
    numberOfCars = m[0]*t[0]

    for i in range(numberOfCars):
        depotNumber = i//m[0]
        depotX = x[n[0]+depotNumber]
        depotY = y[n[0]+depotNumber]
        carRoute = solution[i]
        if len(carRoute):
            fitness += euclidianDist(depotX, depotY, x[carRoute[0]], y[carRoute[0]])
            for j in range(1,len(carRoute)):
                fitness += euclidianDist(x[carRoute[j-1]], y[carRoute[j-1]], x[carRoute[j]], y[carRoute[j]])
            fitness += euclidianDist(x[carRoute[-1]], y[carRoute[-1]], depotX, depotY)
    return fitness


def euclidianDist(x1, y1, x2, y2):
    return math.sqrt((x2-x1)**2+(y2-y1)**2)