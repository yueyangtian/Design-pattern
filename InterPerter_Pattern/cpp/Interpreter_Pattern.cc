#include <iostream>
#include <vector>
#include <stdlib.h>
using namespace std;


class interpreter {
public:
    virtual int InterPreter() = 0;
};
int Option(vector<string>& v, vector<interpreter*>& iv);
class AddInterpreter : public interpreter {
public:
    AddInterpreter(interpreter* f, interpreter* s):frist(f),second(s){}
public:
    int InterPreter() {
        return frist->InterPreter() + second->InterPreter();
    }
private:
    interpreter* frist;
    interpreter* second;
};
class MutilInterpreter : public interpreter {
public:
    MutilInterpreter(interpreter* f, interpreter* s):frist(f),second(s){}
public:
    int InterPreter() {
        return frist->InterPreter() * second->InterPreter();
    }
private:
    interpreter* frist;
    interpreter* second;
};

class NumberInterpreter : public interpreter {
public:
    NumberInterpreter(const string s):num(s){}
public:
    int InterPreter() {
        return atoi(num.c_str());
    }
private:
    string num;
};
vector<string> split (string s, string delimiter) {
    size_t pos_start = 0, pos_end, delim_len = delimiter.length();
    string token;
    vector<string> res;

    while ((pos_end = s.find (delimiter, pos_start)) != string::npos) {
        token = s.substr (pos_start, pos_end - pos_start);
        pos_start = pos_end + delim_len;
        res.push_back (token);
    }

    res.push_back (s.substr (pos_start));
    return res;
}
int Parser(string s) {
    vector<string> v = split(s, " ");
    vector<interpreter*> iv;
    return Option(v, iv);
}
int Option(vector<string>& v, vector<interpreter*>& iv) {
    if (iv.size() == 1 && v.size() == 0) {
        return iv[0]->InterPreter();
    }
    if (v.size() == 0) {
        return 0;
    } 

    string n = v[0];
    if (n == "*") {
        if (iv.size() < 2 ) {
            return 0;
        }
        interpreter* second = iv.back();
        iv.pop_back();
        interpreter* frist = iv.back();
        iv.pop_back();

        interpreter* i = new MutilInterpreter(frist, second);
        iv.push_back(i);

    } else if (n == "+") {
        if (iv.size() < 2 ) {
            return 0;
        }
        interpreter* second = iv.back();
        iv.pop_back();
        interpreter* frist = iv.back();
        iv.pop_back();

        interpreter* i = new AddInterpreter(frist, second);
        iv.push_back(i);
    } else {
        interpreter* i = new NumberInterpreter(n);
        iv.push_back(i);
    }
    vector<string> vt;
    for (int i = 1; i < v.size(); ++i) {
        vt.push_back(v[i]);
    }
    return Option(vt, iv);
}
void test_interpreter() {
    cout << "6 100 11 + * 10 * is " << Parser("6 100 11 + * 10 *") << endl;
}
int main() {
    test_interpreter();
    return 0;
}