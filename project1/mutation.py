from random import randint


def mutation(solution, mutationType):
    randint(0,0)
    solSize = len(solution)
    randCar1 = randint(0, solSize-1)
    count=0
    while not solution[randCar1]:
        randCar1 = randint(0, solSize-1)
        count+=1
        if count>100:
            print("warning")
            count=0
    randCar2 = randint(0, solSize-1)
    if solution[randCar1] and solution[randCar2]:
        if mutationType == "switch":
            randIndex1 = randint(0, len(solution[randCar1])-1)
            randIndex2 = randint(0, len(solution[randCar2])-1)
            solution[randCar1][randIndex1], solution[randCar2][randIndex2] = solution[randCar2][randIndex2], solution[randCar1][randIndex1]

        if mutationType == "move":
            randIndex1 = randint(0, len(solution[randCar1])-1)
            randIndex2 = randint(0, len(solution[randCar2])-1)
            solution[randCar2].insert(randIndex2, solution[randCar1].pop(randIndex1))
    return solution

