# Decode the stream cipher
Given the one time pad (stream ciphers) in a text file, decode the last cipher text and print the message.
The One time pad is a stream cipher that uses key equal to the length of the message.In our case , it is given that the One-Time Pad of all the cipher texts uses the same key.

## Logic
Given 2 messages m1 and m2 and their corresponding cipher texts c1 and c2, if the key used is same for One-Time Pad we can use the following to decode m1 and m2 - 
c1 ^ c2 = m1 ^ m2

Hint :- 'a'^' ' = 'A' and 'A'^' ' = 'a'

## Usage
go build cipher.go
./cipher --help

Example - ./cipher -file=sample.txt