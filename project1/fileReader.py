from random import randint, choice, random
from fitness import fitness
from mutation import mutation
from math import floor
import copy



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


def main(mutationRate, survivalProp, initPopulation, generations):
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

    population = [genRandSol(m,n,t) for x in range(initPopulation)]

    for i in range(generations):
        population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
        selection = population[:floor(survivalProp*len(population))]
        population = copy.deepcopy(selection)
        i = 0
        while len(population) < initPopulation-len(selection):
            if random() < mutationRate:
                population.append(copy.deepcopy(mutation(selection[i % len(selection)], choice(["switch", "move"]))))
            else:
                population.append(copy.deepcopy(selection[i % len(selection)]))
            i += 1
    print(fitness(selection[0],x,y,m,n,t))

main(0.8, 0.2, 100, 1000)