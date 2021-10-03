#include <iostream>
using namespace std;


class image {
public:
    virtual void display() = 0;
};

class RealImage : public image {
public:
    RealImage(const string& n):name(n) {}
public:
    string name;
public:
    void display() {
        load();
        cout << "image displayed" << endl;
    }
private:
    void load() {
        cout << name << " image loading" << endl;
    }
};

class ProxyImage : public image {
public:
    ProxyImage():image("") {}
public:
    void proxy(const string& n) {
        if (image.name != n) {
            image.name = n;
        }
    }
    void display() {
        image.display();
    }
private:
    RealImage image;
};

void test_proxy() {
    ProxyImage p1;
    p1.proxy("1.jpg");
    p1.display();
}


int main() {
    test_proxy();
    return 0;
}
