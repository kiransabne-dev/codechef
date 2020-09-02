// Example Input

// 2
// 10 4
// 10 8

// Example Output

// 0
// 1

// Explanation

// Example case 1: Chef attacks with power 4
// , Darth's health becomes 6. Chef attacks with power 2, Darth's health becomes 4. Chef attacks with power 1 and Darth's health becomes 3, but Chef's attack power becomes 0

// .

// Example case 2: Chef attacks with power 8
// , Darth's health becomes 2. Chef attacks with power 4, Darth's health becomes 0. Chef kills Darth.

#include <iostream>
using namespace std;

int main(){
   // cout << "Hello" << endl;
    int t;
    cin >> t;
    while(t--){
        int health, power;
        cin >> health;
        cin >> power;
       // cout << "health " << health << " and  power -> " << power << endl;
        // if(power >= health/2 + 2){
        //     cout << 1 << endl;
        // }else  {
        //     cout << 0 << endl;
        // }

        while(health != 0 && power != 0){
            if (health > power ){
                health = health - power;
            } else {
                health = 0;
                
            }
            power = power / 2;
        }
        if(health) {
            cout << 0 << endl;
        } else {
            cout << 1 << endl;
        }
    }
    return 0;
}

// 10 7
// 3 3
// 0 1

// 20 7
// 13 3
// 10 1
// 9 0

// 20 8
// 12 4
// 8 2
// 6 1
// 5 0

// 20 9
// 11 4
// 7 2
// 5 1
// 4 0

// 20 10
// 10 5
// 5 2
// 3 1
// 2 0

// 20 11
// 9 5
// 4 2
// 2 1
// 1 0

// 20 12
// 8 6
// 2 3
// 0 1

