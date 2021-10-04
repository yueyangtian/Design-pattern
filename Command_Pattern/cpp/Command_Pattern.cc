#include <iostream>
#include <vector>
using namespace std;

class Command {
public:
    virtual void Execute() = 0;
};

class Light : public Command {
public:
    void Execute() {
        TurnOn();
    }
    void TurnOn() {
        cout << "Light turn on" << endl;
    }
};

class TV : public Command {
public:
    void Execute() {
        TurnOn();
    }
    void TurnOn() {
        cout << "tv turn on" << endl;
    }
};

class Switcher {
private:
    vector<Command*> CommandList;
public:
    void ExecandStore(Command* c) {
        c->Execute();
        CommandList.push_back(c);
    }
};

void test_command() {
    Switcher s;
    TV t;
    Light l;
    s.ExecandStore(dynamic_cast<Command*>(&t));
    s.ExecandStore(dynamic_cast<Command*>(&l));
}
int main() {
    test_command();
    return 0;
}