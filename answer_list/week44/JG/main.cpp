
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/*Given a N by M matrix of numbers, print out the matrix in a clockwise spiral.

For example, given the following matrix:

[[1,  2,  3,  4,  5],
 [6,  7,  8,  9,  10],
 [11, 12, 13, 14, 15],
 [16, 17, 18, 19, 20]] */

/**
 * Function that puts the top row of the matrix into a vector and returns a smaller matrix. 
 *
 * @param M     Matrix
 * @param V     Vector to append the values to
 
 
 
 */
int Right(std::vector<std::vector<int> > &M, std::vector<int> &V)
{
  for (auto curr : M.at(0))
  {
    V.push_back(curr);
  }
  M.erase(M.begin());
  M.shrink_to_fit();
  return (0);
}
int Left(std::vector<std::vector<int> > &M, std::vector<int> &V)
{
  while (!M.back().empty())
  {
    V.push_back(M.back().back());
    M.back().pop_back();
  }
  M.pop_back();
  return (0);
}
int Down(std::vector<std::vector<int> > &M, std::vector<int> &V)
{
  for (int index = 0; index < M.size(); index++)
  {
    V.push_back(M.at(index).back());
    M.at(index).pop_back();
  }

  M.shrink_to_fit();
  return (0);
}
int Up(std::vector<std::vector<int> > &M, std::vector<int> &V)
{
  for (int index = M.size() - 1; index >= 0; index--)
  {
    V.push_back(M.at(index).front());
    M.at(index).erase(M.at(index).begin());
  }

  M.shrink_to_fit();
  return (0);
}

void printMatrix(const std::vector<std::vector<int> > &matrix)
{
  std::cout << "\nMatrix :";
  for (auto row : matrix)
  {
    std::cout << std::endl;
    for (auto col : row)
    {
      std::cout << col << " ";
    }
  }
}
void printVector(const std::vector<int> &matrix)
{
  std::cout << "\nSpiraled :" << std::endl;
  for (auto row : matrix)
  {
    std::cout << row << " ";
  }
}

int main()
{

  std::vector<std::vector<int> > matrix = {{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}};
  std::cout << "Original Matrix:" << std::endl;
  printMatrix(matrix);
  std::vector<int> spiraled;

  int i = 0;
  while (!matrix.empty())
  {
    if (i == 0)
      Right(matrix, spiraled);
    else if (i == 1)
      Down(matrix, spiraled);
    else if (i == 2)
      Left(matrix, spiraled);
    else if (i == 3)
      Up(matrix, spiraled);
    printMatrix(matrix);
    i++;
    if (i == 5)
      i = 0;
  }
  printMatrix(matrix);
  printVector(spiraled);

  return 0;
}
