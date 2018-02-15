import math

def fitnessInit(solution,x,y,m,n,t,q,Q,d,D):
    fitness=0
    for routenr in range(len(solution)):
        qvalue=0
        depotNumber=routenr//m[0]
        dvalue=routeLength(solution[routenr],x,y,n,depotNumber)
        for customer in solution[routenr]:
            qvalue+=q[customer]
            dvalue+=d[customer]
        if dvalue>D[0]:
            fitness+=(D[0]-dvalue)**2
        if qvalue>Q[0]:
            fitness+=(Q[0]-qvalue)**2
    return fitness



def fitness(solution, x, y, m, n, t):
    fitness = 0
    numberOfCars = m[0]*t[0]
    for i in range(numberOfCars):
        depotNumber = i//m[0]
        fitness+= routeLength(solution[i],x,y,n,depotNumber)
    return fitness


def euclidianDist(x1, y1, x2, y2):
    return math.sqrt((x2-x1)**2+(y2-y1)**2)

def routeLength(carRoute,x,y,n,depotNumber):
    length=0
    if carRoute:
        depotX=x[depotNumber+n[0]]
        depotY=y[depotNumber+n[0]]
        length += euclidianDist(depotX, depotY, x[carRoute[0]], y[carRoute[0]])
        for j in range(1,len(carRoute)):
            length += euclidianDist(x[carRoute[j-1]], y[carRoute[j-1]], x[carRoute[j]], y[carRoute[j]])
        length+= euclidianDist(x[carRoute[-1]], y[carRoute[-1]], depotX, depotY)
    return length