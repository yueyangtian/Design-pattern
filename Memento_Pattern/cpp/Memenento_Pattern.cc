#include <iostream>
#include <vector>
using namespace std;

typedef string Text;

class TextMemento {
private:
    vector<Text> v;
public:
    Text Pop() {
        if (v.size() == 0) {
            return "";
        }
        Text t = v.back();
        v.pop_back();
        return t;

    }
    void Push(Text t) {
        v.push_back(t);
    }
};


class TextEdit {
private:
    Text text;
    TextMemento memento;
public:
    void Push(const char* s) {
        memento.Push(text);
        text = text + s;
    }
    void Revert() {
        text = memento.Pop();
    }
    void Show() {
        cout << text << endl;
    }
};

void test_memento() {
    TextEdit edit;
    edit.Push("hello ");
    edit.Push("world ");
    edit.Show();
    edit.Revert();
    edit.Show();
}

int main() {
    test_memento();
    return 0;
}