class Solution {

  public int findRotateSteps(String ring, String key) {
    int[][] dp = initDP(key.length(), ring.length());

    for (int i = 0; i < key.length(); i++) {
      for (int j = 0; j < ring.length(); j++) {
        if (key.charAt(i) == ring.charAt(j)) {
          if (i == 0) {
            dp[i][j] = minStep(0, j, ring.length()) + 1;
          } else {
            for (int x = 0; x < ring.length(); x++) {
              if (dp[i - 1][x] != Integer.MAX_VALUE) {
                dp[i][j] = min(dp[i][j], dp[i - 1][x] + minStep(x, j, ring.length()) + 1);
              }
            }
          }
        }
      }
    }
    return min(dp[dp.length - 1]);
  }

  private int[][] initDP(int rowCnt, int colCnt) {
    int[][] dp = new int[rowCnt][colCnt];
    for (int i = 0; i < rowCnt; i++) {
      for (int j = 0; j < colCnt; j++) {
        dp[i][j] = Integer.MAX_VALUE;
      }
    }
    return dp;
  }

  private int minStep(int pos, int target, int length) {
    return min(abs(target - pos), length - abs(target - pos));
  }

  private int min(int a, int b) {
    return a <= b ? a : b;
  }

  private int min(int... a) {
    int m = Integer.MAX_VALUE;
    for (int i : a) {
      if (i < m) {
        m = i;
      }
    }
    return m;
  }

  private int abs(int x) {
    return x >= 0 ? x : 0 - x;
  }
}
