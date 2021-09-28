#include <iostream>
using namespace std;

class Shape {
public:
    virtual void Draw() = 0;
};

class Circle : public Shape {
public:
    void Draw() {
        cout << "Circle" << endl;
    }
};
class Square : public Shape {
public:
    void Draw() {
        cout << "Square" << endl;
    }    
};
class Rectangle : public Shape {
public:
    void Draw() {
        cout << "Rectangle" << endl;
    }    
};
class Factory {
public:
    Shape* Create(std::string type) {
        if (type == "Circle") {
            return new Circle;
        } else if (type == "Square") {
            return new Square;
        } else if (type == "Rectangle") {
            return new Rectangle;
        } else {
            return NULL;
        }
    }
};

void test_factory() {
    Factory factory;
    auto s = factory.Create("Circle");
    s->Draw();
}

int main() {
    test_factory();
    return 0;
}