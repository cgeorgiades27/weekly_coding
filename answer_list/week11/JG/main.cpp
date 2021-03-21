// Your First C++ Program

#include <iostream>
#include <vector>
#include<string>
#include<assert.h>


  bool isValid(std::string s){
      //# Fill this in. 
  return false;
  }
   

int main() {
    
std::string s;

//# Test Program
s = "()(){(())" ;
//# should return False
assert((isValid(s)==false) && "This should be false");
   

s = "";
//# should return True
assert((isValid(s)==true) && "This should be true");

s = "([{}])()";
//# should return True
assert((isValid(s)==false) && "This should be true");

    return 0;
}
