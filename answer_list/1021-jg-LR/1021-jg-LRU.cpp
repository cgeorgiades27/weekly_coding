

#include <iostream>
#include <cassert>
#include <map>
#include <vector>
/* This problem was asked by Google.

Implement an LRU (Least Recently Used) cache. 


Each operation should run in O(1) time. */

class LRUCache
{

    //attributes

public:
    std::vector<std::pair<int, std::string>> LRU;
    //It should be able to be initialized with a cache size n
    LRUCache(int size) //paramaterized constructor
    {
        std::vector<std::pair<int, std::string>> LRU(size);
        this->LRU = LRU;
    }
    LRUCache() {} //default constructor

    //functions
    //set(key, value): sets key to value. If there are already n items in the cache and we are adding a new item, then it should also remove the least recently used item.
    void set(int key, std::string value)
    {
        bool hit = false;
        std::pair<int, std::string> inserMe(key, value);
        // itterate through LRU
        for (long unsigned int index = 0; index < LRU.size(); index++)
        {
            if (LRU[index].first == key)
            {
                // if hit
                std::cout << "hit" << std::endl;
                LRU.erase(LRU.begin() + index);   //erase the pair
                LRU.insert(LRU.begin(), inserMe); //add to begining
                hit = true;
            }
        }
        if (hit == false)
        {
            // if not a hit insert at front

            LRU.insert(LRU.begin(), inserMe);
            LRU.pop_back();
        }
    }
    //get(key): gets the value at key. If no such key exists, return null.
    std::string get(int key)
    {
        for (long unsigned int index = 0; index < LRU.size(); index++)
        {
            if (LRU[index].first == key)
            {
                // if hit
                return (LRU[index].second);
            }
        }
        return NULL;
    }
};

int main()
{

    LRUCache myLRU;              //calls default consturctor
    LRUCache paramatrizedLRU(3); //calss paramaterized constructor
    assert((paramatrizedLRU.LRU.size() == 3) && "There should be 3 entries ");

    // Unit Tests//
    paramatrizedLRU.set(1, "thanos");
    paramatrizedLRU.set(2, "ironman");
    paramatrizedLRU.set(3, "thor");

    assert((paramatrizedLRU.LRU.size() == 3) && "There should be 3 entries ");
    //hit middle value
    paramatrizedLRU.set(1, "thanos");
    assert((paramatrizedLRU.LRU[0].first == 1 && paramatrizedLRU.LRU[1].first == 3 && paramatrizedLRU.LRU[2].first == 2) && "order should be thanos,thor,ironman  ");
    // hit last value
    paramatrizedLRU.set(2, "ironman");
    assert((paramatrizedLRU.LRU[0].first == 2 && paramatrizedLRU.LRU[1].first == 1 && paramatrizedLRU.LRU[2].first == 3) && "order should be ironman,thanos,thor  ");
    //miss
    paramatrizedLRU.set(4, "blackPanther");
    assert((paramatrizedLRU.LRU[0].first == 4 && paramatrizedLRU.LRU[1].first == 2 && paramatrizedLRU.LRU[2].first == 1) && "order should be blackPanther,ironman,thanos  ");
    std::cout << "Value: " << paramatrizedLRU.get(2) << std::endl;
    return 0;
}
