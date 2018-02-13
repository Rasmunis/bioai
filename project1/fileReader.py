from random import randint, choice, random, sample, seed
from fitness import fitness
from mutation import mutation
from crossover import crossover
import copy
import numpy as np
import matplotlib.pyplot as plt
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
        solution[k*m[0]+randint(0,m[0]-1)].append(i)
    return solution


def clusterSol2(x,y,m,n,t):
    solution=clusterSol(x,y,m,n,t)
    for count in range(10):
        solution2=[]
        cmx=[]
        cmy=[]
        for i in range(m[0]*t[0]):
            solution2.append([])
            cmx.append(0)
            cmy.append(0)
            for customer in solution[i]:
                cmx[i]+=x[customer]
                cmy[i]+=y[customer]
            if solution[i]:
                cmx[i]=cmx[i]/len(solution[i])
                cmy[i]=cmy[i]/len(solution[i])
        for i in range(n[0]):
            closest=0
            closestDist=(x[i]-cmx[0])**2+(y[i]-cmy[0])**2
            for j in range(1,m[0]*t[0]):
                if ((x[i]-cmx[j])**2+(y[i]-cmy[j])**2<closestDist):
                    closestDist=(x[i]-cmx[j])**2+(y[i]-cmy[j])**2
                    closest=j
            solution2[closest].append(i)
        solution=copy.deepcopy(solution2)
    
    return solution


def writeSolutionToFile(name,solution,fitness,d,q,m,n,t):
    file=open(name+".txt", "w+")
    file.write(str(fitness)+"\n")
    for vehiclenr in range(m[0]*t[0]):
        if solution[vehiclenr]:
            duration=0
            cost=0
            file.write(str(1+vehiclenr/m[0])+"  "+str((vehiclenr+1)%m[0])+"  ")
            for customer in solution[vehiclenr]:
                duration+=d[customer]
                cost+=q[customer]
            file.write(str(duration)+"  "+str(cost)+"  ")
            for cust in solution[vehiclenr]:
                file.write(str(cust)+" ")
            file.write("\n")


def isValid(sol,q,Q):
    if Q==0:
        return 1
    for vehicle in sol:
        demand=0
        for customer in vehicle:
                demand+=q[customer]
        if demand>Q:
            return 0
    return 1
    
        #population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))

        
        # SELECTION METHOD 1
        # selection = population[:int(survivalProp*len(population))]

        # SELECTION METHOD 2
        #selection = copy.deepcopy([population[0]])
        #selection.extend([population[randint(int(survivalProp/2*len(population)), len(population)-1)] for solu in range(int(survivalProp/2*len(population)))])
        #selection.extend([population[randint(0, int(survivalProp/2*len(population)))] for solu in range(int(survivalProp/2*len(population)))])

        # TWEAK
    """selection = copy.deepcopy([population[0]])
        selection.extend(sample(population[int(survivalProp / 2 * len(population)):], int(survivalProp / 2 * len(population))))
        selection.extend(sample(population[:int(survivalProp / 2 * len(population))], int(survivalProp / 2 * len(population))))

        population = copy.deepcopy(selection)
        i = 0
        while len(population) < initPopulation:
            if random() < mutationRate:
                child = mutation(copy.deepcopy(selection[i % len(selection)]), choice(["switch", "move"]))
            elif random() < crossoverRate:
                child = crossover(selection)
            else:
                child = copy.deepcopy(selection[i % len(selection)])
            if isValid(child,q,Q[0]):
                population.append(child)
            i += 1
        if gen % 100 == 0: print("GEN NO", gen, '\n', 'FITNESS ', fitness(population[0],x,y,m,n,t), '\n')"""

    """selection=[]
        stepSize=int(initPopulation/5)
        selection.extend(sample(population[0:stepSize], int(survivalProp*0.4*initPopulation)))
        selection.extend(sample(population[stepSize:2*stepSize],int(survivalProp*0.3*initPopulation)))
        selection.extend(sample(population[2*stepSize:3*stepSize],int(survivalProp*0.2*initPopulation)))
        selection.extend(sample(population[3*stepSize:4*stepSize], int(survivalProp*0.1*initPopulation)))"""
        
    """population=[copy.deepcopy(best)]
        
        i=0
        while (len(population)<initPopulation):
            if random() < mutationRate:
                child = mutation(copy.deepcopy(selection[i % len(selection)]), "move")
            elif random() < crossoverRate:
                child = crossover(selection)
            else:
                child = copy.deepcopy(selection[i % len(selection)])
            if isValid(child,q,Q[0]):
                population.append(child)
            i+=1"""




    population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
    print(fitness(population[0],x,y,m,n,t))
    writeSolutionToFile("test",population[0],fitness(population[0],x,y,m,n,t),d,q,m,n,t)
    plot(population[0], x, y, m, n, t)
    plt.plot(range(len(fitnessList)), fitnessList, 'ro')
    plt.axis([0, generations, 580, 800])
    plt.show()

def main(initPopulation, generations, crossoverRate,pressure):
    #seed(3)
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
    fitnessList=[]
    solution=clusterSol2(x,y,m,n,t)
    plot(solution,x,y,m,n,t)
    population = [clusterSol2(x,y,m,n,t) for it in range(2*initPopulation)]
    #population = [genRandSol(m,n,t) for it in range(2*initPopulation)]
    for sol in population:
        if (not (isValid(sol,q,Q[0]))):
            population.remove(sol)
    print(len(population))
    stepSize=3.0/generations
    for gen in range(generations):
        #pressure=pressure+stepSize
        population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
        fitnessList.append(fitness(population[0],x,y,m,n,t))
        if gen%10==0:
            print(fitnessList[gen], gen)
        newGen=[]
        while(len(newGen)<initPopulation):
            index=int((initPopulation)*(random()**pressure))
            if random()<crossoverRate:
                index2=int(initPopulation*(random()**pressure))
                child = crossover([population[index], population[index2]])
            else:
                child = mutation(copy.deepcopy(population[index]),"move")
            if isValid(child,q,Q[0]):
                newGen.append(child)
        population=newGen
    population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
    print(fitness(population[0],x,y,m,n,t))
    writeSolutionToFile("test",population[0],fitness(population[0],x,y,m,n,t),d,q,m,n,t)
    plot(population[0], x, y, m, n, t)
    plt.plot(range(len(fitnessList)), fitnessList, 'ro')
    plt.axis([0, generations, 575, 1000])
    plt.show()






main(800,300,0.3,1.45)







