#include<iostream>
#include <fstream>
#include<vector>
#include <string>
using namespace std;

int horizontal(vector<string> &s, int r, int c){
    int m = s.size(), n = s[r].size();
    int ans = 0;
    // forward
    if(c + 3 < n && s[r][c] == 'X' && s[r][c+1] == 'M' && s[r][c+2] == 'A' && s[r][c+3] == 'S'){
        ans++;
    }
    
    // backward
    if(c - 3 >= 0 && s[r][c] == 'X' && s[r][c-1] == 'M' && s[r][c-2] == 'A' && s[r][c-3] == 'S'){
        ans++;
    }

    return ans;
}

int vertical(vector<string> &s, int r, int c){
    int m = s.size(), n = s[r].size();
    int ans = 0;
    // down
    if(r + 3 < m && s[r][c] == 'X' && s[r+1][c] == 'M' && s[r+2][c] == 'A' && s[r+3][c] == 'S'){
        ans++;
    }
    
    // up
    if(r - 3 >= 0 && s[r][c] == 'X' && s[r-1][c] == 'M' && s[r-2][c] == 'A' && s[r-3][c] == 'S'){
        ans++;
    }

    return ans;
}

int diagonal(vector<string> &s, int r, int c){
    int m = s.size(), n = s[r].size();
    int ans = 0;
    // down, right: r++, c++
    if(r + 3 < m && c + 3 < n && s[r][c] == 'X' && s[r+1][c+1] == 'M' && s[r+2][c+2] == 'A' && s[r+3][c+3] == 'S'){
        ans++;
    }
    
    // down, left: r++, c--
    if(r + 3 < m && c - 3 >= 0 && s[r][c] == 'X' && s[r+1][c-1] == 'M' && s[r+2][c-2] == 'A' && s[r+3][c-3] == 'S'){
        ans++;
    }
    
    // up, right: r--, c++
    if(r - 3 >= 0 && c + 3 < n && s[r][c] == 'X' && s[r-1][c+1] == 'M' && s[r-2][c+2] == 'A' && s[r-3][c+3] == 'S'){
        ans++;
    }

    // up, left: r--, c--
    if(r - 3 >= 0 && c - 3 >= 0 && s[r][c] == 'X' && s[r-1][c-1] == 'M' && s[r-2][c-2] == 'A' && s[r-3][c-3] == 'S'){
        ans++;
    }

    return ans;
}


int count(vector<string> &s){
    int m = s.size(), n = s[0].size();

    int xmas = 0;
    for(int i = 0; i < m; i++){
        for(int j = 0; j < n; j++){
            xmas += horizontal(s, i, j) + vertical(s, i, j) + diagonal(s, i, j);
        }
    }
    return xmas;
}

int shapeMatch(vector<string> &s, int r, int c){
    int m = s.size(), n = s[0].size();
    int ans = 0;
    /*
    M.M
    .A.
    S.S
    */
   if(r+2 < m && c+2 < n && s[r][c] == 'M' && s[r][c+2] == 'M' && s[r+1][c+1] == 'A' && s[r+2][c] == 'S' && s[r+2][c+2] == 'S'){
    ans++;
   }

    /*
    M.S
    .A.
    M.S
    */
   if(r+2 < m && c+2 < n && s[r][c] == 'M' && s[r][c+2] == 'S' && s[r+1][c+1] == 'A' && s[r+2][c] == 'M' && s[r+2][c+2] == 'S'){
    ans++;
   }

   /*
    S.M
    .A.
    S.M
    */
   if(r+2 < m && c+2 < n && s[r][c] == 'S' && s[r][c+2] == 'M' && s[r+1][c+1] == 'A' && s[r+2][c] == 'S' && s[r+2][c+2] == 'M'){
    ans++;
   }

   /*
    S.S
    .A.
    M.M
    */
   if(r+2 < m && c+2 < n && s[r][c] == 'S' && s[r][c+2] == 'S' && s[r+1][c+1] == 'A' && s[r+2][c] == 'M' && s[r+2][c+2] == 'M'){
    ans++;
   }

   return ans;
}

int countShape(vector<string> &s){
    int m = s.size(), n = s[0].size();

    int xmas = 0;
    for(int i = 0; i < m; i++){
        for(int j = 0; j < n; j++){
            xmas += shapeMatch(s, i, j);
        }
    }
    return xmas;
}

int main() {
    ifstream f("../testcases/4.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    vector<string> input;
    while(getline(f, s)){
        input.push_back(s);
    }

    cout << "Part a: " << count(input) << endl;
    cout << "Part b: " << countShape(input) << endl;

    // Expected answers:
    // Part a: 2493
    // Part b: 1890
    f.close();
    
    return 0;
}