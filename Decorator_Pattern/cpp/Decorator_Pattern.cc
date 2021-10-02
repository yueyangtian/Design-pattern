#include <iostream>
using namespace std;

class Shape {
public:
    virtual void draw() = 0;
};

class Cricle : public Shape {
public:
    void draw() {
        cout << "cricle draw" << endl;
    }
};

class CricleDecorator {
public:
    CricleDecorator(Shape* s):shape(s) {}
public:
    void draw() {
        cout << "draw in decorator ";
        shape->draw();
    }
private:
    Shape* shape;
};

void test_decorator() {
    Cricle c;
    c.draw();
    CricleDecorator cd(&c);
    cd.draw();
}
int main() {
    test_decorator();
    return 0;
}