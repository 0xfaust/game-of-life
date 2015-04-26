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

void read_grid (char* fname, int grid[10][10])
{
	FILE* map;
	char comma;
	int i, j, k;
	
	map = fopen(fname, "r");
	if(map == NULL)
	{
		printf("Opening file failed\n");
	}
	else
	{
		while(!feof(map))
		{
			for(k=0; k<10; k++)
			{
				for(i=0; i<10; i++)w
				{
					for(j=0; j<10; j++)
					{
						fscanf(map, "%d%c", &grid[i][j], &comma);
					}
				}
			}
		}
	}
	fclose(map);
}		
			
int main()
{
    int grid[10][10];
    int temp[10][10];
    int i, j, alive, dead, counter, timer;
	read_grid("grid.txt", grid);
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
