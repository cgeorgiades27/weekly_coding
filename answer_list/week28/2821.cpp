/*
######################
# Week 28, 2021 - CG #
######################

There are `N` prisoners standing in a circle, waiting to be executed. The executions are carried out starting with the kth person, and removing every successive kth person going clockwise until there is no one left.

Given N and k, write an algorithm to determine where a prisoner should stand in order to be the last survivor.

For example, if `N = 5` and `k = 2`, the order of executions would be `[2, 4, 1, 5, 3]`, so you should return `3`.

Bonus: Find an `O(log N)` solution if `k = 2`.
*/

#include <iostream>
#include <bitset>
#include <cstdlib>

int main(int argc, char *argv[])
{
  if (argc < 3)
    {
      std::cerr << "Too few args. USAGE " << argv[0] << " N K\n";
      std::exit(EXIT_FAILURE);
    }
  
  size_t N = std::atoi(argv[1]);
  size_t K = std::atoi(argv[2]);
  std::cout << "N = " << N << ", K = " << K << std::endl;

  typedef std::bitset<16> b_set;
  
  // prepare
  b_set bset;
  size_t last;
  size_t count = 0;
  
  for (size_t i = 0; i < N; ++i)
    bset.set(i);

  // basecase 1
  if (K >= N)
    {
      std::cerr << "K cannot be >= N, they won't be killed!\n"; 
      std::exit(EXIT_FAILURE);
    }
  
  // basecase 2
  if (K == 1)
    {
      last = N;
      goto END;
    }
  
  while (bset.count() > 1)
    {
      for (size_t i = 0; i < N; ++i)
	{
	  if (bset.test(i))
	   {
	     ++count;
	     if (count == K)
	       {
		 bset.reset(i);
		 count = 0;
	       }
	   }
	}
    }

  for (size_t i = 0; i < bset.size(); ++i)
    {
      if (bset.test(i))
	last = i + 1; // adjust for index offset
    }

 END:
  std::cout << "Best place to start: " << last << std::endl;
  
  return 0;
}
