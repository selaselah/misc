#include <vector>     
#include <algorithm>     
using namespace std;
struct stu {
  int score;
  bool operator<(const stu &s) const{
    return score <= s.score;
  }
};

int main(int argvc, char **argv){
  vector<stu> v;
  stu s;
  for(int i = 0; i < 17; i++){
    s.score = 1;
    v.push_back(s);
  }
  sort(v.begin(), v.end());
  return 0;
}
