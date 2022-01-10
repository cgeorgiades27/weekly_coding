// Your First C++ Program

#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <stack>
#include <cmath>
#include <list>
#include <algorithm>
#define NDEBUG

// This problem was asked by Facebook.

// Given a array of numbers representing the stock prices of a company in chronological order, write a function that calculates the maximum profit you could have made from buying and selling that stock once. You must buy before you can sell it.

// For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at 5 dollars and sell it at 10 dollars.
int largestProfit(std::vector<int> &mylist)
{
  int localMax = 0;
  int globalMax = 0;
  for (unsigned int index = 0; index < mylist.size() - 1; index++)
  {

    int *result = &*std::max_element(mylist.begin() + index, mylist.end());
    localMax = *result - (mylist[index]);
    if (localMax > globalMax)
    {
      globalMax = localMax;
    }
  }

  return globalMax;
}
int main()
{

  //# Test Program
  std::vector<int> stockPriceFB = {9, 11, 8, 5, 7, 10};
  assert((largestProfit(stockPriceFB) == 5) && "This return 5");

  std::vector<int> stockPriceAAPL = {5, 11, 8, 5, 100, 10};
  assert((largestProfit(stockPriceAAPL) == 95) && "This return 95");

  std::vector<int> stockPriceTSLA = {5, 4, 80, 5, 1, 100};
  assert((largestProfit(stockPriceTSLA) == 99) && "This return 99");

  std::vector<int> stockPricePTON = {35, 185, 80, 5, 1, 100};
  assert((largestProfit(stockPricePTON) == 150) && "This return 150");

  return 0;
}
