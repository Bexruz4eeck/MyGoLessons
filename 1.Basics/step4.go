package main

func main() { // STRING, NUMBER, BOOLEAN 

// STRING 

// имя - тип - значение - значение, o'zg - nomi -tipi - value - qiymat
 var message string  = "Hello Tashkent"
 println(message)

 var message1 string                       // DECLARE ---> вы представили,tanitdim  
 message1 = "Hello Tashkent"             // ASSIGMEND ---> я сравнял,tenglashtirdim
 println((message1))

 var box1, box2 string = "first", "second"
 println(box1,box2)

//NOTE! ---> ПРИМЕЧАНИЕ!
 var  box = "This is varible"  // типы не заданы но они всеравно работают, TIPI berilmagan ammo bu ham islaydi!
 println(box)                  // Если не дать ТИПИ, то он найдет автоматический и будет работать медленно,TIPI berilmasa o'zi avotomaticheskiy topib sekin ishlaydi!
                     

                          //если тип задан, код работает лучше, чем ограничение скорости

//____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________


// NUMBER 
// INT 
   
 var a, b, c int              //целое число,BUTUN SON Ex: 7, -5, 11, 23, -94, 768, 811
 println(box)                // ---> 0 0 0

var n1, n2, n3 int = 11, 22, 33
println(n1, n2, n3)

//___________________________________________________________________________________________________________________________________________________________________________________________________

// FLOAT-плавать,suzmoq

var x, y, z float32    //ОСТАВШИЙСЯ НОМЕР, QOLDIK SON  Ex 1.5, -11.6, 123.5, 4567.11, -12.5
println(x,y,z)        // ---> +0.000000e+000 +0.000000e+000 +0.000000e+000 

var  f1, f2, f3 float64 = 11.1, 22.2, 33.3
println(f1, f2, f3)

var flo1, flo2, flo3 float64 = 11.1, 22.2, 33.3
println(flo1 + flo2 + flo3)

//_______________________________________________________________________________________________________________________________________________________________________________________________________

// BOOLEAN-логическое значение,mantiqiy 

var erkakMI bool = true   // false
println(erkakMI)

//________________________________________________________________________________________________________________________________________________________________________________________________________________
  
// HAMMA O'ZGARUVCHILAR, ВСЕ ИЗМЕНЕНИЯ,
// Yani bitta satrga togridan togri harhil tipda qiymat birlashtirish mumkin,То есть есть возможность объединять значения разных типов в одну строку

var q, w, e, r, = 11, 22.3, "Hello", true
println(q, w, e, r)

//_________________________________________________________________________________________________________________________________________________________________________________________________________________________________

// NEW NOTE - Новое примечания!
// :  ---> bu operator ishlatilsa VAR va TIP yozilishi shartmas!, если используется этот оператор, то VAR и TIP писать не надо!
// o'zgaruvchisini va typini o'zi avtomat tanlab quyadi, автоматически выбирает переменную и тип
 
num, flo, stf,trufal := 11, 22.3,"Hello", false
println(num, flo, stf, trufal)  //bir satrlik o'zgaruvchi typlar kodi! однострочный код типов переменных!

var bag1, bag2, bag3, bag4 = 33, 44.5, "Hello", true
println(bag1, bag2, bag3, bag4)



//vBox := 11
//println(vBox)

//vFlo := 22.3
//println(vFol)

//vStr := "Hello"
//println(vStr)

//vBool := true 
//println(vBool)

} *960