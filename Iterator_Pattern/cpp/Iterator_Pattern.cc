#include <iostream>
#include <vector>
using namespace std;

class IntIterator {
public:
    virtual IntIterator* Next() = 0;
    virtual IntIterator* Frist() = 0;
    virtual int Current() = 0;
};

class IntIter : public IntIterator {
public:
    IntIter(vector<int>& _v):v(_v),idx(0) {}
private:
    vector<int>& v;
    int idx;
public:
    IntIterator* Next() {
        idx++;
        return this;
    }
    IntIterator* Frist() {
        idx = 0;
        return this;
    }
    int Current() {
        return v[idx];
    }

};

void test_iterator() {
    vector<int> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);

    IntIterator* it = new IntIter(v);
    it = it->Next();
    cout << "No2 it access " << it->Current() << endl;
}
int main() {
    test_iterator();
    return 0;
}