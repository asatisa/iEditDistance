// See https://aka.ms/new-console-template for more information
using cseditdistance;
using System.Xml.Linq;

//var name = Console.ReadLine();
var currentDate = DateTime.Now;

Console.WriteLine("Calculation of Edit Distance! version.0.5");
Console.WriteLine($"{Environment.NewLine}On {currentDate:d} at {currentDate:t}!");

RunTest r = new RunTest();


int result = 0;

//string text1 = "ABBYY";
//string text2 = "ABBY";

//result = LevenshteinDistance.Calculate(text1, text2);

//Console.WriteLine($"Calculation : text1 = {text1} , text2 = {text2} : Return value = {result}");
result = r.Calculate("ABBYYYY", "ABBY");
result = r.Calculate("hospital", "haspita");
result = r.Calculate("haspita", "hospital");
result = r.Calculate("1.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "1.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล");
result = r.Calculate("2.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "2..เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ.กา.รเ.มีองฺ.ท้องถนสุงฺ.กู้^.ชินร^.แส?เทศบาล");
result = r.Calculate(".เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ.กา.รเ.มีองฺ.ท้องถนสุงฺ.กู้^.ชินร^.แส?เทศบาล","เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล");
result = r.Calculate("3.เงินกู้สวัสดิการแก่ข้าราชการ", "3..เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ");
result = r.Calculate(".เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ", "เงินกู้สวัสดิการแก่ข้าราชการ");
result = r.Calculate("4.สินเชื่อ", "4..สิน!■ชื่อ.");
result = r.Calculate("สินเชื่อ", ".สิน!■ชื่อ.");
result = r.Calculate("5.สินเชื่อ", "5.ลัน!■ซื่ซี");
result = r.Calculate("สินเชื่อ", "ลัน!■ซื่ซี");

//Console.Write($"{Environment.NewLine}Press any key to exit...");
//Console.ReadKey(true);

