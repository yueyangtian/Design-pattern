#include <iostream>
#include <vector>
#include <stdio.h>

using namespace std;
class Observer {
public:
    virtual void Notify(string) = 0;
};

class Object {
public:
    virtual void Attach(Observer*) = 0;
    virtual void NotifyAll() = 0; 
    virtual void SetStat(int) = 0;
protected:
    vector<Observer*> list;
    int stat;
};
class Object1 : public Object {
public:
    void Attach(Observer* ob) {
        list.push_back(ob);
    }
    void NotifyAll() {
        auto i = list.begin();
        char buf[1024];
        sprintf(buf, "from object 1 stat %d notified", stat);
        while (i != list.end()) {
            (*i)->Notify(buf);
            i++;
        }
    }
    void SetStat(int s) {
        stat = s;
        NotifyAll();
    }
};

class Observer1 : public Observer {
public:
    void Notify(string s) {
        cout << "observer1 " << s <<endl;
    }
};

class Observer2 : public Observer {
public:
    void Notify(string s) {
        cout << "observer2 " << s <<endl;
    }
};

void test_observer() {
    Object* obj = new Object1;
    Observer1 ob1;
    Observer2 ob2;
    Observer2 ob3;
    obj->Attach(&ob1);
    obj->Attach(&ob2);
    obj->Attach(&ob3);

    obj->SetStat(1);
}

int main() {
    test_observer();
    return 0;
}