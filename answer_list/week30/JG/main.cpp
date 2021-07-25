
#include <iostream>
#include <string>
#include <assert.h>
#include <cmath>
#include <vector>
#include <algorithm>

/* This question was asked by Zillow.

You are given a `2-d` matrix where each cell represents number of coins in that cell. Assuming we start at `matrix[0][0]`, and can only move right or down, find the maximum number of coins you can collect by the bottom right corner.

For example, in this matrix:
```
0 3 1 1
2 0 0 4
1 5 3 1
```
The most we can collect is `0 + 2 + 1 + 5 + 3 + 1 = 12` coins.
*/

/**
 * Function that returns the sum of all the right moves from current position.
 *
 * @param M     Matrix
 * @param x     current x position 
 * @param y     current y position
 
 
 */
int addUpCurrentRow(const std::vector<std::vector<int>> &M, int x, int y)
{
  int col = M[0].size() - y;
  int sum = 0;
  while (col > 1)
  {
    if (y + 1 >= M[0].size())
      sum = sum + 0;

    else
      sum = sum + M[x][++y]; //r

    col--;
  }

  return sum;
}
/**
 * Function that returns the sum of all the down moves from current position.
 *
 * @param M     Matrix
 * @param x     current x position 
 * @param y     current y position
 
 
 */
int addUpCurrentCol(const std::vector<std::vector<int>> &M, int x, int y)
{
  int row = M.size() - x;
  int sum = 0;
  while (row > 1)
  {
    if (x + 1 >= M.size())
      sum = sum + 0;

    else
      sum = sum + M[++x][y]; //r

    row--;
  }

  return sum;
}

/**
 * Function that returns the maximim number of coins you can collect only making right or down turns.
 *
 * @param M     Matrix
 *  
 */
int getMaxCoins(const std::vector<std::vector<int>> &matrix)
{

  //get row count
  int remainingRightTurns = matrix.size();
  //get col count
  int remianingDownTurns = matrix[0].size();
  //to exit you will move a rows down and cols over
  int totalMoves = remainingRightTurns + remianingDownTurns - 2; // -2 because the start and finish are not counted
  //var that stores the amount of times we have moved right
  int CR = 0;
  //var that stores the amount of times we have moved down
  int CD = 0;
  //sum of the currrent choices
  int total = 0;

  //we always start at position [0][0], go ahead and add it to total
  total = total + matrix[0][0];

  while (totalMoves != 0)
  { //while you still have moves left

    if (addUpCurrentRow(matrix, CD, CR) > addUpCurrentCol(matrix, CD, CR)) // if sum of the rows is larger than the sum og the col
      total = total + matrix[CD][++CR];                                    //r
    else
      total = total + matrix[++CD][CR]; //d

    totalMoves--;
  }

  std::cout << "\nThe max number of coins is " << total << std::endl;

  return total;
}

void printMatrix(const std::vector<std::vector<int>> &matrix)
{
  for (auto row : matrix)
  {
    std::cout << std::endl;
    for (auto col : row)
    {
      std::cout << col << " ";
    }
  }
}

int main()
{

  //#Test Program
  //define the matrix
  std::vector<std::vector<int>> matrix = {{1, 2, 4, 9}, {1, 8, 6, 0}, {9, 6, 8, 5}, {0, 0, 2, 9}};
  printMatrix(matrix);
  assert((getMaxCoins(matrix) == 39) && "answer should be 39");

  //test 2
  matrix = {{0, 3, 1, 1}, {2, 0, 0, 4}, {1, 5, 3, 1}};
  printMatrix(matrix);
  assert((getMaxCoins(matrix) == 12) && "answer should be 12");

  //test 3
  matrix = {{1, 2, 1, 2}, {1, 1, 2, 2}, {2, 2, 1, 1}, {1, 4, 1, 2}};
  printMatrix(matrix);
  assert((getMaxCoins(matrix) == 13) && "answer should be 13");

  return 0;
}
