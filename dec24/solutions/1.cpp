#include<iostream>
#include <fstream>
#include<vector>
#include <string>
using namespace std;

int checkDiff(vector<int> &list1, vector<int> &list2){
    sort(list1.begin(), list1.end());
    sort(list2.begin(), list2.end());
    int ans = 0;
    for(int i = 0; i < list1.size(); i++){
        ans += abs(list1[i] - list2[i]);
    }

    return ans;
}

int similarityScore(vector<int> &list1, vector<int> &list2){
    unordered_map<int, int> freq;
    for(int num : list2){
        freq[num]++;
    }

    int score = 0;
    for(int num : list1){
        score += num * freq[num];
    }
    return score;
}

int main() {
    ifstream f("../testcases/1.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    int x, y;
    vector<int> list1, list2;
    while(f >> x && f >> y){
        list1.push_back(x);
        list2.push_back(y);
    }
    // part a
    int ans = checkDiff(list1, list2);
    cout << "Part a: " << ans << endl;

    // part b
    int score = similarityScore(list1, list2);
    cout << "Part b: " << score << endl;
    
    // Expected answers:
    // Part a: 1223326
    // Part b: 21070419
    f.close();
    
    return 0;
}