// Explanation

// Example case 1: Chef and Rick generate the optimal integers 3
// and 5 respectively. Each of them has 1

// digit, so Rick cheats and wins the game.

// Example case 2: Chef and Rick could generate e.g. 6877
// and 99 respectively. Since Rick's integer has 2 digits and Chef cannot generate an integer with less than 4

// digits, Rick wins.

// Example case 3: Chef and Rick could generate e.g. 86
// and 888 respectively. Chef's integer has 2 digits and Rick cannot generate an integer with less than 3 digits, so Chef wins.
// 1, 9, 18, 27, 36, ...
// 1, 9, 99, 999, 9999, 99999, ...

// Example Input

// 3
// 3 5
// 28 18
// 14 24

// Example Output

// 1 1
// 1 2
// 0 2

#include <iostream>
using namespace std;

int main(){
    int t;
    cin >> t;
    while(t--){
        int c, r, dc, dr;
        cin >> c >> r;
        if(c % 9 == 0){
            dc = c / 9;
        } else {
            dc = (c / 9) + 1;
        }

        if(r % 9 == 0){
            dr = r / 9;
        } else {
            dr = (r / 9) + 1;
        }

        if(dc == dr || dr < dc){
            cout << 1 << " " << dr << endl;
        } else if(dc < dr){
            cout << 0 << " " << dc << endl;
        }
    }

    return 0;
}