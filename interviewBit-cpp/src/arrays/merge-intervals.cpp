#include <iostream>

using namespace std;

// https://www.interviewbit.com/problems/merge-intervals/
struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};

bool isOverlappingInterval(Interval interval1, Interval interval2) {
    return ((interval1.start >= interval2.start && interval1.start <= interval2.end)
        || (interval1.end >= interval2.start && interval1.end <= interval2.end)
        || (interval2.start >= interval1.start && interval2.start <= interval1.end)
        || (interval2.end >= interval1.start && interval2.end <= interval1.end));
}

// T(n) : O(n) ; S(n) : O(n)
vector<Interval> mergeIntervals(vector<Interval> &intervals, Interval newInterval) {
    
    int n = intervals.size();
    vector<Interval> mergedIntervals;
    if (n == 0) {
        mergedIntervals.push_back(newInterval);
        return mergedIntervals;
    }

    int start = intervals[0].start, end = intervals[0].end;
    if (newInterval.end < start) {
        mergedIntervals.push_back(newInterval);
        for (auto &interval : intervals) {
            mergedIntervals.push_back(interval);
        }
        return mergedIntervals;
    }
    
    bool flag = false;
    mergedIntervals.push_back(newInterval);
    for (int i = 0; i < n; i++) {
        if (isOverlappingInterval(interval[i], mergedIntervals[i-1])) {
            mergedIntervals[i-1].start = min(mergedIntervals[i-1].start, newInterval.start);
            mergedIntervals[i-1].end = max(mergedIntervals[i-1].end, newInterval.end);
            flag = true;
        } else {
            mergedIntervals.push_back(intervals[i]);
        }
    }
    
    if (!flag) {
        mergedIntervals.push_back(newInterval);
    }

    return mergedIntervals;
}

// Driver Code for testing
int main() {
    
    return 0;
}
