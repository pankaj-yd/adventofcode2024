#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
using namespace std;

bool isIncreasingIdxGood(vector<int> report, int left, int idx, int right){
    // check left, idx, right is good
    bool iGood = true;

    // check left, idx
    if(left >= 0){
        int diff = report[idx] - report[left];
        if(diff > 3 || diff < 1){
            iGood = false;
        }
    }

    // check i, i+1
    if(iGood && right < report.size()){
        int diff = report[right] - report[idx];
        if(diff > 3 || diff < 1){
            iGood = false;
        }
    }
    return iGood;
}

bool isSafeIncreasing(vector<int> &report, bool tolerate){
    int n = report.size();
    for(int i = 1; i < n; i++){
        int diff = report[i] - report[i-1];
        if(1 <= diff && diff <= 3){
            continue;
        } else if( (diff > 3 || diff < 1) && !tolerate){
            return false;
        } else {
            // remove i-1 and check if i fits
            // check i-2, i, i+1 is good
            bool iGood = isIncreasingIdxGood(report, i-2, i, i+1);
            if(iGood){
                tolerate = !tolerate;
                i++;
                continue;
            }

            // remove i
            // check if i-2, i-1, i+1 is good
            bool iMinusOneGood = isIncreasingIdxGood(report, i-2, i-1, i+1);
            if(iMinusOneGood){
                tolerate = !tolerate;
                i++;
                continue;
            }
            return false;
        }
    }
    return true;
}

bool isDecreasingIdxGood(vector<int> report, int left, int idx, int right){
    // check left, idx, right is good
    bool iGood = true;

    // check left, idx
    if(left >= 0){
        int diff = report[left] - report[idx];
        if(diff > 3 || diff < 1){
            iGood = false;
        }
    }

    // check i, i+1
    if(iGood && right < report.size()){
        int diff = report[idx] - report[right];
        if(diff > 3 || diff < 1){
            iGood = false;
        }
    }
    return iGood;
}

bool isSafeDecreasing(vector<int> &report, bool tolerate){
    int n = report.size();
    for(int i = 1; i < n; i++){
        int diff = report[i-1] - report[i];
        if(1 <= diff && diff <= 3){
            continue;
        } else if( (diff > 3 || diff < 1) && !tolerate){
            return false;
        } else {
            // remove i-1 and check if i fits
            // check i-2, i, i+1 is good
            bool iGood = isDecreasingIdxGood(report, i-2, i, i+1);
            if(iGood){
                tolerate = !tolerate;
                i++;
                continue;
            }

            // remove i
            // check if i-2, i-1, i+1 is good
            bool iMinusOneGood = isDecreasingIdxGood(report, i-2, i-1, i+1);
            if(iMinusOneGood){
                tolerate = !tolerate;
                i++;
                continue;
            }
            return false;
        }
    }
    return true;
}

int countSafeReports(vector<vector<int>> &reports, bool tolerate){
    int numSafeReports = 0;
    for(vector<int> &report : reports){
        if(report.size() <= 1){
            numSafeReports++;
        } else if(isSafeIncreasing(report, tolerate) || isSafeDecreasing(report, tolerate)){
            numSafeReports++;
        }
    }

    return numSafeReports;
}

int main() {
    ifstream f("../testcases/2.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    vector<vector<int>> reports;
    while(getline(f, s)){
        vector<int> report;
        stringstream ss(s);
        int x;
        while(ss >> x){
            report.push_back(x);
        }
        reports.push_back(report);
    }
    cout << "Part a: " << countSafeReports(reports, false) << endl;
    cout << "Part b: " << countSafeReports(reports, true) << endl;

    // Expected answers:
    // Part a: 585
    // Part b: 626
    f.close();
    
    return 0;
}