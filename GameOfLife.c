#include <stdio.h>
        
void printgrid (int i, int j,int grid[10][10])
{
    for(i=0; i<10; i++)
    {
        for(j=0; j<10; j++)
        {
            printf("%d ", grid[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}
int main()
{
    int grid[10][10] =
    {
        0,0,0,0,0,0,0,0,0,0,
        0,0,1,0,0,0,0,0,0,0,
        1,0,1,0,0,0,0,0,0,0,
        0,1,1,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
    };
    int temp[10][10] =
    {
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,
    };

    int i, j, alive, dead, counter, timer;
    printgrid (i,j,grid);
    alive=1;
    for(timer=0; timer<20; timer++)
    {
		for(counter=0; counter<10; counter++)
		{
			for(i=0; i<10; i++)
			{
				for(j=0; j<10; j++)
				{
				int neighbours;
				neighbours=grid[i+1][j] + grid[i][j+1] + grid[i-1][j] + grid[i][j-1] + grid[i+1][j+1] +
							grid[i-1][j-1] + grid[i-1][j+1] + grid[i+1][j-1];
					if(grid[i][j] == 1)
					{
						if(neighbours == 2 || neighbours == 3)
						{
							temp[i][j] = 1;
						}
						else{
							temp[i][j]=0;
						}
					}
					else if(grid[i][j] == 0)
					{
						if(neighbours == 3)
						{
							temp[i][j] = 1;
						}
						else{
							temp[i][j]=0;
						}
					}
				}
			}
		}
		 for(i=0; i<10; i++)
			{
				for(j=0; j<10; j++)
				{
					grid[i][j]=temp[i][j];
				}
			}
		printgrid (i,j,grid);
    }           
    return 0;
}
