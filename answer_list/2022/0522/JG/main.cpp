
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/* ## Week 5, 2022

return the longest palindrom in a given stting


/*
 * Function that returns the longest palindrome substring in a given string.
 *
 * 
 * @param parseMe     string to search for lingest palindrome in
 
 */

std::string findLongestPalindrome(std::string parseMe)
{
  int counter = 0; // used to keep track ofwhat position from the right we are at
  std::string currentLongestPalindrome = "";
  std::string absoulteGlobalPalindrome = "";
  for (unsigned int index = 0; index < parseMe.length() - 1; index++)
  {
    currentLongestPalindrome = "";
    counter = 0; //start checking from the last char
    unsigned int tmpindex = index;
    do
    {
      counter++;
      if (parseMe.at(tmpindex) == parseMe.at(parseMe.length() - counter))
      {
        currentLongestPalindrome += parseMe.at(tmpindex);
        tmpindex++;
      }
    } while (parseMe.length() - counter > 0 && tmpindex < parseMe.length());

    if (currentLongestPalindrome.length() > absoulteGlobalPalindrome.length()) // if the current palindrome is longer than the absoulte palindrome make it the absolute palindrome
      absoulteGlobalPalindrome = currentLongestPalindrome;
  }
  std::cout << " Longest paindrome in, " << parseMe << " is " << absoulteGlobalPalindrome << std::endl;
  return absoulteGlobalPalindrome;
}

int main()
{

  //# Test Program

  std::string test = "adabbae";
  assert((findLongestPalindrome(test) == "abba") && "abba should be longest palindrome");

  std::string test2 = "adababae";

  assert((findLongestPalindrome(test2) == "ababa") && "ababa should be longest palindrome");

  std::string test3 = "cbbd";

  assert((findLongestPalindrome(test3) == "bb") && "bb should be longest palindrome");

  findLongestPalindrome("lollipop");

  return 0;
}
