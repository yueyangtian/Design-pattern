#include <iostream>
using namespace std;

class HardDrive {
public:
    void load() {
        cout << "hard drive loading" << endl;
    }
};

class Memory {
public:
    void read() {
        cout << "memory reading" << endl;
    }
};

class CPU {
public:
    void execute() {
        cout << "cpu executing" << endl;
    }
};

class Computer {
public:
    void start() {
        hardDrive.load();
        cpu.execute();
        memory.read();
    }
public:
    HardDrive hardDrive;
    Memory memory;
    CPU cpu;
};

void test_Facade() {
    Computer c;
    c.start();
}
int main() {
    test_Facade();
    return 0;
}