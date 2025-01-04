#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
#include <unordered_set>
using namespace std;

struct Pos {
    int x, y;
    bool operator==(const Pos& other) const {
        return x == other.x && y == other.y;
    }
    size_t hash() const {
        return std::hash<int>()(x) ^ (std::hash<int>()(y) << 1);  // Combining x and y
    }
};


// Specialize std::hash for Pos to use the hash() member function
namespace std {
    template <>
    struct hash<Pos> {
        size_t operator()(const Pos& p) const {
            return p.hash();  // Call the member hash function
        }
    };
}

unordered_map<char, vector<Pos>> towers;
int rows = 0, cols = 0;


void preProcess(vector<string> input){
    // set rows, cols
    rows = input.size();
    cols = input[0].size();

    for (int i = 0; i < rows; i++){
        for(int j = 0; j < cols; j++){
            char ch = input[i][j];
            if ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9'){
                towers[ch].push_back({i, j});
            }
        }
    }
}

bool isValid(int x, int y){
    return x >= 0 && x < rows && y >= 0 && y < cols;
}

int solve1(){
    unordered_set<Pos> antinodes;
    for(const pair<char, vector<Pos>> &sameTowers : towers){
        const vector<Pos> &positions = sameTowers.second;
        for(int i = 0; i < positions.size(); i++){
            for(int j = i+1; j < positions.size(); j++){
                const Pos &p1 = positions[i], &p2 = positions[j];

                Pos node1 = {2 * p1.x - p2.x, 2 * p1.y - p2.y};
                Pos node2 = {2 * p2.x - p1.x, 2 * p2.y - p1.y};
                if(isValid(node1.x, node1.y)){
                    antinodes.insert(node1);
                }
                if(isValid(node2.x, node2.y)){
                    antinodes.insert(node2);
                }
            }
        }
    }

    return antinodes.size();
}


void storeAllAntinodes(unordered_set<Pos> &antinodes, Pos node, Pos diff){
    while(true){
        node = {node.x + diff.x, node.y + diff.y};
        if(!isValid(node.x, node.y)){
            break;
        }
        antinodes.insert(node);
    }
}

int solve2(){
    unordered_set<Pos> antinodes;
    for(const pair<char, vector<Pos>> &sameTowers : towers){
        const vector<Pos> &positions = sameTowers.second;
        for(int i = 0; i < positions.size(); i++){
            for(int j = i+1; j < positions.size(); j++){
                const Pos &p1 = positions[i], &p2 = positions[j];
                antinodes.insert(p1);
                antinodes.insert(p2);

                // calculate and add all antinodes too
                Pos node1 = p1;
                Pos diff1 = {p1.x - p2.x, p1.y - p2.y};
                storeAllAntinodes(antinodes, node1, diff1);
                
                Pos node2 = p2;
                Pos diff2 = {p2.x - p1.x, p2.y - p1.y};
                storeAllAntinodes(antinodes, node2, diff2);
            }
        }
    }

    return antinodes.size();
}

int main() {
    ifstream f("../testcases/8.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    vector<string> input;
    while(getline(f, s)){
        input.push_back(s);
    }

    // prePrcoess
    preProcess(input);
    // part a
    cout << "Part a: " << solve1() << endl;

    // part b
    cout << "Part b: " << solve2() << endl;
    
    // Expected answers:
    // Part a: 332
    // Part b: 1174
    f.close();
    
    return 0;
}