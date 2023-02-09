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

//Console.WriteLine($"Calculation : text1 = {text1} , text2 = {text2} :
//Return value = {result}");
result = r.Calculate("kitten", "sitting");
result = r.Calculate("มกราคม", "มกรคม");
result = r.Calculate("มกราคม", "มกคม");

result = r.Calculate("กุมถาพันธ์", "มกรคม");
result = r.Calculate("มีนาคม", "มกรคม");

result = r.Calculate("ABBYY", "ABBY");
result = r.Calculate("HONDA", "HONDO");

result = r.Calculate("ABBYYYY", "ABBY");
result = r.Calculate("hospital", "haspita");
result = r.Calculate("haspita", "hospital");
result = r.Calculate("1.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "1.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล");
result = r.Calculate("2.เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "2..เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ.กา.รเ.มีองฺ.ท้องถนสุงฺ.กู้^.ชินร^.แส?เทศบาล");
result = r.Calculate(".เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ.กา.รเ.มีองฺ.ท้องถนสุงฺ.กู้^.ชินร^.แส?เทศบาล", "เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล");
result = r.Calculate("3.เงินกู้สวัสดิการแก่ข้าราชการ", "3..เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ");
result = r.Calculate(".เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ", "เงินกู้สวัสดิการแก่ข้าราชการ");
result = r.Calculate("4.สินเชื่อ", "4..สิน!?ชื่อ.");
result = r.Calculate("สินเชื่อ", ".สิน!?ชื่อ.");
result = r.Calculate("5.สินเชื่อ", "5.ลัน!?ซื่ซี");
result = r.Calculate("สินเชื่อ", "ลัน!?ซื่ซี");
result = r.Calculate("สินเชื่อ", "ลินเชื่อ");
result = r.Calculate("สินเชื่อ", "ลิมเชื่อ");
result = r.Calculate("สินเชื่อ", "ลิน1ชื่อ");
result = r.Calculate("สินเชื่อ", "ลินเชื่");
result = r.Calculate("GILY", "GEELY");
result = r.Calculate("HONDA", "HYUNDAI");
result = r.Calculate("GEELY", "GILY");
result = r.Calculate("HYUNDAI", "HONDA");
result = r.Calculate("FLOMAX", "VOLMAX");
result = r.Calculate("VOLMAX", "FLOMAX");
result = r.Calculate("เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องลิ่นสังกัด อบต. อบจ. และเทศบาล");
result = r.Calculate("เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", ".เงินกสฺ.วัสุดฺ.ทาร.แก่ข้าราชการ.กา.รเ.มีองฺ.ท้องถนสุงฺ.กู้^.ชินร^.แส?เทศบาล");
result = r.Calculate("เงินกู้สวัสดิการแก่ข้าราชการ การเมืองท้องถิ่นสังกัด อบต. อบจ. และเทศบาล", "เงินก้สวัสดิการแก*ข้าราซการ การเมืองท-องกนสังกัด ชินต. อบจ. แสเทศบาส");
result = r.Calculate("สินเชื่อที่อยู่อาศัยที่ธนาคารออกค่าจำนองให้สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย (ห้ามปิดก่อน 3 ปี)", ".สินเชื่อ.ที่.อยู่อา.ศัยฺที่ธนาคารออกค่าจํานองฺให้สําหรับลูกค้าใน.อาน่าจของผู้ บริ๋หารกลุ่มเครือข่าย(หามปิดก่อน 3 บี)");
result = r.Calculate("สินเชื่อที่อยู่อาศัยที่ธนาคารออกค่าจำนองให้สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย(ห้ามปิดก่อน 3 ปี)", ".สิน!?ชื่อ.ที่.'อุ.ยู่'อุาศัยุ.ที่ธนาคารฺออฺกเค่าจํานองให้สำหรับลูกค้าในอำนาจของผู้ บร์หารกลุ่มเครํ่อขาย(ห้ามปืดกืชิน 3 บี)");
result = r.Calculate("สินเชื่อที่อยู่อาศัยที่ธนาคารออกค่าจำนองให้สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย(ห้ามปิดก่อน 3 ปี)", " ลัน!?ซื่ซีที่.อุธ์อุๆ.ศัยุ.ที่ธนาคารอซีกุค่า.จํานองให้สําหรับุฐ.กุ.ค่าในอํา");
result = r.Calculate("สินเชื่ออเนกประสงค์แบบพิเศษ สำหรับองค์กรปกครองส่วนท้องถิ่น สินเชื่ออเนกประสงค์แบบพิเศษ", "สําหรับองค์กรปกครองส่วนท้องถิ่น");
result = r.Calculate("สินเชื่ออเนกประสงค์แก่พนักงานรัฐวิสาหกิจ (ดำรงเงินฝาก)", ".ชินุเฃีอ.'อุ.เนุทป.รุะสุงค์สิกุ.หรับพนัก.งานุรัซีวิสาหกิจ“ตํารงฺ.เงินฝาก");
result = r.Calculate("สินเชื่อ Home For Cash สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย สินเชื่อ Home For Cash", "สินเชื่อ Home For Cash สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย");
result = r.Calculate("สินเชื่อ Home For Cash สำหรับลูกค้าในอำนาจของผู้บริหารกลุ่มเครือข่าย สินเชื่อ Home For Cash", "สินเชื่อ Home For Cash");
result = r.Calculate("สินเชื่อ Home For Cash อัตราดอกเบี้ยพิเศษ ในอำนาจของผู้บริหารกลุ่มเครือข่าย (เพื่อป้องกัน Refinance)", ".สินเชื่ชิ.HQmefFor Cash อัตรุาดอฺ.กุเบี้ย.พิเศ ขาย (เพือป้องกัน Refinance)");


//Console.Write($"{Environment.NewLine}Press any key to exit...");
//Console.ReadKey(true);

