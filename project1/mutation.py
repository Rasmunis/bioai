from random import randint


def mutation(solution, mutationType):
    solSize = len(solution)
    randCar1 = randint(0, solSize-1)
    randCar2 = randint(0, solSize-1)

    if mutationType == "switch":
        randIndex1 = randint(0, len(solution[randCar1])-1)
        randIndex2 = randint(0, len(solution[randCar2])-1)
        solution[randCar1][randIndex1], solution[randCar2][randIndex2] = solution[randCar2][randIndex2], solution[randCar1][randIndex1]

    if mutationType == "move":
        randIndex1 = randint(0, len(solution[randCar1])-1)
        randIndex2 = randint(0, len(solution[randCar2])-1)
        solution[randCar2].insert(randIndex2, solution[randCar1].pop(randIndex1))
    return solution


print(mutation([[1,2,3], [4,5,6], [7,8,9], [10,11,12]], 'switch'))