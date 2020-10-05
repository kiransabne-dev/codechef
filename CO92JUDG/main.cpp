// Input:

// 3
// 5
// 3 1 3 3 4
// 1 6 2 5 3
// 5
// 1 6 2 5 3
// 3 1 3 3 4
// 3
// 4 1 3
// 2 2 7
#include <iostream>
#include <vector>
#include <numeric>
#include <algorithm>

using namespace std;

int main(){
    int t;
    cin >> t;
    while(t--){
        int noOfMatches;
        cin >> noOfMatches;
        vector <int> a(noOfMatches);
        vector <int> b(noOfMatches);
        //cout << "scoreA[1] " << scoreA[1] << endl;
        for(int i = 0; i < noOfMatches; ++i){
            cin >> a[i];
        }
        for(int i = 0; i < noOfMatches; ++i){
            cin >> b[i];
        }
        int x = accumulate(a.begin(), a.end(), 0) - *max_element(a.begin(), a.end());
        int y = accumulate(b.begin(), b.end(), 0) - *max_element(b.begin(), b.end());
        if (x < y) cout << "Alice" << endl;
        else if(x > y) cout << "Bob" << endl;
        else cout << "Draw" << endl;

    }

    return 0;
}