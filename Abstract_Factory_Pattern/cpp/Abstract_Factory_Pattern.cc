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
class Color {
public:
    virtual void Tint() = 0;
};
class Red : public Color {
public:
    virtual void Tint() {
        cout << "Red" <<endl;
    }
};
class Black : public Color {
public:
    virtual void Tint() {
        cout << "Black" <<endl;
    }
};

class Factory {
public:
    virtual Shape* CreateShape() = 0;
    virtual Color* CreateColor() = 0;
};

class RedCircleFactory : public Factory {
public:
    Shape* CreateShape() {
        return new Circle;
    }
    Color* CreateColor() {
        return new Red;
    }
};
class BlackRectangleFactory : public Factory {
        Shape* CreateShape() {
        return new Rectangle;
    }
    Color* CreateColor() {
        return new Black;
    }
};

void test_factory() {
    Factory* f = new RedCircleFactory;
    f->CreateColor()->Tint();
    f->CreateShape()->Draw();
}

int main() {
    test_factory();
    return 0;
}