

/*
    Weekly Coding Problem

    Week 25, 2022

    Given an array of integers that are out of order, determine the bounds of
   the window that must be sorted in order for the entire array to be sorted.

    For example, given `[3, 7, 5, 6, 9]`, you should return `(1,3)`.
*/

#include <iostream>
#include <vector>
#include <assert.h>

// types
typedef std::vector<int> vec;
struct Tests {
    vec input;
    vec result;
};

// prototypes
vec getSortBounds(vec&);

// globals 
vec v1{3, 7, 5, 6, 9}; 
vec v2{3, 7, 1, 4, 2, 9}; 
vec v1r{1, 3};
vec v2r{0, 4}; 
std::vector<Tests> t{{v1, v1r}, {v2, v2r}};


int main()
{

    return 0;
}

vec getSortBounds(vec &v)
{
  for (auto i : v) {
    
  }
}

