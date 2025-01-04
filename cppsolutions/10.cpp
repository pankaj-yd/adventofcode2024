#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
#include <unordered_set>
using namespace std;

bool valid(vector<string> &input, int i, int j){
    return i >= 0 && i < input.size() && j >= 0 && j < input[0].size();
}

string getKey(int i, int j){
    return to_string(i) + "-" + to_string(j);
}

vector<vector<int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

void dfs(vector<string> &input, int i, int j, unordered_set<string> &score){
    if (!valid(input, i, j)){
        return;
    }
    // cout << "i: " << i << ", j: " << j << ", val:" << input[i][j] << endl;
    if(input[i][j] == '9'){
        score.insert(getKey(i, j));
        return;
    }

    for(vector<int> &dir : dirs){
        int nr = i + dir[0], nc = j + dir[1];

        if (valid(input, nr, nc) && input[nr][nc] == input[i][j] + 1 ){
            dfs(input, nr, nc, score);
        }
    }
}

int solve1(vector<string> &input){
    int totalScore = 0;
    for(int i = 0; i < input.size(); i++){
        for(int j = 0; j < input[0].size(); j++){
            if(input[i][j] == '0'){
                unordered_set<string> score;
                dfs(input, i, j, score);
                totalScore += score.size();
                // cout << "score: " << i << ","<< j << " = " << score.size() << endl;
            }
        }
    }
    return totalScore;
}

void dfs2(vector<string> &input, int i, int j, int &score){
    if (!valid(input, i, j)){
        return;
    }
    // cout << "i: " << i << ", j: " << j << ", val:" << input[i][j] << endl;
    if(input[i][j] == '9'){
        score++;
        return;
    }

    for(vector<int> &dir : dirs){
        int nr = i + dir[0], nc = j + dir[1];

        if (valid(input, nr, nc) && input[nr][nc] == input[i][j] + 1){
            dfs2(input, nr, nc, score);
        }
    }
}


int solve2(vector<string> &input){
    int totalScore = 0;
    for(int i = 0; i < input.size(); i++){
        for(int j = 0; j < input[0].size(); j++){
            if(input[i][j] == '0'){
                dfs2(input, i, j, totalScore);
                // cout << "score: " << i << ","<< j << " = " << score.size() << endl;
            }
        }
    }
    return totalScore;
}

int main() {
    ifstream f("../testcases/10.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    vector<string> input;
    while(getline(f, s)){
        input.push_back(s);
    }

    // part a
    long long partA = solve1(input);
    cout << "Part a: " << partA << endl;

    // part b
    long long partB = solve2(input);
    cout << "Part b: " << partB << endl;
    
    // Expected answers:
    // Part a: 557
    // Part b: 1062
    f.close();
    
    return 0;
}