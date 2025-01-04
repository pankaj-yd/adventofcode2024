#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
using namespace std;

bool canCalibrate(vector<long long> &res, vector<vector<long long>> &operands, long long i, long long j, long long curr){
    if(curr == res[i] && j == operands[i].size()){
        return true;
    }

    if(i >= res.size() ||  j >= operands[i].size()){
        return false;
    }

    return canCalibrate(res, operands, i, j+1, curr + operands[i][j]) 
    || canCalibrate(res, operands, i, j+1, curr * operands[i][j]);

}


long long solve1(vector<long long> &res, vector<vector<long long>> &operands){
    long long ans = 0;
    for(long long i = 0; i < res.size(); i++){
        if (operands[i].size() > 0 && canCalibrate(res, operands, i, 1, operands[i][0])){
            ans += res[i];
        }
    }
    return ans;
}


bool canCalibrate2(vector<long long> &res, vector<vector<long long>> &operands, long long i, long long j, long long curr){
    if(curr == res[i] && j == operands[i].size()){
        return true;
    }

    if(i >= res.size() ||  j >= operands[i].size()){
        return false;
    }

    return canCalibrate2(res, operands, i, j+1, curr + operands[i][j]) 
    || canCalibrate2(res, operands, i, j+1, curr * operands[i][j]) 
    || canCalibrate2(res, operands, i, j+1, stol(to_string(curr) + to_string(operands[i][j])));

}

long long solve2(vector<long long> &res, vector<vector<long long>> &operands){
    long long ans = 0;
    for(long long i = 0; i < res.size(); i++){
        if (operands[i].size() > 0 && canCalibrate2(res, operands, i, 1, operands[i][0])){
            ans += res[i];
        }
    }
    return ans;
}

int main() {
    ifstream f("../testcases/7.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    vector<long long> res;
    vector<vector<long long>> operands;
    while(getline(f, s)){
        istringstream ss(s);
        long long ans;
        char ch;
        ss >> ans >> ch;
        res.push_back(ans);
        vector<long long> curr;
        while(ss >> ans){
            curr.push_back(ans);
        }
        operands.push_back(curr);
    }
    // part a
    cout << "Part a: " << solve1(res, operands) << endl;

    // part b
    cout << "Part b: " << solve2(res, operands) << endl;
    
    // Expected answers:
    // Part a: 12553187650171
    // Part b: 96779702119491
    f.close();
    
    return 0;
}