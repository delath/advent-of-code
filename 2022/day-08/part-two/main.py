SIZE = 99
grid = []
for y in range (0,SIZE):
    temp = input()
    row = []
    
    for x in range (0,SIZE):
        row.append(temp[x])

    grid.append(row)

maxScenicScore = 0

for y in range (1,SIZE-1):
    for x in range (1,SIZE-1):

        leftDistance = x
        topDistance = y
        rightDistance = (SIZE-1)-x
        bottomDistance = (SIZE-1)-y

        viewDistanceA = 0
        viewDistanceB = 0
        viewDistanceC = 0
        viewDistanceD = 0

        for z in range (1, leftDistance+1):
            if(grid[x][y]<=grid[x-z][y] or z==leftDistance):
                viewDistanceA = z
                break

        for z in range (1, topDistance+1):
            if(grid[x][y]<=grid[x][y-z] or z==topDistance):
                viewDistanceB = z
                break

        for z in range (1, rightDistance+1):
            if(grid[x][y]<=grid[x+z][y] or z==rightDistance):
                viewDistanceC = z
                break

        for z in range (1, bottomDistance+1):
            if(grid[x][y]<=grid[x][y+z] or z==bottomDistance):
                viewDistanceD = z
                break
        
        currentScenicScore = viewDistanceA * viewDistanceB * viewDistanceC * viewDistanceD

        if(currentScenicScore>maxScenicScore):
            maxScenicScore=currentScenicScore

print(maxScenicScore)