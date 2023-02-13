using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Diagnostics;

namespace cseditdistance
{
  internal class MyMatch
  {
        public static void Print2DArray(Array arr)
        {
            //Console.WriteLine();
            int min_x = arr.GetLowerBound(0);
            int min_y = arr.GetLowerBound(1);
            int max_x = arr.GetUpperBound(0);
            int max_y = arr.GetUpperBound(1);
            String[,] array2d = (string[,])arr;

            Console.WriteLine("min_x = " + min_x + ", max_x = " + max_x);
            Console.WriteLine("min_y = " + min_y + ", max_y = " + max_y);
            for (int x = min_x; x <= max_x; x++)
            {
                Console.Write("|");
                for (int y = min_y; y <= max_y; y++)
                {
                    Console.Write(AddSpace(array2d[x, y], 2));
                }
                Console.WriteLine(" |");
            }

        }

        public static void Print2DIntArray(Array arr, bool isDebug)
        {
            //Console.WriteLine();
            int min_x = arr.GetLowerBound(0);
            int min_y = arr.GetLowerBound(1);
            int max_x = arr.GetUpperBound(0);
            int max_y = arr.GetUpperBound(1);
            int[,] array2d = (int[,])arr;

            if (isDebug)
            {
                Debug.WriteLine("min_x = " + min_x + ", max_x = " + max_x);
                Debug.WriteLine("min_y = " + min_y + ", max_y = " + max_y);
            } else
            {
                Console.WriteLine("min_x = " + min_x + ", max_x = " + max_x);
                Console.WriteLine("min_y = " + min_y + ", max_y = " + max_y);
            }

            int padding = 3;
            for (int x = min_x; x <= max_x; x++)
            {
                if (isDebug) { Debug.Write("|"); } else { Console.Write("|"); }
                
                for (int y = min_y; y <= max_y; y++)
                {
                    if (isDebug) { 
                        Debug.Write(AddSpace(array2d[x, y], padding));
                    } else {
                        Console.Write(AddSpace(array2d[x, y], padding));
                    }
                }

                if (isDebug) { Debug.WriteLine(" |"); } else { Console.WriteLine(" |"); }
            }

        }
        private static String AddSpace(int s, int num)
        {
            //int i = s.Length;
            String ret = s.ToString().PadLeft(num, ' ');
            return ret;
        }
        private static String AddSpace(String s, int num)
        {
            //int i = s.Length;
            String ret = s.ToString().PadLeft(num, ' ');
            return ret;
        }
    }
}
