
grid = []
for y in range (0,99):
    temp = input()
    row = []
    
    for x in range (0,99):
        row.append(temp[x])

    grid.append(row)

treesVisible = 392

for y in range (1,98):
    for x in range (1,98):

        leftDistance = x
        topDistance = y
        rightDistance = 98-x
        bottomDistance = 98-y

        visible = True
        for z in range (1, leftDistance+1):
            if(grid[x][y]<=grid[x-z][y]):
                visible = False
                break
        
        if(visible):
            treesVisible=treesVisible+1
            continue


        visible = True
        for z in range (1, topDistance+1):
            if(grid[x][y]<=grid[x][y-z]):
                visible = False
                break
        
        if(visible):
            treesVisible=treesVisible+1
            continue


        visible = True
        for z in range (1, rightDistance+1):
            if(grid[x][y]<=grid[x+z][y]):
                visible = False
                break
        
        if(visible):
            treesVisible=treesVisible+1
            continue


        visible = True
        for z in range (1, bottomDistance+1):
            if(grid[x][y]<=grid[x][y+z]):
                visible = False
                break
        
        if(visible):
            treesVisible=treesVisible+1
            continue

print(treesVisible)
