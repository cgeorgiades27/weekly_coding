
#include <iostream>
#include <vector>
#include <string>
#include <assert.h>
#include <cmath>
#include <memory>

/* ## Week 28, 2021

There are `N` prisoners standing in a circle, waiting to be executed. The executions are carried out starting with the kth person, and removing every successive kth person going clockwise until there is no one left.

Given `N` and `k`, write an algorithm to determine where a prisoner should stand in order to be the last survivor.

For example, if `N = 5` and `k = 2`, the order of executions would be `[2, 4, 1, 5, 3]`, so you should return `3`.

Bonus: Find an `O(log N)` solution if `k = 2`. */

struct Person
{
  /* data */
  int id;
  bool alive;
  std::shared_ptr<Person> next;
};

std::shared_ptr<Person> addPersonToLineUp(std::shared_ptr<Person> &P, int ID)
{

  auto newperson = std::make_shared<Person>();

  newperson->id = ID;
  newperson->alive = true;

  newperson->next = P->next;
  P->next = newperson;
  P = newperson;

  return P;
}

/**
 * Function that returns the best position to stand to be the last person standing.
 *
 * 
 * @param k     first to die i.e place to start 
 * @param n     amount of prisoners
 
 */

int getLastToDie(int k, int n)
{
  //make Head
  auto Head = std::make_shared<Person>();
  std::shared_ptr<Person> tmp;
  std::shared_ptr<Person> tmp2;
  int killCount = 0;

  for (int i = 0; i <= n; i++)
  { // will initilize a circular linked list of size N
    if (i == 0)
    {

      Head->id = 1;
      Head->alive = true;
    }
    else if (i == 1)
    {
      tmp = addPersonToLineUp(Head, i);
    }
    else
    {
      tmp = addPersonToLineUp(tmp, i);
    }
    if(n==1){  //base case
      Head->next=Head;
      tmp=Head;
      break;
    }
    else if (i == n - 1)
      tmp->next = Head;
  }
  if (k < n)
  {
    do // itterate until you get to, k, the starting position
    {
      tmp = tmp->next;

    } while (tmp->id != k);
  }
  tmp->alive = false; // kill first person
  killCount++;

  do //iterate k times
  {
    for (int i = 0; i < k; i++)
    {
      tmp = tmp->next;
      if (i < k - 1 || k == 1)
        tmp2 = tmp;

      if (tmp->alive == true && (i + 1 == k))
      {
        tmp->alive = false;
        killCount++;
      }
      else if (tmp->alive == false && (i <= k - 1))
        i--;
    }

  } while (killCount != n);

  return tmp2->id;
}

int main()
{
  int k;

  int n;
  std::cout << "Enter number of prisoners: " << std::endl;
  std::cin >> n;
  std::cout << "Enter starting position: " << std::endl;
  std::cin >> k;
  std::cout << "Best place to stand is: "<<getLastToDie(k, n) << std::endl;

  //# Test Program

  assert((getLastToDie(2, 5) == 3) && "3 should be last alive");

  assert((getLastToDie(1, 5) == 5) && "5 should be last alive");

  assert((getLastToDie(5, 5) == 2) && "2 should be last alive");

  assert((getLastToDie(2, 6) == 5) && "5 should be last alive");

  return 0;
}
