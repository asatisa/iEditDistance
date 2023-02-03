using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace cseditdistance
{
  
  public class RunTest
  {
    private static int i_count = 0;
    public int Calculate(string source1, string source2) //O(n*m)
    {
      return CalculateLevenshteinDistance_Percent(source1, source2);
    }

    public int CalculateLevenshteinDistance(string source1, string source2) //O(n*m)
    {  
      int return_int = 0;
      var currentDate = DateTime.Now;

      return_int = LevenshteinDistance.Calculate(source1, source2);

      Console.WriteLine($"{currentDate:d}-{currentDate:t}: Result = {return_int} #: When text1 = {source1} , text2 = {source2}");
      return return_int;
     }

    public int CalculateLevenshteinDistance_Percent(string source1, string source2) //O(n*m)
    {
      i_count = i_count + 1;
      int return_int = 0;
      var currentDate = DateTime.Now;
      var source1Length = source1.Length;
      var source2Length = source2.Length;
      var maxLength = 0;
      if (source1Length >= source2Length)
      {
        maxLength = source1Length;
      } else
      {
        maxLength = source2Length;
      }

      float percent = 0F;
      return_int = LevenshteinDistance.Calculate(source1, source2);
      percent = ((float)maxLength - (float)return_int) / (float)maxLength * 100F;

      //Console.WriteLine($"{i_count} : Result = {return_int}, {maxLength}, {percent}% #: When text1 = {source1} , text2 = {source2}");
      Console.WriteLine($"{i_count} : Result = {return_int}, {maxLength}, {percent}% #");
      //return_int = (int)percent;
      return return_int;
    }
  }
}
