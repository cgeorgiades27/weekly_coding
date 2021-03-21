// week 12, 2021 - @author Chris Georgiades

#include <iostream>

size_t consecOnes(const int &);

int main()
{
    int input;

    while (true)
    {
        std::cout << "\tEnter a number in decimal form (-999 to quit): ";
        std::cin >> input;
        if (input == -999)
            break;
        std::cout << "\tThe consecutive number of 1 bits is: " << consecOnes(input) << '\n';
    }
    return 0;
}

size_t consecOnes(const int &num)
{
    size_t ones = 0;
    size_t max = 0;

    for (size_t i = 0; i < sizeof(num) * 8; ++i)
    {
        if (num >> i & 1)
        {
            ++ones;
            // reset max
            if (ones > max)
                max = ones;
        }
        else // reset consecutive ones
            ones = 0;
    }
    return max;
}