#include <iostream>
#include <list>

int main()
{
    std::list<int> l;
    std::list<int>::iterator itr;

    l.push_front(1);
    itr = l.begin();
    l.push_front(2);
    l.pop_back();
    std::cout << *itr << '\n';
    return 0;
}