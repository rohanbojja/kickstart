using System;
using System.Linq;
using System.Collections.Generic;

namespace kick_start
{
    class Program
    {
        static void Main(string[] args)
        {
            int cases = int.Parse(Console.ReadLine());
            for (int testCase = 1; testCase < cases + 1; testCase++)
            {
                var s = Console.ReadLine();
                int indexA = 0;
                int ans = 0;
                List<int> kicks = new List<int>();
                List<int> starts = new List<int>();

                while (indexA >= 0)
                {
                    indexA = s.IndexOf("KICK", indexA);
                    if (indexA != -1)
                    {
                        kicks.Add(indexA);
                        indexA++;
                    }
                }
                indexA = 0;
                while (indexA >= 0)
                {
                    indexA = s.IndexOf("START", indexA);
                    if (indexA != -1)
                    {
                        starts.Add(indexA);
                        indexA++;
                    }
                }
                foreach(int kickIndex in kicks){
                    ans += starts.Count( x => x>kickIndex);
                }
                Console.WriteLine($"Case #{testCase}: {ans}");
            }
        }
    }
}
