//
// Created by Brandon Li on 6/3/23.
//

// currently WA
// https://codebreaker.xyz/submission/245769

#include <bits/stdc++.h>
using namespace std;

typedef long long LL;
typedef long double LD;
typedef unsigned long long ULL;
typedef pair<int, int> pii;
typedef pair<LL, LL> pll;
typedef pair<string, string> pss;
typedef vector<int> vi;
typedef vector<LL> vl;
typedef vector<vi> vvi;
typedef vector<pii> vii;
typedef vector<pll> vll;
typedef vector<vl> vvl;
typedef map<int, int> mii;
typedef map<LL, LL> mll;
typedef set<int> si;
typedef set<LL> sl;
typedef set<char> sc;

double EPS = 1e-9;
int INF = 1000000005;
LL INFF = 1000000000000000005LL;
double PI = acos(-1);
int dirx[8] = { -1, 0, 0, 1, -1, -1, 1, 1 };
int diry[8] = { 0, 1, -1, 0, -1, 1, -1, 1 };

#define fast ios_base::sync_with_stdio(false);cin.tie(NULL);cout.tie(0);

#ifdef TESTING
#define DEBUG fprintf(stderr, "====TESTING====\n")
#define VALUE(x) cerr << "The value of " << #x << " is " << x << endl
#define debug(...) fprintf(stderr, __VA_ARGS__)
#else
#define DEBUG
#define VALUE(x)
#define debug(...)
#endif

#define FOR(a, b, c) for (int(a) = (b); (a) < (c); ++(a))
#define FORA(a, b, c) for (auto(a) = (b); (a) != (c); ++(a))
#define FORN(a, b, c) for (int(a) = (b); (a) <= (c); ++(a))
#define FORD(a, b, c) for (int(a) = (b); (a) >= (c); --(a))
#define FORSQ(a, b, c) for (int(a) = (b); (a) * (a) <= (c); ++(a))
#define FORC(a, b, c) for (char(a) = (b); (a) <= (c); ++(a))
#define FOREACH(a, b) for (auto&(a) : (b))
#define REP(i, n) FOR(i, 0, n)
#define REPN(i, n) FORN(i, 1, n)
#define MAX(a, b) a = max(a, b)
#define MIN(a, b) a = min(a, b)
#define SQR(x) ((LL)(x) * (x))
#define RESET(a, b) fill(a, a + sizeof(a) / sizeof(a[0]), b)
#define RESET2D(a, b) fill(&a[0][0], &a[0][0] + sizeof(a) / sizeof(a[0][0]), b)
#define SET(a) RESET(a, 0)
#define SET2D(a) RESET2D(a, 0)
#define fi first
#define se second
#define mp make_pair
#define pb push_back
#define ALL(v) v.begin(), v.end()
#define VINDEX(v, i) v.begin() + i
#define SIZE(v) (int)v.size()
#define SORT(v) sort(ALL(v))
#define SORTS(v) sort(ALL(v), sortbysec)
#define REVERSE(v) reverse(ALL(v))
#define ALLA(arr, sz) arr, arr+sz
#define SIZEA(arr) sizeof(arr)/sizeof(arr[0])
#define ALLS(arr) ALLA(arr, SIZEA(arr))
#define SORTA(arr, sz) sort(ALLA(arr, sz))
#define REVERSEA(arr, sz) reverse()
#define PERMUTE next_permutation
#define TC(t) while (t--)
#define sp <<' '<<
#define endl '\n'
#define GET(n) LL n; cin >> n
#define GETT(n, t) t n; cin >> n
#define GETN LL n; cin >> n;
#define NEXT(n, a) REP(i, n) cin >> a[i]
#define NEXTN(n, a, i) REP(i, n) cin >> a[i]
#define print(x) cout << x << endl;
#define printw(x) cout << x;


LL gcd(LL a, LL b ){
    if (b == 0) return a;
    return gcd(b, a%b);
}

LL lcm(LL a, LL b) {
    return a/gcd(a, b)*b;
}

bool prime(LL a) {
    if (a == 1) return 0;
    FORN(i, 2, round(sqrt(a))) if (a%i == 0) return 0;
    return 1;
}

bool sortbysec(const pair<int,int> &a,
               const pair<int,int> &b) {
    return (a.second < b.second);
}

void rebalance(priority_queue<LL, vector<LL>, greater<LL>>& min_heap,
               priority_queue<LL, vector<LL>, less<LL>>& max_heap) {
    while (min_heap.size() > max_heap.size() + 1) {
        max_heap.push(min_heap.top());
        min_heap.pop();
    }
    while (max_heap.size() > min_heap.size() + 1) {
        min_heap.push(max_heap.top());
        max_heap.pop();
    }
}

void add_element(LL val,
                 priority_queue<LL, vector<LL>, greater<LL>>& min_heap,
                 priority_queue<LL, vector<LL>, less<LL>>& max_heap) {
    if (max_heap.empty() || val < max_heap.top()) {
        max_heap.push(val);
    } else {
        min_heap.push(val);
    }

    rebalance(min_heap, max_heap);
}

void remove_element(priority_queue<LL, vector<LL>, greater<LL>>& min_heap,
                    priority_queue<LL, vector<LL>, less<LL>>& max_heap) {
    // if one heap is greater than the other by more than 1, remove the top
    // else if both heaps are equal, remove from max_heap

//    if (max_heap.size() > min_heap.size() + 1) {
//        max_heap.pop();
//    } else

    if (min_heap.size() > max_heap.size()) {
        min_heap.pop();
    } else if (max_heap.empty() != 1){
        max_heap.pop();
    }

    rebalance(min_heap, max_heap);
}

int main() {
    fast GETN;

    priority_queue<LL, vl, greater<LL>> min_heap;
    priority_queue<LL, vl, less<LL>> max_heap;

    REP(i, n) {
        GETT(x, string);
        if (x == "PUSH") {
            GET(num);
            add_element(num, min_heap, max_heap);
        } else if (x == "POP") {
            remove_element(min_heap, max_heap);
        }
    }

    vl max_heap_v;
    while (!max_heap.empty()) {
        max_heap_v.pb(max_heap.top());
        max_heap.pop();
    }

    REVERSE(max_heap_v);

    for (auto& x : max_heap_v) {
        printw(x);
        printw(" ");
    }

    while (!min_heap.empty()) {
        printw(min_heap.top());
        printw(" ");
        min_heap.pop();
    }
}
