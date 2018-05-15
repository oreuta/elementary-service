# Elementary-Service


=======
     Elementary Service is a web service providing REST API for the following elementary tasks:

- **Fibonacci sequence**

------------

Output all Fibonacci numbers that satisfy passed to the function restriction : are in the specified range
  or have the specified length.
  
  **Input parameters:** context object with fields min and max, or with field length
  
  **Output:** array of numbers
- **Lucky tickets**

------------
 There are 2 ways to calculate if the ticket is lucky:
 
  *Simple one* - if we have a ticket with six-figure number and the sum of the first three digits is equal to the sum of the last three, then this ticket is considered to be lucky.
  
  *Hard one* - if the sum of even numbers is equal to the sum of odd numbers,
  then the ticket is considered to be lucky.
  Determine which method of the calculation will contain more lucky tickets at a given interval.
  
 **Input parameters:** context object with min and max fields
 
  **Output:** an object with information about the winning method and the number of lucky tickets for each method.

------------


-  **Palindrome**

Check whether the string or part of it is a palindrome. For example, the string panama is not a palindrome, but its parts "ana" and "ama" is. Strings smaller than 2 symbols are not palindrome.
  **Input parameters:** string
  
  **Output:** string with found palindromes

------------


-  **Numerical sequence**

Output a comma-separated string of n length that consisting of natural
numbers, the square of which is not less than the given m.

**Input parameters:** the length and the min square value

**Output:** string with numerical sequence

------------


- **Square root**

Output array of square roots numbers of given numbers

**Input parameters:** the array of numbers

**Output:** the array of square roots numbers

------------


-  **Triangles sort**

Output vertices of triangle objects in descending order by their area.

**Input parameters:** array of triangle objects

 **Output:** an ordered array of triangle vertices

------------



<<<<<<< HEAD
1. _Task1_ - description
2. _Task2_ - description
3. _Task3_ - description
4. ...
5.
6. _Task6_ - The objective is to build the sequence of natural digits given the length of the sequence and minimal value.

Input data - the length itself and the value of the minimum (the sequence of squares (the meaning of each digit) shouldn't be less that this value).

Output data - the string with the sequence, commas are to be used as separators.
=======
>>>>>>> 18313a86f7bc812b614fd8d65c62a23caa52e06d
>>>>>>> f9ad387c294cc7a8231dfc42b139c26ff9de007b

Elementary-Service is deployed to [heroku](https://www.heroku.com/)