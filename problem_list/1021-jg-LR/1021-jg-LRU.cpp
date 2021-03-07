

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
    }
    LRUCache() {} //default constructor

    //functions
    //set(key, value): sets key to value. If there are already n items in the cache and we are adding a new item, then it should also remove the least recently used item.
    void set(int key, std::string value){

    };
    //get(key): gets the value at key. If no such key exists, return null.
    std::string get(int key, std::string value)
    {
        return NULL;
    };
};

int main()
{
    std::cout << "Hello World!";
    LRUCache myLRU;              //calls default consturctor
    LRUCache paramatrizedLRU(3); //calss paramaterized constructor

    paramatrizedLRU.set(1, "thanos");
    paramatrizedLRU.set(2, "ironman");
    paramatrizedLRU.set(3, "thor");

    assert((paramatrizedLRU.LRU.size() == 3) && "Yet another way to add assert message");

    // Unit Tests

    return 0;
}
