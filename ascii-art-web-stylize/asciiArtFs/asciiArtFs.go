package asciiArtFs

import (
	"asciiArtWeb/asciiArtFs/myFunctions" // Importing custom functions from myFunctions package.
	// استيراد الدوال المخصصة من حزمة myFunctions.

	"fmt"   // Used for formatted I/O operations like printing and error handling.
	// تُستخدم لعمليات الإدخال والإخراج المنسقة مثل الطباعة ومعالجة الأخطاء.

	"log"   // Used for logging error messages or other important runtime information.
	// تُستخدم لتسجيل رسائل الخطأ أو المعلومات الأخرى المهمة أثناء التشغيل.
)

func AsciiArtFs(text string, banner string) (string, error) {
	// AsciiArtFs function: generates ASCII art based on input text and the specified banner file.
	// دالة AsciiArtFs: تولد فن ASCII بناءً على النص المدخل وملف البانر المحدد.

	banner = "asciiArtFs/" + banner + ".txt"
	// Set the file path for the selected banner by appending ".txt" to the banner name.
	// تعيين مسار الملف للبنر المحدد عن طريق إضافة ".txt" إلى اسم البانر.

	standard, err := myfunctions.Read(banner)
	// Read the banner file and store its content in the `standard` variable.
	// قراءة ملف البانر وتخزين محتوياته في المتغير `standard`.

	if err != nil {
		return "NotFound", fmt.Errorf("")
		// If the banner file is not found, return "NotFound" and an empty error.
		// إذا لم يتم العثور على ملف البانر، يتم إرجاع "NotFound" وخطأ فارغ.
	}

	asciiChars := myfunctions.BytesToAsciiMap(standard)
	// Convert the banner's byte content into an ASCII map for further processing.
	// تحويل محتوى البايت من البانر إلى خريطة ASCII لمزيد من المعالجة.

	result, err := myfunctions.WriteResult(text, asciiChars)
	// Generate the ASCII art by writing the result based on the input text and ASCII map.
	// توليد فن ASCII بكتابة النتيجة بناءً على النص المدخل وخريطة ASCII.

	if err != nil {
		log.Println(err)
		// Log any errors that occur during ASCII generation.
		// تسجيل أي أخطاء تحدث أثناء توليد ASCII.

		return "", fmt.Errorf("")
		// Return an empty string and an empty error if ASCII generation fails.
		// إرجاع سلسلة فارغة وخطأ فارغ إذا فشل توليد ASCII.
	}

	res := String(result)
	// Convert the result slice of strings into a single formatted string.
	// تحويل شريحة النتائج من السلاسل النصية إلى سلسلة منسقة واحدة.

	return res, nil
	// Return the formatted ASCII art string and nil error.
	// إرجاع سلسلة فن ASCII المنسقة وخطأ nil.
}

func String(result []string) string {
	// String function: converts a slice of strings into a single HTML-formatted string.
	// دالة String: تقوم بتحويل شريحة من السلاسل النصية إلى سلسلة HTML منسقة.

	str := ""
	// Initialize an empty string to hold the final result.
	// تهيئة سلسلة فارغة لتخزين النتيجة النهائية.

	for _, v := range result {
		// Iterate through each string in the result slice.
		// تكرار على كل سلسلة نصية في شريحة النتائج.

		v = replaceSpaces(v)
		// Call replaceSpaces function to trim trailing spaces from the string.
		// استدعاء دالة replaceSpaces لقص الفراغات الزائدة من السلسلة.

		str += v + "<br>"
		// Append the string to the result with an HTML line break ("<br>").
		// إضافة السلسلة إلى النتيجة مع فاصل أسطر HTML ("<br>").
	}

	return str
	// Return the final HTML-formatted string.
	// إرجاع السلسلة المنسقة النهائية.
}

func replaceSpaces(str string) string {
	// replaceSpaces function: removes trailing spaces from a string.
	// دالة replaceSpaces: تزيل الفراغات الزائدة من نهاية السلسلة.

	res := ""
	// Initialize an empty string to store the trimmed result.
	// تهيئة سلسلة فارغة لتخزين النتيجة بعد إزالة الفراغات.

	for i := range str {
		// Iterate through each character in the string.
		// التكرار على كل حرف في السلسلة.

		if OnlySpaces(str[i:]) {
			break
			// If the remaining part of the string consists only of spaces, stop.
			// إذا كانت بقية السلسلة تتكون فقط من فراغات، توقف.
		}

		res += string(str[i])
		// Append non-space characters to the result string.
		// إضافة الأحرف غير الفراغية إلى سلسلة النتيجة.
	}

	return res
	// Return the trimmed string without trailing spaces.
	// إرجاع السلسلة بعد إزالة الفراغات الزائدة.
}

func OnlySpaces(str string) bool {
	// OnlySpaces function: checks if a string consists only of spaces.
	// دالة OnlySpaces: تتحقق مما إذا كانت السلسلة تتكون فقط من فراغات.

	for _, v := range str {
		// Iterate through each character in the string.
		// التكرار على كل حرف في السلسلة.

		if v != ' ' {
			return false
			// If any character is not a space, return false.
			// إذا كان أي حرف ليس فراغاً، يتم إرجاع false.
		}
	}

	return true
	// If all characters are spaces, return true.
	// إذا كانت جميع الأحرف فراغات، يتم إرجاع true.
}






// package asciiArtFs

// import (
// 	"asciiArtWeb/asciiArtFs/myFunctions"
// 	"fmt"
// 	"log"
// )

// func AsciiArtFs(text string, banner string) (string, error) {
// 	banner = "asciiArtFs/" + banner + ".txt"
// 	standard, err := myfunctions.Read(banner)
// 	if err != nil {
// 		return "NotFound", fmt.Errorf("")
// 	}
// 	asciiChars := myfunctions.BytesToAsciiMap(standard)
// 	result, err := myfunctions.WriteResult(text, asciiChars)
// 	if err != nil {
// 		log.Println(err)
// 		return "", fmt.Errorf("")
// 	}
// 	res := String(result)
// 	return res, nil
// }

// func String(result []string) string {
// 	str := ""
// 	for _, v := range result {
// 		v = replaceSpaces(v)
// 		str += v + "<br>"
// 	}
// 	return str
// }

// func replaceSpaces(str string) string {
// 	res := ""
// 	for i := range str {
// 		if OnlySpaces(str[i:]) {
// 			break
// 		}
// 		res += string(str[i])
// 	}
// 	return res
// }

// func OnlySpaces(str string) bool {
// 	for _, v := range str {
// 		if v != ' ' {
// 			return false
// 		}
// 	}
// 	return true
// }