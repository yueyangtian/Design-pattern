#include<iostream>
using namespace std;

class Computer {
public:
    virtual void insertInSquarePort() = 0;
};

class Mac : public Computer {
public:
    void insertInSquarePort() {
        cout << "Mac insert square port" << endl;
    }
};

class Windows {
public:
    void insertInCriclePort() {
        cout << "Windows insert cricle port" << endl;
    }
};

class WindowsAdapter : public Computer {
public:
    void insertInSquarePort() {
        cout << "[adapter] ";
        win.insertInCriclePort();
    }
public:
    Windows win;
};

void test_adapter() {
    Mac m;
    m.insertInSquarePort();

    Windows w;
    w.insertInCriclePort();

    WindowsAdapter wa;
    wa.win = w;
    wa.insertInSquarePort();
}

int main() {
    test_adapter();
    return 0;
}