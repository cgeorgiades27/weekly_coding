
/*
Weekly Coding Problem

Week 30, 2021

This question was asked by Zillow.

You are given a `2-d` matrix where each cell represents number of coins in that cell. Assuming we start at `matrix[0][0]`, and can only move right or down, find the maximum number of coins you can collect by the bottom right corner.

For example, in this matrix:

0 3 1 1
2 0 0 4
1 5 3 1

The most we can collect is `0 + 2 + 1 + 5 + 3 + 1 = 12` coins.

 */

#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <cstdlib>
#include <iomanip>

typedef std::vector<std::vector<int>> matrix;
int maxMatrixSum(matrix &, int &, int, int);
void printMatrix(matrix &, std::ostream &);

int main()
{
  matrix m0{{0, 2, 1}, {3, 0, 5}, {1, 0, 3}, {1, 4, 1}};
  matrix m1{{1, 7, 8}, {3, 6, 9}, {1, 10, 2}, {7, 4, 1}};
  matrix m2{{1, 1, 0, 1}, {1, 1, 2, 5}, {2, 1, 6, 4}, {0, 1, 1, 3}};
  matrix m3{{1, 1, 0, 1}, {1, 1, 2, 6}, {2, 1, 6, 4}, {0, 1, 0, 3}};
  int max = 0;

  std::vector<matrix> v{m0, m1, m2, m3};
  std::cout << "\nChoose one of the following matrices:\n";
  for (int i = 0; i < v.size(); ++i)
  {
    std::cout << "Matrix: " << i << "\n";
    printMatrix(v[i], std::cout);
  }

  int choice;
  std::cout << "Choose a matrix number (0-" << v.size() - 1 << "): ";
  std::cin >> choice;
  std::cout << "Testing Matrix: " << choice << '\n';
  printMatrix(v[choice], std::cout);
  std::cout << "Max sum: " << maxMatrixSum(v[choice], max, 0, 0) << '\n';
  return 0;
}

// Solution
int maxMatrixSum(matrix &m, int &max, int i, int j)
{
  if (i > m.size() - 1 || j > m[i].size() - 1)
    return 0;

  int iN = i + 1;
  int jN = j + 1;

  // Get max of subtrees
  int maxLeft = maxMatrixSum(m, max, i, jN);
  int maxRight = maxMatrixSum(m, max, iN, j);

  int maxOfAll = std::max(maxLeft, maxRight) + m[i][j];

  if (maxOfAll > max)
    max = maxOfAll;

  return maxOfAll;
}

// Print utility
void printMatrix(matrix &m, std::ostream &os)
{
  for (int i = 0; i < m[0].size(); ++i)
  {
    for (int j = 0; j < m.size(); ++j)
    {
      os << std::setw(3) << m[j][i];
      if (j == m.size() - 1)
        os << '\n';
      else
        os << ' ';
    }
  }
}
