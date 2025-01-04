#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <unordered_set>
#include <sstream>
using namespace std;

// For a key, it stores all vals which must come before it
unordered_map<int, unordered_set<int>> order;

bool isValidUpdate(vector<int> &update){
    for(int i = 0; i < update.size(); i++){
        for(int j = i+1; j < update.size(); j++){
            if(order[update[i]].find(update[j]) != order[update[i]].end()){
                return false;
            }
        }
    }
    return true;
}

int solve1(vector<vector<int>> &updates){
    int res = 0;
    for(vector<int> &update : updates){
        if(isValidUpdate(update)){
            int mid = ((int)update.size())/2;
            res += update[mid];
        }
    }
    return res;
}

// Just sort using manual comparator, you know
// one page number is less than other if it must
// come before another one
int solve2(vector<vector<int>> &updates){
    int res = 0;
    for(vector<int> &update : updates){
        if(isValidUpdate(update)){
            continue;
        }
        // make input valid
        auto cmp = [](int a, int b){
            if (order[b].find(a) != order[b].end()){
                return true;
            }
            return false;
        };
        sort(update.begin(), update.end(), cmp);
        int mid = ((int)update.size())/2;
        res += update[mid];
    }
    return res;
}

int main() {
    ifstream f("../testcases/5.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    // get order
    while(getline(f, s)){
        if(s == ""){
            break;
        }
        
        istringstream ss(s);
        int x, y;
        char ch;
        ss >> x >> ch >> y;
        order[y].insert(x);
    }

    // get updates
    vector<vector<int>> updates;
    while(getline(f, s)){
        istringstream ss(s);
        vector<int> update;
        int x;
        while(ss >> x){
            update.push_back(x);
            if(ss.peek() == ','){
                ss.ignore();
            }
        }
        updates.push_back(update);
    }

    cout << "Part a: " << solve1(updates) << endl;
    cout << "Part b: " << solve2(updates) << endl;

    // Expected answers:
    // Part a: 5391
    // Part b: 6142
    f.close();
    return 0;
}