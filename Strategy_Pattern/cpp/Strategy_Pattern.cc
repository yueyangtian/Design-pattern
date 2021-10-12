#include <iostream>
using namespace std;

class strategy {
public:
    virtual int Opreation(int, int) = 0;
};

class AddOpreation : public strategy {
public:
    int Opreation(int a, int b) {
        return a + b;
    }
};

class MutilOpreation : public strategy {
public:
    int Opreation(int a, int b) {
        return a * b;
    }
};

class Context {
private:
    strategy& Option;
public:
    Context(strategy& s):Option(s) {};
public:
    int Exec(int a, int b) {
        return Option.Opreation(a, b);
    }
};


void test_strategy() {
    AddOpreation ap;
    MutilOpreation mp;
    Context add(ap);
    Context mutil(mp);

    cout << "1 + 2 = " << add.Exec(1, 2) << endl;
    cout << "1 * 2 = " << mutil.Exec(1, 2) << endl;
}

int main() {
    test_strategy();
    return 0;
}