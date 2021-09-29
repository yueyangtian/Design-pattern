#include<iostream>
using namespace std;

class Car {
public:
    virtual void Drive() = 0;
    virtual void Look() = 0;
};

class CarBuilder {
public:
    virtual CarBuilder& SetSpeed(int sp) = 0;
    virtual CarBuilder& SetColor(string cr) = 0;
    virtual Car* Build() = 0;
};

class car : public Car {
private:
    int speed;
    string color;
public:
    car(const int sp, const string cr):speed(sp),color(cr){}
    void Drive() {
        cout << "speed " << speed << endl;
    }
    void Look() {
        cout << "see " << color << endl;
    }
};

class cb : public virtual CarBuilder {
private:
    int speed;
    string color;
public:
    CarBuilder& SetSpeed(int sp) {
        speed = sp;
        return *this;
    }
    CarBuilder& SetColor(string cr) {
        color = cr;
        return *this;
    }

    Car* Build() {
        Car* c = new car(speed, color);
        return c;
    }
};


void test_builder() {
    Car* c1 = cb().SetSpeed(200).SetColor("yellow").Build();
    c1->Drive();
    c1->Look();

    Car* c2 = cb().SetSpeed(180).SetColor("red").Build();
    c2->Drive();
    c2->Look();
}


int main() {
    test_builder();
    return 0;
}
