from random import randint, choice, random, sample, seed, shuffle
from fitness import fitness, routeLength, fitnessInit
from mutation import mutation
from crossover import crossover, crossoverInit
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
    for i in range(totVehicles):
        color=color_list[i]
        depotIndex= n[0]+int(i/m[0])
        segments=[]
        if solution[i]:
            segments.append([(x[depotIndex], y[depotIndex]), (x[solution[i][0]], y[solution[i][0]])])
            l = len(solution[i])
            this=0
            next=0
            for j in range(l-1):
                this=solution[i][j]
                next=solution[i][j+1]
                segments.append([(x[this],y[this]), (x[next], y[next])])
            segments.append([(x[next],y[next]), (x[depotIndex], y[depotIndex])])
        
            lc=mc.LineCollection(segments, colors=color, linewidths=1)
            ax.add_collection(lc)
        ax.autoscale()
        ax.margins(0.1)
    plt.plot(x[n[0]:n[0]+t[0]], y[n[0]:n[0]+t[0]], 'x')
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
    for car in solution:
        for j in range(1,len(car)-2):
            closest = j+1
            closestDist =(x[car[closest]]-x[car[j]])**2+(y[car[closest]]-y[car[j]])**2
            for k in range(j+2,len(car)):
                dist = (x[car[k]]-x[car[j]])**2+(y[car[k]]-y[car[j]])**2
                if dist<closestDist:
                    closestDist=dist
                    closest=k
            car[closest],car[j+1]=car[j+1],car[closest]

    return solution


def clusterSol2(x,y,m,n,t):
    solution=clusterSol(x,y,m,n,t)
    for count in range(30):
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
            file.write(str(1+vehiclenr/m[0])+"   "+str((vehiclenr+1)%m[0])+"  ")
            for customer in solution[vehiclenr]:
                duration+=d[customer]
                cost+=q[customer]
            file.write(str(duration)+"   "+str(cost)+"  ")
            for cust in solution[vehiclenr]:
                file.write(str(cust+1)+" ")
            file.write("\n")


def isValid(sol,x,y,q,Q,d,D,m,n):
    if Q==0 and D==0:
        return 1
    for vehiclenr in range(len(sol)):
        depotNumber=vehiclenr//m[0]
        demand=0
        duration=routeLength(sol[vehiclenr],x,y,n,depotNumber)
        for customer in sol[vehiclenr]:
            demand+=q[customer]
            duration+=d[customer]
        if demand>Q or (D>0 and duration>D):
            return 0
    return 1


def genValidSol(x,y,m,n,t,q,Q):
    solution=[]
    qvalues=[]
    dvalues=[]
    carnr=0
    carlist=range(m[0]*t[0])
    shuffle(carlist)
    for car in range(m[0]*t[0]):
        solution.append([])
        qvalues.append(0)
    for customer in range(n[0]):
        count=0
        while qvalues[carlist[carnr]]+q[customer]>Q[0]:
            carnr+=1
            if carnr==m[0]*t[0]:
                carnr=0
            count+=1
            if count>m[0]*t[0]:
                return genValidSol(x,y,m,n,t,q,Q)
        if carnr>m[0]*t[0]-1:
            print("oops")
            for carnr in carlist:
                if qvalues[carnr]+q[customer]<=Q[0]:
                    solution[carnr].append(customer)
                    qvalues[carnr]+=q[customer]
                    break
            return genValidSol(x,y,m,n,t,q,Q)
        else:
            solution[carlist[carnr]].append(customer)
            qvalues[carlist[carnr]]+=q[customer]
    return solution

def evolveInit(populationChunk,x,y,D,d,Q,q,m,n,t,):
    partitionSize=len(populationChunk)
    pressure=2
    crossoverRate=0.6
    populationChunk.sort(key=lambda solution: fitnessInit(solution, x, y, m, n, t,q,Q,d,D))
    newGen=[populationChunk[0]]
    count=0
    while(len(newGen)<partitionSize):
        
        index=int((partitionSize-1)*(random()**pressure))
        chunkSize=randint(2,max([len(car) for car in populationChunk[index]])-1)
        index2=int((partitionSize-1)*(random()**pressure))
        child = crossoverInit([populationChunk[index2], populationChunk[index]],chunkSize)
        newGen.append(child)
    populationChunk=newGen
    return populationChunk

