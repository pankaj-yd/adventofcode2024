#include<iostream>
#include <fstream>
#include<vector>
#include <string>
#include <sstream>
#include <unordered_set>
using namespace std;


void addCheckSum(long long &res, long long fileId, long long start, long long fileSize){
    //  cout << "fileId: " << fileId << ", fileSize: " << fileSize << " , startIdx: " << start << endl; 
    // start * fileId + (start+1) * fileId + .... + (start + fileSize - 1) * fileId
    long long n = start + fileSize - 1;
    res += ((n * (n + 1))/2  - (start * (start - 1))/2) * fileId;

    // for(int i = start; i < start + fileSize; i++){
    //     cout << i << "*" << fileId << endl;
    // }
}

long long solve1(vector<int> input){
    long long n = input.size();
    // odd indices are free spaces
    long long i = 0; // indicate free space
    long long j = n - 1; // fileId from back

    if( !n%2){
        // n is even
        // last idx is freespace
        j--;
    }

    // expanded index
    long long idx = 0;

    // stores checksum
    long long res = 0;
    while(i < j){
        // even input index i
        // means a file
        // add checksum for this
        if(i%2 == 0){
            long long fileSize = input[i];
            addCheckSum(res, i/2, idx, fileSize);

            // update expanded index
            idx += fileSize;
            i++;
            continue;
        }

        // odd input index i, means free space
        // we'll put file from back in free space
        long long freeSpace = input[i];
        long long fileSize = input[j];

        // fit as much as we can
        // update checkSum and expanded index
        long long fit = min(freeSpace, fileSize);
        addCheckSum(res, j/2, idx, fit);
        idx += fit;

        // decrease from freespace and filesize
        input[i] -= fit, input[j] -= fit;
        
        if(input[i] == 0){
            i++;
        }
        if(input[j] == 0){
            j -= 2;
        }
    }

    // add remaining files to checkSum
    while(i < n){
        long long fileSize = input[i];
        addCheckSum(res,i/2, idx, fileSize);

        // update expanded index
        idx += fileSize;
        i += 2;
    }

    return res;
}

long long solve2(vector<int> input){
    long long n = input.size();

    vector<vector<int>> freeSpaces;
    int expandedIndex = 0;
    for(int i = 1; i < n; i += 2){
        expandedIndex += input[i-1];
        freeSpaces.push_back({i, input[i], expandedIndex});
        expandedIndex += input[i];
    }


    // odd indices are free spaces
    long long j = n - 1; // fileId from back

    if( n%2 == 0 ){
        // n is even
        // last idx is freespace
        j--;
    }

    // stores checksum for movable files
    long long res = 0;

    // 00992111777.44.333....5555.6666.....8888..
    while(j >= 0){
        // we'll put file from back in free space
        long long fileSize = input[j];

        for(int k = 0; k < freeSpaces.size(); k++){
            if (freeSpaces[k][0] > j){
                break;
            }

            // input idx, capcity, expanded start idx
            if (freeSpaces[k][1] >= fileSize){
                addCheckSum(res, j/2, freeSpaces[k][2], fileSize);

                input[j] = -input[j];
                // decrease from freespace and filesize
                freeSpaces[k][1] -= fileSize;
                freeSpaces[k][2] += fileSize;
                break;
            }
        }
        
        // whether we moved it successfully or not
        // move to next file
        j -= 2;
    }

    // add remaining files to checkSum
    // cout << "Adding unmoved" << endl;
    long long i = 0; // indicate file
    // expanded index
    long long idx = 0;
    while(i < n){
        if(i != 0){
            idx += input[i-1];
        }

        if(input[i] < 0){
            idx += -input[i];
            i += 2;
            continue;
        }

        long long fileSize = input[i];
        addCheckSum(res,i/2, idx, fileSize);

        // update expanded index
        idx += fileSize;
        i += 2;
    }

    return res;
}

int main() {
    ifstream f("../testcases/9.txt");

    if (!f) {
        cerr << "Error opening the file!" << endl;
        return 1;
    }

    string s;
    getline(f, s);
    
    vector<int> input;
    for(char ch : s){
        input.push_back(ch - '0');
    }

    // part a
    long long partA = solve1(input);
    cout << "Part a: " << partA << endl;

    // part b
    long long partB = solve2(input);
    cout << "Part b: " << partB << endl;
    
    // Expected answers:
    // Part a: 6262891638328
    // Part b: 6287317016845
    f.close();
    
    return 0;
}