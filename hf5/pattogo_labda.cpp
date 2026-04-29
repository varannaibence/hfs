#include <chrono>
#include <csignal>
#include <iostream>
#include <thread>

void kurzorVisszaallit(int) {
    std::cout << "\033[?25h\n";
    std::cout.flush();
    std::_Exit(0);
}

int main() {
    std::signal(SIGINT, kurzorVisszaallit);

    const int szelesseg = 40;
    const int magassag = 15;

    int x = szelesseg / 2;
    int y = magassag / 2;
    int dx = 1;
    int dy = 1;

    std::cout << "\033[2J\033[?25l";

    while (true) {
        std::cout << "\033[H";

        for (int sor = 0; sor < magassag; ++sor) {
            for (int oszlop = 0; oszlop < szelesseg; ++oszlop) {
                if (sor == 0 || sor == magassag - 1) {
                    std::cout << '#';
                } else if (oszlop == 0 || oszlop == szelesseg - 1) {
                    std::cout << '#';
                } else if (oszlop == x && sor == y) {
                    std::cout << 'O';
                } else {
                    std::cout << ' ';
                }
            }
            std::cout << '\n';
        }
        std::cout.flush();

        x += dx;
        y += dy;

        if (x <= 1 || x >= szelesseg - 2) {
            dx *= -1;
        }
        if (y <= 1 || y >= magassag - 2) {
            dy *= -1;
        }

        std::this_thread::sleep_for(std::chrono::milliseconds(60));
    }

    return 0;
}
