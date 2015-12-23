//initialise variables
var i, j;
var gridWidth= 10;
var gridHeight = 10;
var gridSize = gridWidth * gridHeight;
var neighbours;

//declare grid and arrays
var grid = new Array(gridHeight);
for (i=0; i<gridHeight; i++){
	grid[i] = new Array(gridWidth);
}	

var temp = new Array(gridHeight);
for (i=0; i<gridHeight; i++){
	grid[i] = new Array(gridWidth);
}	

//populate grid array
for (i=0; i<gridHeight; i++){
	for (j=0; j<gridWidth;j++){
		grid[i][j] = 0;
	}
}

//count neighbours
function countNeighbours(i, j){
	neighbours = 0;
	neighbours += grid[i+1][j];
	neighbours += grid[i][j+1];
	neighbours += grid[i-1][j];
	neighbours += grid[i][j-1];
	neighbours += grid[i+1][j+1];
	neighbours += grid[i-1][j-1];
	neighbours += grid[i-1][j+1];
	neighbours += grid[i+1][j-1];
	return neighbours;
}

//generate next generation
function nextGen() {
	for (i=1; i<=gridHeight; i++)
	{
		for (j=1; j<=gridWidth; j++)
		{
			neighbours = countNeighbours(i, j);
			if(grid[i][j] ==1){
				if(neighbours == 2 || neighbours == 3){
					temp[i][j] = 1;
				}
				else{
					temp[i][j] = 0;
				}
			}
			else if(grid[i][j]){
				if(neighbours == 3){
					temp[i][j] = 1;
				}
				else{
					temp[i][j] = 0;
				}
			}
		}
	}
}

//change grid from temp to grid
function updateGrid(){
	for (i=1; i<=gridHeight; i++)
	{
		for (j=1; j<=gridWidth; j++)
		{
			grid[i][j] == temp[i][j];
		}
	}
}
