#include <iostream>
using namespace std;
class Logger {
public:
    int level;
    Logger* nextLogger;
public:
    virtual void Log(int leve, const string& message) {
        if (leve == level) {
            write(message);
        } else {
            if (nextLogger) {
                nextLogger->Log(leve, message);
            } else {
                cout << "unknow level " << leve <<endl;
            }
        }
    }
    virtual void setNextLog(Logger* next) {
        nextLogger = next;
    }
private:
    virtual void write(const string& message) = 0;
};

class Debug : public Logger {
private:
    void write(const string& message) {
        cout << "[DEBUG] " << message << endl;
    } 
};

class Info : public Logger {
private:
    void write(const string& message) {
        cout << "[Info] " << message << endl;
    } 
};
class Error : public Logger {
private:
    void write(const string& message) {
        cout << "[ERROR] " << message << endl;
    } 
};

Logger* GetLogger() {
    Debug* d = new Debug;
    d->level = 1;

    Info* i = new Info;
    i->level = 2;

    Error* e = new Error;
    e->level = 3;

    d->setNextLog(i);
    i->setNextLog(e);

    return d;
}

void test_chainof() {
    Logger* log = GetLogger();
    log->Log(1, "hello");
    log->Log(4, "hello2");
}
int main() {
    test_chainof();
    return 0;
}