from random import random, randint, choice
import copy

""" HOW THIS CROSSOVER WORKS
    Take a random chunk from a random car of one parent (P1), with a length between 2 and the length of the car,
    and insert it into a random car of another parent (P2), while deleting the duplicate customers from elsewhere
    in P2.
"""
def crossover(selection,chunkSize,q,Q):
    # select two random solutions from the selection to use for mating
    randParent1 = copy.deepcopy(selection[1])
    randParent2 = copy.deepcopy(selection[0])

    # select two random cars from within the solutions
    relevantCars=[carnr for carnr in range(len(randParent1)) if (len(randParent1[carnr])>chunkSize)]
    if relevantCars:
        randIndex1 = choice(relevantCars)
    else:
        print("failure, chunkSize to large")
        print(chunkSize)
        print([len(car)for car in randParent1])
        return randParent1
    randIndex2 = randint(0, len(randParent2)-1)

    # select two random indices from within Car1 to represent the start and end of the chunk to be transferred
    chunkStart = randint(0,len(randParent1[randIndex1])-chunkSize)
    chunkEnd=chunkStart+chunkSize
    # get the chunk
    chunk = randParent1[randIndex1][chunkStart:chunkEnd]
    capacityChunk=0
    for customer in chunk:
        capacityChunk+=q[customer]
    capacityCar=0
    for customer in randParent2[randIndex2]:
        capacityCar+=q[customer]
    i=0
    if randParent2[randIndex2]:
        i = randint(0, len(randParent2[randIndex2]))
    if not (capacityCar+capacityChunk>Q[0]):
    # delete customers in the chunk from solution 2
        for customer in chunk:
            for car in randParent2:
                if customer in car:
                    car.remove(customer)

        # insert the chunk in Car2
        randParent2[randIndex2][i:i] = chunk
        return randParent2
    k=0
    if(capacityChunk>Q[0]):
        print("chunk invalid")
        return randParent2
    while capacityCar+capacityChunk>Q[0]:
        if (i+k<len(randParent2[randIndex2])):
            capacityCar-=q[randParent2[randIndex2][i+k]]
        else:
            i-=1
            capacityCar-=q[randParent2[randIndex2][i]]
        k+=1
    if k>chunkSize:
        return crossover(selection,chunkSize-1,q,Q)
    if k>len(randParent2[randIndex2]):
        return randParent2

    for t in range(k):
        for carnr in range(len(randParent2)):
            for customernr in range(len(randParent2[carnr])):
                if (randParent2[carnr][customernr]==chunk[t]):
                    randParent2[carnr][customernr],randParent2[randIndex2][i+t]=randParent2[randIndex2][i+t],chunk[t]

    for customer in chunk[k:]:
        for car in randParent2:
            if customer in car:
                car.remove(customer)
    randParent2[randIndex2][i+k:i+k] = chunk[k:]
    return randParent2

def crossoverInit(selection,chunkSize):
    # select two random solutions from the selection to use for mating
    randParent1 = copy.deepcopy(selection[1])
    randParent2 = copy.deepcopy(selection[0])
    
    # select two random cars from within the solutions
    relevantCars=[carnr for carnr in range(len(randParent1)) if (len(randParent1[carnr])>chunkSize)]
    if relevantCars:
        randIndex1 = choice(relevantCars)
    else:
        print("failure, chunkSize to large")
        print(chunkSize)
        print([len(car)for car in randParent1])
        return randParent1
    randIndex2 = randint(0, len(randParent2)-1)

# select two random indices from within Car1 to represent the start and end of the chunk to be transferred
    chunkStart = randint(0,len(randParent1[randIndex1])-chunkSize)
    chunkEnd=chunkStart+chunkSize
    # get the chunk
    chunk = randParent1[randIndex1][chunkStart:chunkEnd]
    i=0
    if randParent2[randIndex2]:
        i = randint(0, len(randParent2[randIndex2]))
    # delete customers in the chunk from solution 2
    for customer in chunk:
        for car in randParent2:
            if customer in car:
                car.remove(customer)
    
        # insert the chunk in Car2
        randParent2[randIndex2][i:i] = chunk
    return randParent2










