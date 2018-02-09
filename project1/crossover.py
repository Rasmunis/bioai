from random import randint
import copy

""" HOW THIS CROSSOVER WORKS
    Take a random chunk from a random car of one parent (P1), with a length between 2 and the length of the car,
    and insert it into a random car of another parent (P2), while deleting the duplicate customers from elsewhere
    in P2.
"""
def crossover(selection):
    # select two random solutions from the selection to use for mating
    randParent1 = copy.deepcopy(selection[randint(0, len(selection))])
    randParent2 = copy.deepcopy(selection[randint(0, len(selection))])

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
        for car in randParent2:
            if customer in car:
                car.remove(customer)

    # insert the chunk in Car2
    randParent2[randIndex2].insert(randint(0, len(randParent2[randIndex2])), chunk)

    return randParent2


