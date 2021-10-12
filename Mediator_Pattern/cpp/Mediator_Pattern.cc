#include <iostream>
#include <stdio.h>
using namespace std;

class ChartRoom {
private:
    string no;
public:
    ChartRoom(const char* n):no(n) {}
public:
    void ShowMessage(string message) {
        cout << no << " " << message << endl;
    }
};

class User {
private:
    string name;
public:
    User(const char* n):name(n) {}
public:
    void SendMessage(string message) {
        char buf[1024];
        sprintf(buf, "[%s] %s", name.c_str(), message.c_str());
        ChartRoom("01").ShowMessage(buf);
    }
};

void test_mediator() {
    User u1("Lili");
    User u2("Tom");
    u1.SendMessage("hello");
    u2.SendMessage("world");
}
int main() {
    test_mediator();
    return 0;
}