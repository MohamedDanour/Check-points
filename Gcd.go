package train

func Gcd(a, b uint) uint {
  for b != 0 {
    a , b = b, a%b
  }
  return a
}

/*

Instructions
Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

The hash equation is computed as follows:
(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

If the resulting character is unprintable add 33 to it.

*/