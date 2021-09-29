#include<iostream>
#include<mutex>
using namespace std;


std::mutex mtx;
class Singleton {
public:
    static Singleton* Instance;
    static Singleton* GetInstance() {
        if (Instance == NULL) {
            mtx.lock();
            if (Instance == NULL) {
                Instance = new Singleton;
            }
            mtx.unlock();
        }
        return Instance;
    }
private:
    Singleton(){};
public:
    Singleton(const Singleton&) = delete;
};

Singleton* Singleton::Instance = NULL;


int main() {
    auto a = Singleton::GetInstance();
    return 0;
}