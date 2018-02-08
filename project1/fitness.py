from importlib import import_module
import math

import_module('fileReader.py')


def fitness(solution):
    fitness = 0
    numberOfCars = m*t

    for i in range(numberOfCars):
        depotNumber = i//t
        depotX = x[n+depotNumber]
        depotY = y[n+depotNumber]
        carRoute = solution[i]

        fitness += euclidianDist(depotX, depotY, x[carRoute[0]], y[carRoute[0]])
        for j in range(1,len(carRoute)):
            fitness += euclidianDist(x[carRoute[j-1]], y[carRoute[j-1]], x[carRoute[j]], y[carRoute[j]])
        fitness += euclidianDist(x[carRoute[-1]], y[carRoute[-1]], depotX, depotY)

        return fitness


def euclidianDist(x1, y1, x2, y2):
    return math.sqrt((x2-x1)**2+(y2-y1)**2)