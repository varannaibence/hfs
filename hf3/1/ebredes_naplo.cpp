#include <iostream>
#include <string>

struct EbredesAdat {
    std::string nap;
    int ora;
    int perc;
};

int main()
{
    EbredesAdat naplo[7] = {
        {"Hetfo",     6, 45},
        {"Kedd",      7, 10},
        {"Szerda",    6, 30},
        {"Csutortok", 7,  0},
        {"Pentek",    6, 55},
        {"Szombat",   8, 20},
        {"Vasarnap",  9, 15}
    };

    std::cout << std::left;
    std::cout.width(12); std::cout << "Nap";
    std::cout << "Ebredes ideje\n";
    std::cout << "-----------------------------\n";

    for (int i = 0; i < 7; ++i) {
        std::cout.width(12); std::cout << naplo[i].nap;
        if (naplo[i].perc < 10)
            std::cout << naplo[i].ora << ":0" << naplo[i].perc << "\n";
        else
            std::cout << naplo[i].ora << ":" << naplo[i].perc << "\n";
    }

    return 0;
}