def main(partitionNr,name,initPopulation, generations, crossoverRate,pressure):
    #seed(8)
    def evolve(populationChunk):
        for gen in range(generations):
        #pressure=pressure+stepSize
            populationChunk.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
            fitnessList.append(fitness(populationChunk[0],x,y,m,n,t))
            newGen=[populationChunk[0]]
            while(len(newGen)<partitionSize):
                index=int((partitionSize-1)*(random()**pressure))
                if random()<crossoverRate:
                    chunkSize=randint(2,max([len(car) for car in populationChunk[index]])-1)
                    index2=int((partitionSize-1)*(random()**pressure))
                    child = crossover([populationChunk[index2], populationChunk[index]],chunkSize,q,Q)
                    while not isValid(child,x,y,q,Q[0],d,D[0],m,n):
                        index=int((partitionSize-1)*(random()**pressure))
                        index2=int((partitionSize-1)*(random()**pressure))
                        chunkSize=randint(2,max([len(car) for car in populationChunk[index]])-1)
                        child = crossover([populationChunk[index2], populationChunk[index]],chunkSize,q,Q)
                else:
                    child = mutation(copy.deepcopy(populationChunk[index]),"move")
                if isValid(child,x,y,q,Q[0],d,D[0],m,n):
                    newGen.append(child)
            populationChunk=newGen
        return populationChunk
    x=[]
    y=[]
    D=[]
    d=[]
    q=[]
    Q=[0]
    m=[0]
    n=[0]
    t=[0]
    reader(name+'.txt',x,y,D,d,q,Q,m,n,t)
    fitnessList=[]
    population = [clusterSol2(x,y,m,n,t) for it in range(initPopulation)]
    #population = [genRandSol(m,n,t) for it in range(2*initPopulation)]
    failurelist=[]
    for nr in range(len(population)):
        if not isValid(population[nr],x,y,q,Q[0],d,D[0],m,n):
            failurelist.append(nr)
    for nr in reversed(failurelist):
        population.remove(population[nr])
    if not population:
        population = [genValidSol(x,y,m,n,t,q,Q) for it in range(initPopulation/2)]
        population.sort(key=lambda solution: fitnessInit(solution,x,y,m,n,t,q,Q,d,D))
        count=0
        while not isValid(population[10],x,y,q,Q[0],d,D[0],m,n):
            print(count, fitnessInit(population[0],x,y,m,n,t,q,Q,d,D))
            count+=1
            population=evolveInit(population,x,y,D,d,Q,q,m,n,t)
            if count%100==0:
                print(count, fitnessInit(population[0],x,y,m,n,t,q,Q,d,D))
    failurelist=[]
    for nr in range(len(population)):
        if not isValid(population[nr],x,y,q,Q[0],d,D[0],m,n):
            failurelist.append(nr)
    for nr in reversed(failurelist):
        population.remove(population[nr])
    plot(population[0],x,y,m,n,t)
    print(fitness(population[0],x,y,m,n,t))
    while len(population)<initPopulation:
        population.append(population[0])

    partitionSize=initPopulation/partitionNr
    partitions=range(0,initPopulation,partitionSize)
    partitions.append(initPopulation)
    print(partitions)
    for run in range(partitionNr):
        population[partitions[run]:partitions[run+1]]=copy.deepcopy(evolve(population[partitions[run]:partitions[run+1]]))
        print(fitnessList[len(fitnessList)-1])
    partitionSize=initPopulation
    population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))

    population=evolve(population)
    population.sort(key=lambda solution: fitness(solution, x, y, m, n, t))
    print(fitness(population[0],x,y,m,n,t))
    writeSolutionToFile(name+'solution.txt',population[0],fitness(population[0],x,y,m,n,t),d,q,m,n,t)
    plot(population[0], x, y, m, n, t)
    plot(population[1],x,y,m,n,t)
    plot(crossover([population[1],population[0]], 5, q,Q),x,y,m,n,t)
    plt.plot(range(len(fitnessList)), fitnessList, 'ro')
    plt.axis([0, len(fitnessList), min(fitnessList)-10, max(fitnessList)])
    plt.show()






main(1,'p01',200,100,1,5)







