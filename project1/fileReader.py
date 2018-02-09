from random import randint, choice, random
from fitness import fitness
from mutation import mutation
from crossover import crossover
from math import floor
import copy
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.collections import LineCollection
from matplotlib.colors import ListedColormap, BoundaryNorm
import pylab as pl
from matplotlib import collections as mc

def plot(solution, x,y,m,n,t):
    totVehicles=m[0]*t[0]
    color_list = plt.cm.Set3(np.linspace(0, 1, totVehicles))
    fig, ax=pl.subplots()
    plt.hold(True)
    for i in range(0,totVehicles):
        color=color_list[i]
        depotIndex= n[0]+int(i/m[0])
        segments=[]
        if solution[i]:
            segments.append([(x[depotIndex], y[depotIndex]), (x[solution[i][0]], y[solution[i][0]])])
        l = len(solution[i])
        this=0
        next=0
        for j in range(0,l-1):
            this=solution[i][j]
            next=solution[i][j+1]
            segments.append([(x[this],y[this]), (x[next], y[next])])
        segments.append([(x[next],y[next]), (x[depotIndex], y[depotIndex])])
        lc=mc.LineCollection(segments, colors=color, linewidths=3)
        ax.add_collection(lc)
        ax.autoscale()
        ax.margins(0.1)
        plt.plot(x[n[0]:n[0]+t[0]], y[n[0]:n[0]+t[0]], 'ro')
    plt.show()


def reader(filename,x,y,D,d,q,Q,m,n,t):
    file=open(filename,'r')
    m[0],n[0],t[0] = [int(i) for i in next(file).split()]
    
    for i in range(1,t[0]+1):
        array=next(file).split()
        D.append(int(array[0]))
        Q[0]=int(array[1])
    for line in file:
        array=[int(i) for i in line.split()]
        x.append(array[1])
        y.append(array[2])
        d.append(array[3])
        q.append(array[4])

def genRandSol(m,n,t):
    totVehicles=m[0]*t[0]
    solution = []
    for i in range(0,totVehicles):
        solution.append([])
    for i in range (0,n[0]):
        vehicle=randint(0,totVehicles-1)
        solution[vehicle].append(i)
    return solution


def clusterSol(x,y,m,n,t):
    solution=[]
    for i in range(m[0]*t[0]):
        solution.append([])
    for i in range(n[0]):
        k=0
        current=(x[k+n[0]]-x[i])**2+(y[k+n[0]]-y[i])**2
        for j in range(1,t[0]):
            next=(x[j+n[0]]-x[i])**2+(y[j+n[0]]-y[i])**2
            if (next<current):
                current=next
                k=j
        solution[k*m[0]].append(i)
    return solution


def main(mutationRate, survivalProp, initPopulation, generations, crossoverRate):
    x=[]
    y=[]
    D=[]
    d=[]
    q=[]
    Q=[0]
    m=[0]
    n=[0]
    t=[0]
    reader('p01.txt',x,y,D,d,q,Q,m,n,t)

    population = [clusterSol(x,y,m,n,t) for it in range(initPopulation)]

    for i in range(generations):
        population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
        selection = population[:floor(survivalProp*len(population))]
        population = copy.deepcopy(selection)
        i = 0
        while len(population) < initPopulation:
            if random() < mutationRate:
                child = mutation(copy.deepcopy(selection[i % len(selection)]), choice(["switch", "move"]))
            elif random() < crossoverRate:
                child = crossover(copy.deepcopy(selection[i % len(selection)]))
            else:
                child = copy.deepcopy(selection[i % len(selection)])
            population.append(child)
            i += 1
            #print(fitness(population[0],x,y,m,n,t))
    population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
    print(fitness(population[0],x,y,m,n,t))
    plot(population[0], x, y, m, n, t)


main(1, 0.1, 100, 5000, 0.7)

