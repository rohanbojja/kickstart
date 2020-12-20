using System;
using System.Linq;

namespace Record_Breaker
{
    class Program
    {
        static void Main(string[] args)
        {
            int cases = int.Parse(Console.ReadLine());
            for (int testCase = 1; testCase < cases + 1; testCase++)
            {
                var n = int.Parse(Console.ReadLine());
                var v = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
                int ans = 0;
                int currentRecord = -1;
                for (int i = 0; i < n; i++)
                {
                    if (i != n - 1)
                    {
                        if (v[i] > currentRecord && v[i] > v[i + 1])
                        {
                            ans += 1;
                            currentRecord = v[i];
                        }
                        currentRecord = Math.Max(currentRecord, v[i]);
                    }
                    else
                    {
                        if (v[i] > currentRecord)
                        {
                            ans += 1;
                        }
                    }
                }
                Console.WriteLine($"Case #{testCase}: {ans}");
            }
        }
    }
}
