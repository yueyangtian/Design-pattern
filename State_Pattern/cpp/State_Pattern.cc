#include <iostream>
using namespace std;

class Vote {
public:
    virtual void DoAction() = 0;
};

class UnVote : public Vote {
public:
    void DoAction() {
        cout << "vote success" << endl;
    }
};

class Voted : public Vote {
public:
    void DoAction() {
        cout << "already voted" << endl;
    }
};

class Votor {
private:
    bool state;
    Vote* v;
public:
    Votor():state(false),v(new UnVote){}
public:
    void DoVote() {
        if (state) {
            v = new Voted;
        }
        v->DoAction();
        state = true;
    }
};

void test_state() {
    Votor v;
    v.DoVote();
    v.DoVote();
}

int main() {
    test_state();
    return 0;
}