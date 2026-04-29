#include <iostream>
#include <utility>

struct DataPair {
  int first_value;
  int second_value;
};

void PrintTopBorder() { std::cout << "===" << std::endl; }

void PrintSeparator() { std::cout << "---" << std::endl; }

void PrintOutputLabel() { std::cout << "Output:" << std::endl; }

void SortByFirstValue(DataPair data[], int length) {
  for (int i = 0; i < length; ++i) {
    for (int j = 0; j < length - 1; ++j) {
      if (data[j].first_value > data[j + 1].first_value) {
        std::swap(data[j].first_value, data[j + 1].first_value);
        std::swap(data[j].second_value, data[j + 1].second_value);
      }
    }
  }
}

int CalculateResult(const DataPair& pair) {
  if (pair.first_value % 2 == 0) {
    if (pair.second_value % 2 == 0) {
      return pair.first_value * pair.second_value;
    }
    return pair.first_value + pair.second_value;
  }

  if (pair.second_value % 2 == 0) {
    return pair.first_value - pair.second_value;
  }

  return pair.first_value;
}

void PrintCalculatedValues(const DataPair data[], int size) {
  for (int i = 0; i < size; ++i) {
    std::cout << CalculateResult(data[i]) << std::endl;
  }
}

void ProcessData(DataPair data[], int length) {
  SortByFirstValue(data, length);
  PrintTopBorder();
  PrintOutputLabel();
  PrintSeparator();
  PrintCalculatedValues(data, length);
}

int main() {
  DataPair data[5];

  data[0].first_value = 5;
  data[0].second_value = 1;
  data[1].first_value = 2;
  data[1].second_value = 4;
  data[2].first_value = 3;
  data[2].second_value = 7;
  data[3].first_value = 1;
  data[3].second_value = 6;
  data[4].first_value = 4;
  data[4].second_value = 5;

  ProcessData(data, 5);
  return 0;
}
