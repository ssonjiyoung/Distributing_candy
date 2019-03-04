# Distribuging candy..

Alice is a kindergarten teacher. She wants to give some candies to the children in her class.  All the children sit in a line and each of them has a rating score according to his or her performance in the class.  Alice wants to give at least 1 candy to each child. If two children sit next to each other, then the one with the higher rating must get more candies. Alice wants to minimize the total number of candies she must buy.

For example, assume her students' ratings are [4, 6, 4, 5, 6, 2]. She gives the students candy in the following minimal amounts: [1, 2, 1, 2, 3, 1]. She must buy a minimum of 10 candies.

 

###Function Description

Complete the candies function in the editor below. It must return the minimum number of candies Alice must buy.

candies has the following parameter(s):

·    n: an integer, the number of children in the class

·    arr: an array of integers representing the ratings of each student

 

###Input Format

The first line contains an integer, , the size of . 
Each of the next  lines contains an integer  indicating the rating of the student at position .

 

###Constraints

·    1 <= n <= 10^5

·    1 <= arr[i] <= 10^5


###Output Format

Output a single line containing the minimum number of candies Alice must buy.

 

# Example

## Sample 1 

###Input
``` sh
3

1

2

2
```

###Output
``` sh
4
```

###Explanation

Here 1, 2, 2 is the rating. Note that when two children have equal rating, they are allowed to have different number of candies. Hence optimal distribution will be 1, 2, 1.


 


## Sample 2

###Input
``` sh
10

2

4

2

6

1

7

8

9

2

1
```
###Output
``` sh
19
```
###Explanation

Optimal distribution will be 1,2,1,2,1,2,3,4,2,1




 
## Sample 3

###Input
``` sh
8

2

4

3

5

2

6

4

5
```

###Output
``` sh
12
```

###Explanation

Optimal distribution will be 1,2,1,2,1,2,1,2

 
 
## Sample 4

###Input
``` sh
8

77526

74568

4993

65614

53041

53041

45421

52081
```

###Output
``` sh
14
```

###Explanation 

Optimal distribution will be 3,2,1,2,1,2,1,2

 

 
## Sample 5

###Input
``` sh
See attachment file.
```
###Output
``` sh
42105
```

 