#include <iostream>
using namespace std;

int main() {
	// your code goes here
    cout << "text" << endl;
    int N;
    int noOfLucky = 0;
    int noOfOdd = 0;
    cin >> N;
    for(int i = 0; i < N; i++){
        int weapons;
        cin >> weapons;
        cout << "weapons " << weapons << endl;
        if(weapons % 2 == 0){
            cout << "even" << endl;
            noOfLucky++;
        } else {
            cout << "odd" << endl;
            noOfOdd++;
        }
        cout << "even " << noOfLucky << " odd " << noOfOdd << endl;
    }

    if (noOfLucky > noOfOdd){
        cout << "READY FOR BATTLE" << endl;
    } else {
        cout << "NOT READY" << endl;
    }

	return 0;
}
