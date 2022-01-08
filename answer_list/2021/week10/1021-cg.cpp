

#include <iostream>
#include <unordered_map>
#include <list>
#include <tuple>

typedef std::list<int>::iterator lItr;
std::unordered_map<char, lItr> umap;
std::list<int> priorityQ;

void set(char, int);
int get(char);
void print();
size_t maxSize = 4;

int main()
{

    set('a', 1);
    set('b', 2);
    set('c', 3);
    set('d', 4);
    set('e', 5);
    set('f', 6);
    set('g', 7);

    // std::cout << get('e') << '\n';
    // std::cout << get('a') << '\n';
    print();

    return 0;
}

void set(char key, int value)
{
    // cache miss
    if (umap.find(key) == umap.end())
    {
        // get rid of oldest if full
        if (priorityQ.size() == maxSize)
        {
            priorityQ.pop_back();
        }
        umap[key] = priorityQ.insert(priorityQ.begin(), value);
    }
    // cache hit
    else
    {
        priorityQ.erase(umap[key]);
        priorityQ.push_front(value);
        umap[key] = priorityQ.begin();
    }
}

int get(char key)
{
    if (umap.find(key) == umap.end())
        return NULL;
    return *(umap[key]);
}

void print()
{
    std::cout << "printing\n";
    for (lItr i = priorityQ.begin(); i != priorityQ.end(); ++i)
        std::cout << *i << '\n';
}