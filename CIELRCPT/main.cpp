#include <bits/stdc++.h>

using namespace std;

int main(){
    cout << "hello " << endl;
    int testCases;
    cin >> testCases;
    while(testCases--){
        int price;
        cin >> price;
        
    }
    return 0;
}




// Sample Input
// 4
// 10
// 256
// 255
// 4096
// Sample Output
// 2
// 1
// 8
// 2
// Explanations
// In the first sample, examples of the menus whose total price is 10 are the following:
// 1+1+1+1+1+1+1+1+1+1 = 10 (10 menus)
// 1+1+1+1+1+1+1+1+2 = 10 (9 menus)
// 2+2+2+2+2 = 10 (5 menus)
// 2+4+4 = 10 (3 menus)
// 2+8 = 10 (2 menus)
// Here the minimum number of menus is 2.

// In the last sample, the optimal way is 2048+2048=4096 (2 menus). Note that there is no menu whose price is 4096.