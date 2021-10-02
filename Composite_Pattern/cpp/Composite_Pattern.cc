#include <iostream>
#include <vector>
using namespace std;

class Inode {
public:
    virtual void search(const string& s) = 0;
};

class ComFloder : public Inode {
public:
    ComFloder(const string& s):name(s) {}
public:
    void search(const string& s) {
        if (s == name) {
            cout << "search file from floder " << name << endl;
        }
        auto i = v.begin();
        while (i != v.end()) {
            (*i)->search(s);
            i++;
        }
    }
    void Add(Inode* p) {
        v.push_back(p);
    }
private:
    vector<Inode*> v;
    string name;
};

class ComLeaf : public Inode {
public:
    ComLeaf(const string& s):name(s) {}
public:
    void search(const string& s) {
        if (s == name) {
            cout << "search file from file " << name << endl;
        }
    }
private:
    string name;
};
void test_composite() {
    ComLeaf l1("l1");
    ComLeaf l2("l2");
    ComLeaf l3("l3");

    ComFloder n0("n0");

    n0.Add(&l1);
    n0.Add(&l2);
    n0.Add(&l3);

    ComFloder n1("n1");

    ComLeaf l11("l11");
    ComLeaf l12("l12");
    ComLeaf l13("l13");

    n1.Add(&l11);
    n1.Add(&l12);
    n1.Add(&l13);

    n0.Add(&n1);

    n0.search("l13");
}
int main() {
    test_composite();
    return 0;
}