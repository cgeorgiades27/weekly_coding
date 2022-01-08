// week 9, 2021

#include <iostream>
#include <vector>

typedef std::vector<int> vec;
int intersection(vec &, vec &, vec &);

int main()
{
    vec v1 = {1, 2, 3, 4};
    vec v2 = {2, 4, 6, 8};
    vec v3 = {3, 4, 5};

    printf("%d\n", intersection(v1, v2, v3));
    return 0;
}

int intersection(vec &v1, vec &v2, vec &v3)
{
    size_t maxNumSize = 100;
    vec arr(maxNumSize, 0);

    vec v;
    v.insert(v.begin(), v1.begin(), v1.end());
    v.insert(v.end(), v2.begin(), v2.end());
    v.insert(v.end(), v3.begin(), v3.end());

    for (size_t i = 0; i < v.size(); ++i)
        ++arr[v[i]];

    for (size_t i = 0; i < arr.size(); ++i)
    {
        if (arr[i] == 3)
            return i;
    }
    return 0;
}