#include <iostream>
#include <map>
using namespace std;

class Shape {
public:
    virtual void draw() = 0;
};

class ColorCricle : public Shape{
public:
    void draw() {
        cout << color << " color draw" << endl;
    }
public:
    string color;
};

class ColorCricleFactory {
public:
    ColorCricle* createColorCricle(const string& color) {
        ColorCricle* c = new ColorCricle;
        c->color = color;
        return c;
    }
};

class Flyweight {
public:
    ColorCricle* getCricle(string color) {
        if (map[color] == NULL) {
            cout << "new color " << color << endl;
            map[color] = factory.createColorCricle(color);
        } else {
            cout << "flyweight color " << color << endl;
        }
        return map[color];
    }
public:
    ColorCricleFactory factory;
    map<string, ColorCricle*> map;
};

void test_flywight() {
    Flyweight flyweight;
    flyweight.getCricle("red");
    flyweight.getCricle("yellow");
    flyweight.getCricle("red");
}

int main() {
    test_flywight();
    return 0;
}