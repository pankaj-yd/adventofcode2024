#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <set>
using namespace std;

const vector<vector<int>> dirs = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
int srow = -1, scol = -1;

void guardStartPos(vector<string> &input){
    for(int i = 0; i < input.size(); i++){
        int idx = input[i].find('^');
        if(idx != string::npos){
            srow = i, scol = idx;
        }
    }
}

bool valid(vector<string> &input, int r, int c){
     if(r < 0 || r >= input.size() || c < 0 || c >= input[0].size()){
        return false;
    }
    return true;
}

int solve1(vector<string> &input){
    if(srow == -1){
        guardStartPos(input);
    }
    int r = srow, c = scol;

    int visited = 0;
    int dir = 0;
    while(true){
        if(!valid(input, r, c)){
            break;
        }

        // visit r, c. mark it visited
        if(input[r][c] != 'X'){
            visited++;
            input[r][c] = 'X';
        }
        
        int nr = r + dirs[dir][0], nc = c + dirs[dir][1];
        // turn if next idx 
        if(valid(input, nr, nc) && input[nr][nc] == '#'){
            dir = (dir+1)%4;
        } else {
            r = nr, c = nc;
        }
    }

    return visited;
}


// brute force
// first store the path
bool isLoop(vector<string> &input){
    set<vector<int>> visited;
    int r = srow, c = scol;
    int dir = 0;
    while(true){
        if(!valid(input, r, c)){
            return false;
        }

        if(visited.find({r, c, dirs[dir][0], dirs[dir][1]}) != visited.end()){
            return true;
        }

        // add current location to guardPath
        visited.insert({r, c, dirs[dir][0], dirs[dir][1]});

        int nr = r + dirs[dir][0], nc = c + dirs[dir][1];
        // turn if next idx is obstacle
        if(valid(input, nr, nc) && input[nr][nc] == '#'){
            dir = (dir+1)%4;
        } else {
            r = nr, c = nc;
        }
    }

}

int solve2(vector<string> &input){
    if(srow == -1){
        guardStartPos(input);
    }

    // find and store the guard path
    set<pair<int,int>> guardPath;
    int r = srow, c = scol;
    int dir = 0;
    while(true){
        if(!valid(input, r, c)){
            break;
        }
        
        // add current location to guardPath
        guardPath.insert({r, c});

        int nr = r + dirs[dir][0], nc = c + dirs[dir][1];
        // turn if next idx is obstacle
        if(valid(input, nr, nc) && input[nr][nc] == '#'){
            dir = (dir+1)%4;
        } else {
            r = nr, c = nc;
        }
    }

    // now we can try placing obstacles at every location in 
    // guard's path and check if it forms a loop
    int res = 0;
    for(auto it = guardPath.begin(); it != guardPath.end(); it++){
        // can not place obstacle at start
        // ignore it
        if(it->first == srow && it->second == scol){
            continue;
        }

        // place obstacle at x, y
        int x = it->first, y = it->second;
        input[x][y] = '#';
        if(isLoop(input)){
            res++;
        }
        input[x][y] = '.';

    }


    return res;

}

int main() {
    ifstream f("../testcases/6.txt");

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
    cout << "Part a: " << solve1(input) << endl;

    // part b
    cout << "Part b: " << solve2(input) << endl;
    
    // Expected answers:
    // Part a: 5129
    // Part b: 1888
    f.close();
    
    return 0;
}