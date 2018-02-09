from random import randint

def crossover(selection, x, y, m, n, t):
    # select two random solutions from the selection to use for mating
    randParent1 = selection[randint(0, len(selection))]
    randParent2 = selection[randint(0, len(selection))]

    # select two random cars from within the solutions
    randIndex1 = randint(0, len(randParent1))
    randIndex2 = randint(0, len(randParent2))

    # select two random indices from within Car1 to represent the start and end of the chunk to be transferred
    chunkStart = randint(0, len(randParent1[randIndex1])-2)
    chunkEnd = randint(chunkStart+2, len(randParent1))

    # get the chunk
    chunk = randParent1[randIndex1][chunkStart:chunkEnd]

    # delete customers in the chunk from solution 2
    for customer in chunk:
        

    # insert the chunk in Car2
    randParent2[randIndex2].insert(randint(0, len(randParent2[randIndex2])), chunk)


