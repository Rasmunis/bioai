x=[]
y=[]
D=[]
d=[]
q=[]
Q=0
m=0
n=0
t=0

def reader(filename):
    file=open(filename,'r')
    m,n,t = [int(i) for i in next(file).split()]
    
    for i in range(1,t+1):
        array=next(file).split()
        D.append(int(array[0]))
        Q=int(array[1])
    for line in file:
        array=[int(i) for i in line.split()]
        x.append(array[1])
        y.append(array[2])
        d.append(array[3])
        q.append(array[4])


reader('p01.txt')

