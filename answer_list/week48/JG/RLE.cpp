
#include <iostream>
#include <string>
#include <assert.h>
#include <cmath>
#include <vector>
#include <algorithm>

void encode(std::string str)
{

    for (int count = 1, index = 0; index < str.length(); index++, count = 1)
    {
        while (str[index] == str[index + 1])
        {
            if (str[index] != str.back() - 1)
            {
                count++;
                index++;
            }
        }
        std::cout << count << str[index];
    }
}

int main()
{
    std::string str = "AAAABBBCCDAA";
    encode(str);
    return 0;
}
