#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
#include <regex>
using namespace std;

// pattern match
// mul(X,Y)
long long solveMul(string &s){
    std::regex e("mul\\([0-9]?[0-9]?[0-9],[0-9]?[0-9]?[0-9]\\)");
    auto it = std::sregex_iterator(s.begin(), s.end(), e);
    auto end = std::sregex_iterator();
    long long res = 0;
    for(; it != end; ++it) {
        regex num("[0-9]?[0-9]?[0-9]");
        auto match = it->str();
        auto it2 = sregex_iterator(match.begin(), match.end(), num);
        auto end2 = sregex_iterator();
        long long currRes = 1;
        for(; it2 != end2; it2++){
            currRes *= stoi(it2->str());
        }
        
        res += currRes;
    }


    return res;
}

// pattern match
// mul(X,Y)
long long solveMulConditional(string &s){
    std::regex e("mul\\([0-9]?[0-9]?[0-9],[0-9]?[0-9]?[0-9]\\)|don't\\(\\).*?do\\(\\)|don't\\(\\)");
    auto it = std::sregex_iterator(s.begin(), s.end(), e);
    auto end = std::sregex_iterator();
    long long res = 0;
    for(; it != end; ++it) {
        auto match = it->str();
        // don't()...do()
        if(match.substr(0, 3) == "mul"){
            regex num("[0-9]?[0-9]?[0-9]");
            auto it2 = sregex_iterator(match.begin(), match.end(), num);
            auto end2 = sregex_iterator();
            long long currRes = 1;
            for(; it2 != end2; it2++){
                currRes *= stoi(it2->str());
            }
            res += currRes;
        } else if(match == "don't()"){
            break;
        }
    }

    return res;
}


int main() {
    ifstream f("../testcases/3.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    string input = "";
    while(getline(f, s)){
        input += s;
    }
    cout << "Part a: " << solveMul(input) << endl;
    cout << "Part b: " << solveMulConditional(input) << endl;

    // Expected answers:
    // Part a: 169021493
    // Part b: 111762583
    f.close();
    
    return 0;
}