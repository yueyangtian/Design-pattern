#include<iostream>
using namespace std;
class printer {
public:
    virtual void printFile() = 0;
};

class compter {
public:
    virtual void print() = 0;
    virtual void setPrinter(printer* p) = 0;
};

class windows : public compter {
public:
    void print() {
        cout << "Print request for windows" << endl;
        p->printFile();
    }
    void setPrinter(printer* p) {
        this->p = p;
    }
private:
    printer* p;
};

class linux : public compter {
public:
    void print() {
        cout << "Print request for linux" << endl;
        p->printFile();
    }
    void setPrinter(printer* p) {
        this->p = p;
    }
private:
    printer* p;
};

class hp_windows : public printer {
public:
    void printFile() {
        cout << "OS:windows Printing by a HP Printer" << endl;
    }
};
class hp_linux : public printer {
public:
    void printFile() {
        cout << "OS:linux Printing by a HP Printer" << endl;
    }
};
class suf_windows : public printer{
public:
    void printFile() {
        cout << "OS:windows Printing by a Suf Printer" << endl;
    }
};

void test_bridge() {
    hp_windows hp_Printer;
    suf_windows suf_windows;
    
    windows win;
    linux linus;

    linus.setPrinter(&hp_Printer);
    win.setPrinter(&suf_windows);

    linus.print();
    win.print();
}
int main() {
    test_bridge();
    return 0;
}