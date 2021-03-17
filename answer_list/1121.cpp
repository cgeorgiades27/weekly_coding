

#include <iostream>
#include <map>
#include <stack>
#include <string>

class Solution
{
public:
    bool isValid(std::string);
    Solution()
    {
        // set up map
        map_[']'] = '[';
        map_[')'] = '(';
        map_['}'] = '{';
    }

private:
    std::map<char, char> map_;
    std::stack<char> stack_;
};

int main()
{
    std::string s1, s2, s3;
    Solution S;
    s1 = "()(){(())";
    // should return False
    std::cout << S.isValid(s1) << std::endl;

    s2 = "";
    // should return True
    std::cout << S.isValid(s2) << std::endl;

    s3 = "([{}])()";
    // should return True
    std::cout << S.isValid(s3) << std::endl;

    return 0;
}

bool Solution::isValid(std::string s)
{
    if (s.size() == 0)
        return true;

    for (size_t i = 0; i < s.size(); ++i)
    {
        if ((s[i] == ')' || s[i] == '}' || s[i] == ']') && !stack_.empty())
        {
            //printf("match: %c\n", s[i]);
            if (stack_.top() != map_[s[i]])
                return false;
            //printf("popping %c\n", stack_.top());
            stack_.pop();
            ++i;
        }
        else
        {
            //printf("pushing %c\n", s[i]);
            stack_.push(s[i]);
        }
    }
    if (!stack_.empty())
        return false;
    return true;
}