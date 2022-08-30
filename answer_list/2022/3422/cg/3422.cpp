

/* 
    Weekly Coding Problem
    
    Week 34, 2022
    
    You 2 integers n and m representing an n by m grid, determine the number of ways you can
    get from the top-left to the bottom-right of the matrix y going only right or down.
    
    Example:
    `n = 2, m = 2`
    
    This should return 2, since the only possible routes are:
    Right, down
    Down, right.
    
    Here's the signature:

    def num_ways(n, m):
      # Fill this in.
    
    print num_ways(2, 2)
    # 2

*/


#include <iostream>
#include <vector>

typedef std::vector<std::vector<int>> matrix;

int num_ways(int, int);
void recurse_mtx(matrix&, int&);

int main()
{

  return 0;
}

int num_ways(int n, int m)
{
  
  std::vector<int> in(n,0);
  matrix mtx(m,in);

  
}


void recurse_mtx(matrix &m, int &count)
{
  for (auto i: m)
    {
      for (auto j: i)
        {
          
        }
    }
}