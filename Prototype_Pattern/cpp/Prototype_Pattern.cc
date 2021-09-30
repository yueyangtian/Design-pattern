#include<iostream>
#include<vector>
using namespace std;

class inode {
public:
    virtual void Print() = 0;
    virtual inode * Clone() = 0;
};

class file : virtual public inode {
public:
    void Print() {
        cout << "file " << name << endl;
    }
    inode * Clone() {
        file* f = new file;
        f->name = this->name;
        return f;
    }
public:
    string name;
};


class floder : public inode {
public:
    void Print() {
        cout << "floder " << name << endl;
        auto i = childen.begin();
        while (i != childen.end())
        {
            (*i)->Print();
            i++;
        }
    }
    inode * Clone() {
        floder* f = new floder;
        f->name = name;
        auto i = childen.begin();
        while (i != childen.end())
        {
            f->childen.push_back((*i)->Clone());
            i++;
        }
        return f;
    }
public:
    string name;
    std::vector<inode*> childen;  
};

void test_prototype() {
    file* f1 = new file;
    f1->name = "file1";

    floder* fc = new floder;
    fc->name = "floderchild";

    floder* ff = new floder;
    ff->name = "floderfather";
    ff->childen.push_back(fc);
    ff->childen.push_back(f1);

    f1->Print();
    fc->Print();
    ff->Print();

    auto f = ff->Clone();
    f->Print();

}


int main() {
    test_prototype();
    return 0;
}